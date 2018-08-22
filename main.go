package main

import (
	"encoding/json"
	"net/http"
	"cloud.google.com/go/logging"
)

type Response struct {
	Ip      string
	Headers http.Header
}

func printHeader(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := Response{r.RemoteAddr, r.Header}
	json.NewEncoder(w).Encode(response)
	lg.Log(logging.Entry{Payload: response})

}
func main() {
	
	// context
	ctx := context.Background()
	client, err := logging.NewClient(ctx, "go-gcp-logging")
	if err != nil {
		panic(err)
	}

	// Initialize a logger
	lg := client.Logger("my-log")

	// Add entry to log buffer
	lg.Log(logging.Entry{Payload: "something happened!"})

	// HTTP handler
	http.HandleFunc("/", printHeader)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
