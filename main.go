package main

import (
	"fmt"
	"net/http"
)

const port = "8000"

func main() {
	http.HandleFunc("/api", DataController)

	fmt.Printf("Web server started on port %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	if err != nil {
		panic(err)
	}
}
