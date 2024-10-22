package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type TimeFormat struct {
	Current_Time string `json:"current_time"`
}
type MultipleLOC map[string]string

func main() {
	/*router := mux.NewRouter()
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
	*/
	http.HandleFunc("/api/time", TimeNow)
	fmt.Println("Listening to Server")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func TimeNow(w http.ResponseWriter, r *http.Request) {
	tzParam := r.URL.Query().Get("tz")

	if tzParam == "" {
		utcTime := time.Now().UTC().Format(time.RFC3339)
		respose := TimeFormat{Current_Time: utcTime}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(respose)
		return
	}
	timezones := strings.Split(tzParam, ",")
	response := make(MultipleLOC)
	for _, tz := range timezones {
		tz = strings.TrimSpace(tz) // Remove any extra spaces

		loc, err := time.LoadLocation(tz)
		if err != nil {
			// If the timezone is invalid, return a 404 error with an appropriate message
			http.Error(w, "invalid timezone", http.StatusNotFound)
			return
		}

		// Get the current time in the specified timezone
		currentTime := time.Now().In(loc).Format(time.RFC3339)
		response[tz] = currentTime
	}

	// Set the content type to JSON and encode the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
