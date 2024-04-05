package media

import "github.com/fengdotdev/coipoblog/pkg/coipocom/components"


const ImageType = "image"

type Image struct {
	Src string // Specifies the path to the image
	Alt string // Specifies an alternate text for the image
}

func (i *Image) Render() string {
	return `<img src="` + i.Src + `" alt="` + i.Alt + `">`
}

func (i *Image) String() string {
	return "Image: " + i.Src
}

func (i *Image) Html() string {
	return `<img src="` + i.Src + `" alt="` + i.Alt + `">`
}
func (i *Image) Iam() (components.Component,string) {
	return i,ImageType
}