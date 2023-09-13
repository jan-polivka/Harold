package main

import (
	"fmt"
	"io"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func form(w http.ResponseWriter, req *http.Request) {
	enableCORS(&w)
	for name, headers := range req.Header {
		for _, header := range headers {
			fmt.Println(name, " ", header)
		}
	}

	body, bodyErr := io.ReadAll(req.Body)
	fmt.Println(bodyErr)
	fmt.Println(string(body))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/v1/form", form)

	http.ListenAndServe(":8080", nil)
}
