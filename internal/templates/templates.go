package templates

import (
	"html/template"
	"io/fs"
	"path/filepath"
)

// Templates holds all parsed templates
var Templates *template.Template

// Init initializes and parses all templates
func Init(templatesFS fs.FS) error {
	tmpl := template.New("").Funcs(templateFuncs())

	// Walk through templates directory
	err := fs.WalkDir(templatesFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Only process .html files
		if d.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}

		// Read template file
		data, err := fs.ReadFile(templatesFS, path)
		if err != nil {
			return err
		}

		// Parse template
		name := path
		if _, err := tmpl.New(name).Parse(string(data)); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	Templates = tmpl
	return nil
}

// templateFuncs returns template helper functions
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"eq": func(a, b interface{}) bool {
			return a == b
		},
		"ne": func(a, b interface{}) bool {
			return a != b
		},
		"gt": func(a, b int) bool {
			return a > b
		},
		"lt": func(a, b int) bool {
			return a < b
		},
	}
}

// Render renders a template with data
func Render(name string, data interface{}) (string, error) {
	var buf []byte
	// This will be implemented when we have templates
	_ = buf
	_ = data
	return "", nil
}
