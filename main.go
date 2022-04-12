package main

import "http"

func main() {
	http.HandleFunc("/ping", func (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	}
}
