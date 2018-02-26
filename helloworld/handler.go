package helloworld

import "net/http"

// ServeHTTP is our header entry point
func ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Hello, World"))

	return nil
}
