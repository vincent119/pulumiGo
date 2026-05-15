package iac

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestRestoreRegexFromYAML(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"escaped slash", `\/path\/to`, "/path/to"},
		{"double backslash", `\\d+`, `\d+`},
		{"no escaping", `^abc$`, `^abc$`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RestoreRegexFromYAML(tt.input); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSanitizeRegexForYAML(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"plain slash needs escaping", "/path/to", `\/path\/to`},
		{"already escaped", `\/path\/to`, `\/path\/to`},
		{"no slash", `^abc$`, `^abc$`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SanitizeRegexForYAML(tt.input)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTruncateString(t *testing.T) {
	tests := []struct {
		input  string
		maxLen int
		want   string
	}{
		{"hello", 10, "hello"},
		{"hello world", 5, "hello..."},
		{"", 5, ""},
		{"abc", 3, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := truncateString(tt.input, tt.maxLen); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestLocateYAMLError(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		if got := LocateYAMLError("content", nil); got != "No error" {
			t.Errorf("got %q", got)
		}
	})

	t.Run("error without line number", func(t *testing.T) {
		err := errors.New("some yaml error")
		got := LocateYAMLError("line1\nline2", err)
		if got == "No error" {
			t.Error("expected non-empty error string")
		}
	})

	t.Run("error with line number", func(t *testing.T) {
		err := errors.New("yaml: line 2: mapping values are not allowed here")
		content := "key: value\nbad: : value\nkey2: value2"
		got := LocateYAMLError(content, err)
		if got == "No error" {
			t.Error("expected error context string")
		}
	})
}

func TestReadWriteYamlFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.yaml")

	data := map[string]interface{}{
		"name":    "test",
		"version": "1.0",
	}

	if err := WriteYamlFile(path, data); err != nil {
		t.Fatalf("WriteYamlFile: %v", err)
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("file not created: %v", err)
	}

	got, err := ReadYamlFile(path)
	if err != nil {
		t.Fatalf("ReadYamlFile: %v", err)
	}

	if got["name"] != "test" {
		t.Errorf("name: got %v, want test", got["name"])
	}
}

func TestReadYamlFile_NotFound(t *testing.T) {
	_, err := ReadYamlFile("/nonexistent/path/file.yaml")
	if err == nil {
		t.Error("expected error for missing file")
	}
}
