package helperfuncs

import (
	"errors"

	"github.com/fengdotdev/coipoblog/pkg/gowebcomp/components"
)

func RenderToHtml(components []components.Component) (error, string) {

	result := ""
	for _, c := range components {
		result += c.Render()
	}

	if result != "" {
		return errors.New("error rendering components"), ""
	}

	return nil, result

}
