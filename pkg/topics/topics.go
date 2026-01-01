package topics

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type eventInfo struct {
	Name      string
	Signature string
}

func GenerateEventTopics(abiPath, pkg, outFile string) error {
	events, err := parseEvents(abiPath)
	if err != nil {
		return err
	}
	if len(events) == 0 {
		return nil
	}

	var buf bytes.Buffer
	buf.WriteString("// Code generated - DO NOT EDIT.\n\n")
	buf.WriteString("package " + pkg + "\n\n")
	buf.WriteString("import (\n")
	buf.WriteString("\t\"github.com/ethereum/go-ethereum/common\"\n")
	buf.WriteString("\t\"github.com/ethereum/go-ethereum/crypto\"\n")
	buf.WriteString(")\n\n")

	buf.WriteString("const (\n")
	for _, ev := range events {
		buf.WriteString(fmt.Sprintf("\t%sEvent = %q\n", ev.Name, ev.Name))
	}
	buf.WriteString(")\n\n")

	buf.WriteString("var (\n")
	for _, ev := range events {
		buf.WriteString(fmt.Sprintf("\t%sEventSignature = crypto.Keccak256Hash([]byte(%q))\n", ev.Name, ev.Signature))
	}
	buf.WriteString(")\n\n")

	buf.WriteString("var Events = map[string]common.Hash{\n")
	for _, ev := range events {
		buf.WriteString(fmt.Sprintf("\t%sEvent: %sEventSignature,\n", ev.Name, ev.Name))
	}
	buf.WriteString("}\n")

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(outFile), 0o755); err != nil {
		return err
	}
	return os.WriteFile(outFile, src, 0o644)
}

func parseEvents(path string) ([]eventInfo, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	abiRaw, err := extractABI(raw)
	if err != nil {
		return nil, err
	}

	parsed, err := abi.JSON(bytes.NewReader(abiRaw))
	if err != nil {
		return nil, err
	}

	var events []eventInfo
	for _, ev := range parsed.Events {
		if ev.Anonymous {
			continue
		}
		sig := buildSignature(ev)
		events = append(events, eventInfo{Name: ev.Name, Signature: sig})
	}

	sort.Slice(events, func(i, j int) bool { return events[i].Name < events[j].Name })
	return events, nil
}

func extractABI(raw []byte) ([]byte, error) {
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 {
		return nil, errors.New("empty ABI")
	}

	if trimmed[0] == '[' {
		return trimmed, nil
	}

	var artifact struct {
		ABI json.RawMessage `json:"abi"`
	}
	if err := json.Unmarshal(raw, &artifact); err != nil {
		return nil, err
	}
	if len(artifact.ABI) == 0 {
		return nil, errors.New("missing abi field")
	}
	return artifact.ABI, nil
}

func buildSignature(ev abi.Event) string {
	parts := make([]string, len(ev.Inputs))
	for i, input := range ev.Inputs {
		parts[i] = input.Type.String()
	}
	return fmt.Sprintf("%s(%s)", ev.Name, strings.Join(parts, ","))
}
