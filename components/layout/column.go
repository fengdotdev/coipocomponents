package layout

import "github.com/fengdotdev/coipoblog/pkg/coipocom/components"


const ColumnType = "column"

type Column struct {
	Components []components.Component
}


func (c *Column) Html() string {
	panic("unimplemented")
}

func (c *Column) Render() string {
	panic("unimplemented")
}

func (c *Column) String() string {
	panic("unimplemented")
}

func (c *Column) Iam() (components.Component, string) {
	return c, ColumnType
}
