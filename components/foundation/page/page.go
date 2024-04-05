package page

import (
	"fmt"
	"net/http"

	"github.com/fengdotdev/coipoblog/pkg/gowebcomp/components"
	"github.com/fengdotdev/coipoblog/pkg/gowebcomp/helperfuncs"
)



var (
	pageTemplateHtml = `<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>%s</title>
		<link rel="icon" href="%s">
	</head>
	<body>
	%s
	</body>
	</html>`
)


type WebPage struct {
	Title string
	FaviconUrl string
	component components.Component
}




func Page(w http.ResponseWriter, r *http.Request,title string, faviconUrl string,component components.Component) {


	err, content := helperfuncs.RenderToHtml(component)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	page := []byte(fmt.Sprintf(pageTemplateHtml, title, faviconUrl, content))
	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}