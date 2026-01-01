package naming

import "testing"

func TestToSnake(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"Token", "token"},
		{"TestToken", "test_token"},
		{"ERC20Token", "erc20_token"},
		{"TokenV2", "token_v2"},
		{"token", "token"},
		{"token-name", "token_name"},
		{"token name", "token_name"},
		{"", ""},
	}
	for _, tt := range tests {
		got := ToSnake(tt.in)
		if got != tt.want {
			t.Fatalf("ToSnake(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestToLowerCamel(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"Token", "token"},
		{"token", "token"},
		{"T", "t"},
		{"", ""},
	}
	for _, tt := range tests {
		got := ToLowerCamel(tt.in)
		if got != tt.want {
			t.Fatalf("ToLowerCamel(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestToUpperCamel(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"token", "Token"},
		{"Token", "Token"},
		{"t", "T"},
		{"", ""},
	}
	for _, tt := range tests {
		got := ToUpperCamel(tt.in)
		if got != tt.want {
			t.Fatalf("ToUpperCamel(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
