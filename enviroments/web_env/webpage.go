package webenv

import (
	"net/http"

	"github.com/fengdotdev/coipoblog/pkg/coipocom/components"
)


type WebPage struct {
	Title string
	FaviconUrl string
	Component components.Component
	ExternalCSSURL []string
	
}



func NewWebPage(w http.ResponseWriter, r *http.Request) *WebPage {
	return &WebPage{
		Title: "Empty Page",
		FaviconUrl: "",
	}
}