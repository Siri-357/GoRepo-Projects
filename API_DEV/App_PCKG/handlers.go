package App

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"Full_Name" xml:"Name"`
	City    string `json:"City" xml:"City"`
	Zipcode string `json:"Zip_Code" xml:"Zip"`
}
type CustomersWrapper struct {
	XMLName   xml.Name   `xml:"Customers"`
	Customers []Customer `xml:"Customer"`
}

func Greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello Server!")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Siri", City: "Plano", Zipcode: "75024"},
		{Name: "Konda", City: "Plano", Zipcode: "75024"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		wrappedCustomers := CustomersWrapper{Customers: customers}
		xml.NewEncoder(w).Encode(wrappedCustomers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
	/*format := r.URL.Query().Get("format")
	if format == "xml" {
		w.Header().Add("Content-Type", "application/xml")
		wrappedCustomers := CustomersWrapper{Customers: customers}

		if err := xml.NewEncoder(w).Encode(wrappedCustomers); err != nil {
			http.Error(w, "Error encoding XML", http.StatusInternalServerError)
			return
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(customers); err != nil {
			http.Error(w, "Error encoding XML", http.StatusInternalServerError)
			return
		}
	}*/
}
