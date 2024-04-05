package devserver

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port int
	Routes map[string]func(http.ResponseWriter, *http.Request)
	Mux *http.ServeMux
}

func NewServer() *Server {

	return &Server{
		Port: 8080,
		Routes: make(map[string]func(http.ResponseWriter, *http.Request)),
		Mux: http.NewServeMux(),
	}
}


func (s *Server) PopulateRoutes(routes ...map[string]func(http.ResponseWriter, *http.Request)) {

	if len(routes) > 0 {
		for _, route := range routes {
			for path, handler := range route {
				s.Routes[path] = handler
			}
		}
		return
	}
	// Default routes
	s.Routes = Routes
}


func (s *Server) Start(port... int) {
	if len(port) > 0 {
		s.Port = port[0]
	}
	addr := fmt.Sprintf(":%d", s.Port)
	fmt.Println("Server started at http://localhost" + addr)	


	for path, handler := range s.Routes {
		s.Mux.HandleFunc(path, handler)
	}
	http.ListenAndServe(addr, s.Mux)
}