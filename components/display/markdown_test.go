package display

import (
	"testing"

	"github.com/fengdotdev/coipoblog/pkg/gowebcomp/components"
)

var (
	markdownContent string = `
	# Hello World
	`
)

func TestMarkdown_Render(t *testing.T) {
	markdown := components.Markdown{
		MarkdownContent: markdownContent,
	}

	if markdown.Render() == "" {
		t.Error("Error rendering markdown")
	}
	t.Error(markdown.Render())
}