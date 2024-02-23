package main

import (
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", indexHtml)
	http.HandleFunc("/first.css", firstCss)
	http.HandleFunc("/second.css", secondCss)

	http.ListenAndServe(":3000", nil)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
  body, _ := os.ReadFile("index.html")
  w.Write(body)
}

func firstCss(w http.ResponseWriter, r *http.Request) {
  body, _ := os.ReadFile("first.css")
  w.Write(body)
}

func secondCss(w http.ResponseWriter, r *http.Request) {
  body, _ := os.ReadFile("second.css")
  time.Sleep(time.Second)
  w.Write(body)
}
