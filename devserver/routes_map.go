package devserver

import (
	"net/http"
	"github.com/fengdotdev/coipoblog/pkg/gowebcomp/devserver/routes"
)

var (
	Routes = map[string]func(http.ResponseWriter, *http.Request){
		"/": routes.Home,
		"/about": routes.About,
	}
)