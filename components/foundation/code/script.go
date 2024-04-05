package code

import "github.com/fengdotdev/coipoblog/pkg/gowebcomp/components"

type Script struct {
	Content string
}

func NewScript(code string) *Script {
	return &Script{
		Content: code,
	}
}

func (s *Script) Render() string {
	return `<script>` + s.Content + `</script>`
}

func (s *Script) String() string {
	return "Script: " + s.Content
}

func (s *Script) Html() string {
	return `<script>` + s.Content + `</script>`
}

func (s *Script) Iam() components.Component {
	return s
}
