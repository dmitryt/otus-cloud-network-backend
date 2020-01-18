package main

import (
    "net/http"
    "log"
    "os"
    "time"
    "encoding/json"

    "github.com/gorilla/mux"
)

type Data struct {
	Path string `json:"path"`
	Timestamp string `json:"timestamp"`
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

		api := r.PathPrefix("/api").Subrouter()
    api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNotFound)
		})

		api1 := api.PathPrefix("/v1").Subrouter()

    api1.HandleFunc("/send/{greetings}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
				value := vars["greetings"]
				d := Data{Path: value, Timestamp: time.Now().String()}
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				json.NewEncoder(w).Encode(d)
		})

		r.Use(loggingMiddleware)

    http.ListenAndServe(":8080", r)
}