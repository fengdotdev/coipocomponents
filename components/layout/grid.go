package layout

import "github.com/fengdotdev/coipoblog/pkg/coipocom/components"


const GridType = "grid"


type Grid struct {
	Columns    int
	Rows       int
	Components map[int]components.Component
}


func (g *Grid) Html() string {
	panic("unimplemented")
}

func (g *Grid) Render() string {
	panic("unimplemented")
}

func (g *Grid) String() string {
	panic("unimplemented")
}

func (g *Grid) Iam() (components.Component, string) {
	return g, GridType
}