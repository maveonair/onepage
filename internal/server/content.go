package server

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"os"
)

func (s *server) ReadPageFile() ([]byte, error) {
	content, err := os.ReadFile(s.pageFilePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *server) MarkdownToHTML() (string, error) {
	content, err := s.ReadPageFile()
	if err != nil {
		return "", err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(content)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer)), nil
}

func (s *server) WritePageFile(bytes []byte) error {
	err := os.WriteFile(s.pageFilePath, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
