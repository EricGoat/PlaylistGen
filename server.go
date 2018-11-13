package main

import (
	"fmt"
	"net/http"
)

const port = "8000"

type Server struct{}

func (s *Server) start() {
	fmt.Printf("Web server started on port %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}

func (s *Server) addRoute(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, s.log(handler))
}

func (s *Server) log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s: %s\n", r.Method, r.URL)
		h(w, r)
	}
}
