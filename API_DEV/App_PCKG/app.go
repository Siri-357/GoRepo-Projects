package App

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//define router
	//router:=http.NewServeMux()
	router := mux.NewRouter()

	router.HandleFunc("/greet", Greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_ID:[0-9]+}", GetCustomer_ID).Methods(http.MethodGet)
	//strating server
	fmt.Println("Listening to Server")
	//log.Fatal(http.ListenAndServe("localhost:8080", router))
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
func GetCustomer_ID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_ID"])
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {

	//vars := mux.Vars(r)
	fmt.Fprint(w, "Post request received")
}

