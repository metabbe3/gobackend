// handler.go
package handler

import (
	"fmt"
	"net/http"
)

// Handler is a simple HTTP handler function.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
