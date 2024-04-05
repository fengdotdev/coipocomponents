package display

import (
	"bytes"
	"html/template"
)

var (
	templMarkdown = `
	<div class="markdown-body">
	{{.MarkdownContent}}
	</div>
	`
)

type Markdown struct {
	MarkdownContent string
}


func NewMarkdown(content string) *Markdown {
	return &Markdown{
		MarkdownContent: content,
	}
}

func (m *Markdown) Render() string {

	tmp,err := template.New("markdown").Parse(templMarkdown)
	if err != nil {
		return ""
	}


    var buf bytes.Buffer
	err = tmp.Execute(&buf, m)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (m *Markdown) String() string {
	return m.MarkdownContent
}

func (m *Markdown) iam() Component {
	return m
}