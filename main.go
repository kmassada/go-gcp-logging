package main

import (
	"encoding/json"
	"net/http"
	"cloud.google.com/go/logging"
	"context"
	"os"
)

type Response struct {
	Ip      string
	Headers http.Header
}

func printHeader(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := Response{r.RemoteAddr, r.Header}
	json.NewEncoder(w).Encode(response)

}
func main() {
	
	// context
	ctx := context.Background()
	client, err := logging.NewClient(ctx,  os.Getenv("PROJECT_ID"))
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
