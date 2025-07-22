package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет от Air!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Сервер запущен на :1488")
	http.ListenAndServe(":1488", nil)
}
