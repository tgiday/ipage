package ipage

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/tgiday/geezdate"
)

// GetTempFilesFromFolders return a slice of all template files path from input of folders and nil error ,or return error.
func GetTempFilesFromFolders(folders []string) ([]string, error) {
	var filepaths []string
	for _, folder := range folders {
		files, err := os.ReadDir(folder)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			if strings.Contains(file.Name(), ".html") {
				filepaths = append(filepaths, folder+file.Name())
			}
		}
	}
	return filepaths, nil
}

// NewTemp return new template.
func NewTemp() *template.Template {
	t := template.New("").Funcs(template.FuncMap{
		"generator": ipagegenerator,
		"time":      tim,
		"parfunc":   par,
		"print":     prin,
	})
	return t
}
func prin(x string) string {
	return fmt.Sprint(x)
}
func par(x string) template.HTML {
	ht, _ := os.ReadFile("layouts/partials/" + x)
	htm := fmt.Sprintf(string(ht))
	html := template.HTML(htm)
	return html
}
func ipagegenerator() template.HTML {
	htm := fmt.Sprintf(`<meta name="generator" content=%s>`, "ipage")
	html := template.HTML(htm)
	return html
}
func tim(s string) string {
	p := geezdate.Today()
	return s + p.String()
}
