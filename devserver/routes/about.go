package routes

import "net/http"

func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}