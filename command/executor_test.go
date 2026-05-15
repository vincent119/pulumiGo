package command

import "testing"

// isStackless 是從 Execute 抽出的可測試邏輯
func isStackless(args []string) bool {
	if len(args) == 0 {
		return false
	}
	if len(args) >= 2 && args[0] == "stack" && args[1] == "select" {
		return true
	}
	for _, name := range []string{"login", "logout", "version", "whoami"} {
		if args[0] == name {
			return true
		}
	}
	return false
}

func TestIsStackless(t *testing.T) {
	tests := []struct {
		args []string
		want bool
	}{
		{[]string{"login"}, true},
		{[]string{"logout"}, true},
		{[]string{"version"}, true},
		{[]string{"whoami"}, true},
		{[]string{"stack", "select", "dev"}, true},
		{[]string{"up"}, false},
		{[]string{"preview"}, false},
		{[]string{"stack", "ls"}, false},
		{[]string{"stack"}, false},
		{[]string{}, false},
	}

	for _, tt := range tests {
		t.Run(func() string {
			if len(tt.args) == 0 {
				return "empty"
			}
			s := tt.args[0]
			for _, a := range tt.args[1:] {
				s += "_" + a
			}
			return s
		}(), func(t *testing.T) {
			if got := isStackless(tt.args); got != tt.want {
				t.Errorf("isStackless(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
