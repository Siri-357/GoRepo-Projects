package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	/*http.HandleFunc("/hello", Hello)
	fmt.Println("Listneing on Port 8080...")
	http.ListenAndServe(":8080", nil)*/

	var wg sync.WaitGroup
	wg.Add(1)
	go Greet(&wg)
	wg.Wait()
	//time.Sleep(1 * time.Second)
	GoodBye()

}
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,World")
}

func Greet(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hey World!")
}
func GoodBye() {
	fmt.Println("GoodBye World!")
}
