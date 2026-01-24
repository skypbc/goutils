package internal

import "testing"

// TestCleanPath проверяет корректность очистки пути от служебных символов ./\ в начале и конце строки.
func TestCleanPath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Trim leading and trailing special chars, keep dots", `./.../\\.\\../////./.path1.`, `.path1.`},
		{"Nested path with dots and folders", `./.../\\.\\../////./.path2/aaa/.bbb/ccc...//.//...\\.`, `.path2/aaa/.bbb/ccc...`},
		{"Path ending with dot", `./.../\\.\\../////./path3.`, `path3.`},
		{"UTF-8 path with dots", `./папка/файл...////`, `папка/файл...`},
		{"Leading slashes with file", `////.///file`, `file`},
		{"Only ./ pattern", `/./`, ``},
		{"Only dots string", `.......`, ``},
		{"Only special chars", `////`, ``},
		{"Clean path without specials", `file`, `file`},
		{"Mixed slashes and dots with utf-8", `\\..///путь/к/файлу.`, `путь/к/файлу.`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CleanPath(tt.input)
			if got != tt.expected {
				t.Errorf("CleanPath(%q) = %q; expected %q",
					tt.input, got, tt.expected)
			}
		})
	}
}
