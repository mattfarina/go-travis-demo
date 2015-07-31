package main

import (
	"fmt"
	"net/http"
	"os"
)

func getFormatString() string {
	return "Hi there, I love %s"
}

func handler(w http.ResponseWriter, r *http.Request) {
	format := getFormatString()
	fmt.Fprintf(w, format, r.URL.Path[1:])
}

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
