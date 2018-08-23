package main

import (
	"encoding/json"
	"net/http"
	"cloud.google.com/go/logging"
	"google.golang.org/api/option"
	"context"
	"os"
	"fmt"
	stdlog "log"
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
	client, err := logging.NewClient(ctx, 
		os.Getenv("PROJECT_ID"),
		option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		panic(err)
	}

	// Initialize a logger
	lg := client.Logger("my-log")

	// Add entry to log buffer
	j := []byte(`{"Hostname": "`+os.Getenv("HOSTNAME")+`", "Count": 3}`)

	message := fmt.Sprintf("Data: %s", json.RawMessage(j))
	stdlog.Output(0, message)

	message = fmt.Sprintf("{Data: %s}", json.RawMessage(j))
	stdlog.Output(0, message)

	lg.Log(logging.Entry{
		Payload: message,
		Severity: logging.Critical,
	})

	// HTTP handler
	http.HandleFunc("/", printHeader)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
