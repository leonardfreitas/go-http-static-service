package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")

	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	http.HandleFunc("/time", getTime)

	log.Println("[HTTP_STATIC_SERVICE] Running")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
