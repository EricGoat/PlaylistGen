package main

import (
	"fmt"
	"net/http"
)

const port = "8000"

type server struct{}

func (s *server) start() {
	fmt.Printf("Web server started on port %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}

func (s *server) addRoute(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, s.log(handler))
}

func (s *server) log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s: %s\n", r.Method, r.URL)
		h(w, r)
	}
}
