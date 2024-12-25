// Package ipage provides a collection of helpers for generating html pages from markdown contents, and reading configuration files
package ipage

import (
	"bytes"
	"errors"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/frontmatter"
)

type generator struct {
	Ctx parser.Context
	Buf *bytes.Buffer
}

func newgen() *generator {
	g := new(generator)
	g.Ctx = parser.NewContext()
	g.Buf = new(bytes.Buffer)

	return g
}

func (g *generator) Write(p []byte) (n int, err error) {
	n, err = g.Buf.Write(p)
	return n, err
}

// GenerateHtml generate html from markdown, use frt to decode the front matter
func GenerateHtml(mdpath string, frt *interface{}) (template.HTML, error) {
	g := newgen()
	cont, err := g.getContent(mdpath)
	if err != nil {
		return "", err
	}
	content := template.HTML(cont)
	err = g.decFmatter(frt)
	if err != nil {
		return content, err
	}
	return content, nil
}

// generateHtml generate the html template from markdown
func (g *generator) getContent(path string) (string, error) {
	gm := goldmark.New(
		goldmark.WithExtensions(
			&frontmatter.Extender{},
			extension.Table,
		),
		goldmark.WithRendererOptions(
			html.WithXHTML(),
		),
	)
	md, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	err = gm.Convert(md, g, parser.WithContext(g.Ctx))
	if err != nil {
		return "", err
	}
	content := g.Buf.String()
	return content, nil
}

// decFmatter decode page frontmatter info
func (g *generator) decFmatter(front *interface{}) error {
	data := frontmatter.Get(g.Ctx)
	if data == nil {
		return errors.New("no frontmatter found")
	}
	if err := data.Decode(&front); err != nil {
		return err
	}
	return nil
}

// ListContent return list of contents from root folder
func ListContent(root string) []string {
	var lst []string
	fileSystem := os.DirFS(root)
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && path != "." {
			lst = append(lst, path)
		}
		return nil
	})
	return lst
}

// GetPgname return the content page name from filepath
func GetPgname(path string) string {
	a := filepath.Dir(path)
	b := filepath.Base(a)
	return b
}
