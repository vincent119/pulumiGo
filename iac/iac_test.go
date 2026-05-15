package iac

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

// parseStackJSON 是從 StackCheck 抽出的可測試邏輯
func parseStackJSON(output []byte) (string, error) {
	var stacks []map[string]interface{}
	if err := json.Unmarshal(output, &stacks); err != nil {
		return "", err
	}
	for _, stack := range stacks {
		if stack["current"] == true {
			name, ok := stack["name"].(string)
			if !ok {
				return "", errors.New("invalid stack name")
			}
			parts := strings.Split(name, "/")
			return parts[len(parts)-1], nil
		}
	}
	return "", errors.New("current stack not found")
}

func TestParseStackJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "simple stack name",
			input: `[{"name":"dev","current":true}]`,
			want:  "dev",
		},
		{
			name:  "org/project/stack format",
			input: `[{"name":"myorg/myproject/prod","current":true}]`,
			want:  "prod",
		},
		{
			name:  "multiple stacks, second is current",
			input: `[{"name":"dev","current":false},{"name":"prod","current":true}]`,
			want:  "prod",
		},
		{
			name:    "no current stack",
			input:   `[{"name":"dev","current":false}]`,
			wantErr: true,
		},
		{
			name:    "empty list",
			input:   `[]`,
			wantErr: true,
		},
		{
			name:    "invalid json",
			input:   `not json`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStackJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr=%v, got err=%v", tt.wantErr, err)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
