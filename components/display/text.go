package display

import "fmt"

type Text struct {
	Text   string
	Size   int    // if 0, then default size will be used
	Color  string // if empty, then default color will be used
	Font   string // if empty, then default font will be used
	Bold   bool
	Italic bool
}

func (t *Text) Render() string {
	size := t.Size
	if size == 0 {
		size = 12
	}
	color := t.Color
	if color == "" {
		color = "black"
	}
	font := t.Font
	if font == "" {
		font = "Arial"
	}
	bold := ""
	if t.Bold {
		bold = "font-weight: bold;"
	}
	italic := ""
	if t.Italic {
		italic = "font-style: italic;"
	}

	stringSize := fmt.Sprintf("%d", size)
	return "<span style=\"font-family: " + font + "; font-size: " + stringSize + "px; color: " + color + "; " + bold + " " + italic + "\">" + t.Text + "</span>"
}


func (t *Text) String() string {
	return t.Text
}


func (t *Text) iam() Component {
	return t
}