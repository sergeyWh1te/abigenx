package topics

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func TestExtractABI(t *testing.T) {
	rawArray := []byte(`[{"type":"event","name":"Ping","inputs":[]}]`)
	got, err := extractABI(rawArray)
	if err != nil {
		t.Fatalf("extractABI array: %v", err)
	}
	if string(got) != string(rawArray) {
		t.Fatalf("extractABI array mismatch: %s", string(got))
	}

	rawArtifact := []byte(`{"abi":[{"type":"event","name":"Pong","inputs":[]}]}`)
	got, err = extractABI(rawArtifact)
	if err != nil {
		t.Fatalf("extractABI artifact: %v", err)
	}
	if !strings.Contains(string(got), `"Pong"`) {
		t.Fatalf("extractABI artifact missing event: %s", string(got))
	}
}

func TestBuildSignature(t *testing.T) {
	addrType, err := abi.NewType("address", "", nil)
	if err != nil {
		t.Fatalf("new type address: %v", err)
	}
	uintType, err := abi.NewType("uint256", "", nil)
	if err != nil {
		t.Fatalf("new type uint256: %v", err)
	}
	ev := abi.Event{
		Name: "Transfer",
		Inputs: abi.Arguments{
			{Type: addrType},
			{Type: addrType},
			{Type: uintType},
		},
	}
	got := buildSignature(ev)
	want := "Transfer(address,address,uint256)"
	if got != want {
		t.Fatalf("buildSignature = %q, want %q", got, want)
	}
}

func TestParseEvents(t *testing.T) {
	dir := t.TempDir()
	abiPath := filepath.Join(dir, "Token.json")
	raw := `[
		{"type":"event","name":"Mint","inputs":[{"name":"to","type":"address"}]},
		{"type":"event","name":"Burn","anonymous":true,"inputs":[{"name":"from","type":"address"}]},
		{"type":"function","name":"balanceOf","inputs":[{"name":"owner","type":"address"}]}
	]`
	if err := os.WriteFile(abiPath, []byte(raw), 0o644); err != nil {
		t.Fatalf("write abi: %v", err)
	}

	events, err := parseEvents(abiPath)
	if err != nil {
		t.Fatalf("parseEvents: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("events = %d, want 1", len(events))
	}
	if events[0].Name != "Mint" {
		t.Fatalf("event name = %q, want %q", events[0].Name, "Mint")
	}
}

func TestGenerateEventTopics(t *testing.T) {
	dir := t.TempDir()
	abiPath := filepath.Join(dir, "Token.json")
	outFile := filepath.Join(dir, "tokenEvents.go")
	raw := `[{"type":"event","name":"Mint","inputs":[{"name":"to","type":"address"}]}]`
	if err := os.WriteFile(abiPath, []byte(raw), 0o644); err != nil {
		t.Fatalf("write abi: %v", err)
	}

	if err := GenerateEventTopics(abiPath, "token", outFile); err != nil {
		t.Fatalf("GenerateEventTopics: %v", err)
	}

	content, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("read outFile: %v", err)
	}
	if !strings.Contains(string(content), "MintEvent") {
		t.Fatalf("output missing event const: %s", string(content))
	}
}
