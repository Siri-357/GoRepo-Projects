package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}

	fmt.Fprintf(w, "Post Request Successful\n")
	Name := r.FormValue("name")
	Address := r.FormValue("address")
	fmt.Fprintf(w, "Name : %s\n", Name)
	fmt.Fprintf(w, "Address : %s\n", Address)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello Daksh!❤️🤗")

}

func main() {
	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", FileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", HelloHandler)

	fmt.Printf("Starting Server at Port 8080:\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
