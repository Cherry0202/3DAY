package sample

import "net/http"

func Hello(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	w.Write([]byte((msg)))
}