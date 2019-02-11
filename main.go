package main

import (
	"context"
	"encoding/json"
	"fmt"
	stdlog "log"
	"net/http"
	"os"

	"cloud.google.com/go/logging"
	"google.golang.org/api/option"
)

// Response represents a response to http.
type Response struct {
	IP      string
	Headers http.Header
}

// Server to handle logging
type Server struct {
	lg *logging.Logger
}

func (s *Server) printHeader(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := Response{r.RemoteAddr, r.Header}
	json.NewEncoder(w).Encode(response)
	s.lg.Log(logging.Entry{
		Payload:  response,
		Severity: logging.Info,
	})

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

	s := Server{lg}

	// Add entry to log buffer
	j := []byte(`{"Data": {"Hostname": "` + os.Getenv("HOSTNAME") + `", "Count": 3}}`)

	message := fmt.Sprintf(`%s`, json.RawMessage(j))
	stdlog.Output(0, message)

	fmt.Println(message)

	lg.Log(logging.Entry{
		Payload:  message,
		Severity: logging.Critical,
	})

	lg.Log(logging.Entry{
		Payload:  json.RawMessage(j),
		Severity: logging.Critical,
	})

	// HTTP handler
	http.HandleFunc("/", s.printHeader)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
