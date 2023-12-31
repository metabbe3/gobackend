// main.go
package main

import (
	"net/http"

	handler "github.com/metabbe3/gobackend/api/handler"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe(":8080", nil)
}
