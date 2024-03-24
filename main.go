package main

import (
	"fmt"
	"formcontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", formcontroller.RouteSubmitGet)
	http.HandleFunc("/result", formcontroller.RouteSubmitPost)

	fmt.Println("Server started at local:8080...")
	http.ListenAndServe(":8080", nil)
}
