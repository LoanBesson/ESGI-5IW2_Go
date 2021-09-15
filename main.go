package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", hourHandler)
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(":9000", nil)
}

func hourHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, time.Now().Format("15h04"))
	}
}

func addHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		author := req.FormValue("author")
		entry := req.FormValue("entry")

		file, _ := os.OpenFile("lines", os.O_WRONLY|os.O_CREATE, 0644)
		defer file.Close()

		file.WriteString(author + ":" + entry + "\n")
		fmt.Fprintf(rw, author+":"+entry)
	}
}
