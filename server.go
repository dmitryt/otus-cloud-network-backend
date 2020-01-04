package main

import (
    "net/http"
    "log"
    "os"
    "encoding/json"

    "github.com/gorilla/mux"
)

type Data struct {
	Type string
	Value string
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do stuff here
			log.Println(r.RequestURI)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
	})
}

func main() {
		r := mux.NewRouter()
		logger := log.New(os.Stdout, "http: ", log.LstdFlags)
		logger.Println("Server is starting...")

    r.HandleFunc("/greetings/{greetings}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
				value := vars["greetings"]
				d := Data{Type: "greetings", Value: value}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(d)
		})

		r.Use(loggingMiddleware)

    http.ListenAndServe(":80", r)
}