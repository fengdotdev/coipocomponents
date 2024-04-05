package layout

import "github.com/fengdotdev/coipoblog/pkg/coipocom/components"

const RowType = "row"

type Row struct {
	Components []components.Component
}

// Html implements components.Component.
func (r *Row) Html() string {
	panic("unimplemented")
}

// Render implements components.Component.
func (r *Row) Render() string {
	panic("unimplemented")
}

// String implements components.Component.
func (r *Row) String() string {
	panic("unimplemented")
}

func (r *Row) Iam() (components.Component, string) {
	return r, RowType
}
