package bindings

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sergeyWh1te/abigenx/pkg/naming"
	"github.com/sergeyWh1te/abigenx/pkg/topics"
)

type Generator struct {
	RunAbigen  func(abigenPath, abiPath, pkg, typeName, outFile string) error
	FixMethods func(abiPath, filePath string) error
	GenTopics  func(abiPath, pkg, outFile string) error
	WalkDir    func(string, fs.WalkDirFunc) error
	Rel        func(string, string) (string, error)
	MkdirAll   func(string, os.FileMode) error
	Stdout     io.Writer
}

func DefaultGenerator() Generator {
	return Generator{
		RunAbigen:  runAbigen,
		FixMethods: fixAbigenMethods,
		GenTopics:  topics.GenerateEventTopics,
		WalkDir:    filepath.WalkDir,
		Rel:        filepath.Rel,
		MkdirAll:   os.MkdirAll,
		Stdout:     os.Stdout,
	}
}

func Generate(abiDir, outDir, abigenPath string, genTopics bool) error {
	gen := DefaultGenerator()
	return gen.Generate(abiDir, outDir, abigenPath, genTopics)
}

func (g Generator) Generate(abiDir, outDir, abigenPath string, genTopics bool) error {
	return g.WalkDir(abiDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".json") {
			return nil
		}

		rel, err := g.Rel(abiDir, path)
		if err != nil {
			return err
		}

		segments := strings.Split(rel, string(filepath.Separator))
		if len(segments) < 1 {
			return nil
		}

		network := ""
		if len(segments) > 1 {
			network = segments[0]
		}

		contractName := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		pkg := naming.ToSnake(contractName)
		fileBase := naming.ToLowerCamel(contractName)
		typeName := naming.ToUpperCamel(contractName)

		targetDir := filepath.Join(outDir, network, pkg)
		if err := g.MkdirAll(targetDir, 0o755); err != nil {
			return err
		}

		outFile := filepath.Join(targetDir, fileBase+".go")
		fmt.Fprintf(g.Stdout, "processing %s -> %s\n", path, outFile)
		if err := g.RunAbigen(abigenPath, path, pkg, typeName, outFile); err != nil {
			return fmt.Errorf("abigen %s: %w", path, err)
		}

		if err := g.FixMethods(path, outFile); err != nil {
			return fmt.Errorf("fix methods %s: %w", outFile, err)
		}

		if genTopics {
			eventsFile := filepath.Join(targetDir, fileBase+"Events.go")
			fmt.Fprintf(g.Stdout, "processing %s -> %s\n", path, eventsFile)
			if err := g.GenTopics(path, pkg, eventsFile); err != nil {
				return fmt.Errorf("generate topics %s: %w", path, err)
			}
		}

		return nil
	})
}

func runAbigen(abigenPath, abiPath, pkg, typeName, outFile string) error {
	cmd := exec.Command(abigenPath,
		"--abi", abiPath,
		"--pkg", pkg,
		"--type", typeName,
		"--out", outFile,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type abiEntry struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func fixAbigenMethods(abiPath, filePath string) error {
	entries, err := loadABIEntries(abiPath)
	if err != nil {
		return err
	}

	mapping := buildMapping(entries)
	if len(mapping) == 0 {
		return nil
	}

	return rewriteFile(filePath, mapping)
}

func loadABIEntries(path string) ([]abiEntry, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var direct []abiEntry
	if err := json.Unmarshal(raw, &direct); err == nil {
		return direct, nil
	}

	var artifact struct {
		ABI []abiEntry `json:"abi"`
	}
	if err := json.Unmarshal(raw, &artifact); err != nil {
		return nil, err
	}
	return artifact.ABI, nil
}

func buildMapping(entries []abiEntry) map[string]string {
	res := make(map[string]string)
	for _, e := range entries {
		if e.Type != "function" {
			continue
		}
		if !strings.Contains(e.Name, "_") {
			continue
		}
		if e.Name != strings.ToUpper(e.Name) {
			continue
		}
		alias := strings.ReplaceAll(e.Name, "_", "")
		if alias == e.Name || alias == "" {
			continue
		}
		res[alias] = e.Name
	}
	return res
}

func rewriteFile(path string, mapping map[string]string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	ast.Inspect(node, func(n ast.Node) bool {
		id, ok := n.(*ast.Ident)
		if !ok {
			return true
		}
		if newName, exists := mapping[id.Name]; exists {
			id.Name = newName
		}
		return true
	})

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	cfg := printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}
	return cfg.Fprint(out, fset, node)
}
