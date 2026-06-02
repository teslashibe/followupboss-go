package mcp

import "testing"

// maxToolDescription mirrors the host's mcptool description limit.
const maxToolDescription = 120

func TestToolsValid(t *testing.T) {
	tools := Provider{}.Tools()
	if len(tools) == 0 {
		t.Fatal("no tools registered")
	}
	seen := map[string]bool{}
	for _, tool := range tools {
		if tool.Name == "" {
			t.Error("tool with empty name")
		}
		if seen[tool.Name] {
			t.Errorf("duplicate tool name %q", tool.Name)
		}
		seen[tool.Name] = true
		if l := len(tool.Description); l == 0 || l > maxToolDescription {
			t.Errorf("tool %q description length %d out of range (1..%d)", tool.Name, l, maxToolDescription)
		}
		if tool.Invoke == nil {
			t.Errorf("tool %q has nil Invoke", tool.Name)
		}
	}
}

func TestPlatform(t *testing.T) {
	if got := (Provider{}).Platform(); got != "followupboss" {
		t.Errorf("Platform() = %q, want followupboss", got)
	}
}
