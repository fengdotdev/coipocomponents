package routes

import (
	"net/http"

	"github.com/fengdotdev/coipoblog/pkg/gowebcomp/components"
	"github.com/fengdotdev/coipoblog/pkg/gowebcomp/page"
)

func Home(w http.ResponseWriter, r *http.Request) {
	page.Page(w, r, "Home", "/favicon.ico", components.("Hello World"))
}