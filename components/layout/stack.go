package layout

import "github.com/fengdotdev/coipoblog/pkg/coipocom/components"

const StackType = "stack"

type Stack struct {
	Components []components.Component
}

// Html implements components.Component.
func (s *Stack) Html() string {
	panic("unimplemented")
}

// Render implements components.Component.
func (s *Stack) Render() string {
	panic("unimplemented")
}

// String implements components.Component.
func (s *Stack) String() string {
	panic("unimplemented")
}

func (s *Stack) Iam() (components.Component, string) {
	return s, StackType
}
