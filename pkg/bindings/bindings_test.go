package bindings

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateBasic(t *testing.T) {
	t.Parallel()

	abiDir := t.TempDir()
	outDir := t.TempDir()
	abiPath := filepath.Join(abiDir, "Token.json")
	if err := os.WriteFile(abiPath, []byte(`[]`), 0o644); err != nil {
		t.Fatalf("write abi: %v", err)
	}

	var got struct {
		pkg      string
		typeName string
		outFile  string
	}
	var fixedPath string
	gen := DefaultGenerator()
	gen.Stdout = &bytes.Buffer{}
	gen.RunAbigen = func(_ string, _ string, pkg string, typeName string, outFile string) error {
		got.pkg = pkg
		got.typeName = typeName
		got.outFile = outFile
		return os.WriteFile(outFile, []byte("package "+pkg), 0o644)
	}
	gen.FixMethods = func(_ string, filePath string) error {
		fixedPath = filePath
		return nil
	}
	gen.GenTopics = func(_ string, _ string, _ string) error {
		t.Fatalf("unexpected GenTopics call")
		return nil
	}

	if err := gen.Generate(abiDir, outDir, "abigen", false); err != nil {
		t.Fatalf("Generate: %v", err)
	}

	wantOut := filepath.Join(outDir, "token", "token.go")
	if got.pkg != "token" {
		t.Fatalf("pkg = %q, want %q", got.pkg, "token")
	}
	if got.typeName != "Token" {
		t.Fatalf("type = %q, want %q", got.typeName, "Token")
	}
	if got.outFile != wantOut {
		t.Fatalf("outFile = %q, want %q", got.outFile, wantOut)
	}
	if fixedPath != wantOut {
		t.Fatalf("fix path = %q, want %q", fixedPath, wantOut)
	}
	if _, err := os.Stat(wantOut); err != nil {
		t.Fatalf("expected output file: %v", err)
	}
}

func TestGenerateNetworkDir(t *testing.T) {
	t.Parallel()

	abiDir := t.TempDir()
	outDir := t.TempDir()
	networkDir := filepath.Join(abiDir, "mainnet")
	if err := os.MkdirAll(networkDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	abiPath := filepath.Join(networkDir, "Token.json")
	if err := os.WriteFile(abiPath, []byte(`[]`), 0o644); err != nil {
		t.Fatalf("write abi: %v", err)
	}

	var gotOut string
	gen := DefaultGenerator()
	gen.Stdout = &bytes.Buffer{}
	gen.RunAbigen = func(_ string, _ string, _ string, _ string, outFile string) error {
		gotOut = outFile
		return os.WriteFile(outFile, []byte("package token"), 0o644)
	}
	gen.FixMethods = func(_ string, _ string) error { return nil }
	gen.GenTopics = func(_ string, _ string, _ string) error { return nil }

	if err := gen.Generate(abiDir, outDir, "abigen", false); err != nil {
		t.Fatalf("Generate: %v", err)
	}

	wantOut := filepath.Join(outDir, "mainnet", "token", "token.go")
	if gotOut != wantOut {
		t.Fatalf("outFile = %q, want %q", gotOut, wantOut)
	}
}

func TestGenerateTopics(t *testing.T) {
	t.Parallel()

	abiDir := t.TempDir()
	outDir := t.TempDir()
	abiPath := filepath.Join(abiDir, "Token.json")
	if err := os.WriteFile(abiPath, []byte(`[]`), 0o644); err != nil {
		t.Fatalf("write abi: %v", err)
	}

	var gotEvents string
	gen := DefaultGenerator()
	gen.Stdout = &bytes.Buffer{}
	gen.RunAbigen = func(_ string, _ string, pkg string, _ string, outFile string) error {
		return os.WriteFile(outFile, []byte("package "+pkg), 0o644)
	}
	gen.FixMethods = func(_ string, _ string) error { return nil }
	gen.GenTopics = func(_ string, _ string, outFile string) error {
		gotEvents = outFile
		return nil
	}

	if err := gen.Generate(abiDir, outDir, "abigen", true); err != nil {
		t.Fatalf("Generate: %v", err)
	}

	wantEvents := filepath.Join(outDir, "token", "tokenEvents.go")
	if gotEvents != wantEvents {
		t.Fatalf("events = %q, want %q", gotEvents, wantEvents)
	}
}

func TestGenerateSkipsNonJSON(t *testing.T) {
	t.Parallel()

	abiDir := t.TempDir()
	outDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(abiDir, "readme.txt"), []byte("nope"), 0o644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	gen := DefaultGenerator()
	gen.Stdout = &bytes.Buffer{}
	gen.RunAbigen = func(_ string, _ string, _ string, _ string, _ string) error {
		t.Fatalf("unexpected RunAbigen call")
		return nil
	}
	gen.FixMethods = func(_ string, _ string) error {
		t.Fatalf("unexpected FixMethods call")
		return nil
	}
	gen.GenTopics = func(_ string, _ string, _ string) error {
		t.Fatalf("unexpected GenTopics call")
		return nil
	}

	if err := gen.Generate(abiDir, outDir, "abigen", false); err != nil {
		t.Fatalf("Generate: %v", err)
	}

	entries, err := os.ReadDir(outDir)
	if err != nil {
		t.Fatalf("readdir: %v", err)
	}
	if len(entries) != 0 {
		names := make([]string, 0, len(entries))
		for _, ent := range entries {
			names = append(names, ent.Name())
		}
		t.Fatalf("unexpected output entries: %s", strings.Join(names, ","))
	}
}
