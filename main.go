package main

import (
	"net/http"
	"log"
	"path"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "xiaoliang2233")
	r.ParseForm()
	if (r.URL.Path == "/") {
		fp := path.Join("./dist/index.html")
		http.ServeFile(w, r, fp)
	} else {
		fp := path.Join("./dist" + r.URL.Path)
		http.ServeFile(w, r, fp)
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9898", nil)
	if (err != nil) {
		log.Fatal("listrnAndServe", err)
	}
}