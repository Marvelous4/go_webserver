package main

import (
        "fmt"
        "log"
        "net/http"
)

func main() {
        mux := http.NewServeMux()

        v1Mux := http.NewServeMux()

        v1Mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "v1 Profile")
        })

        v1Mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "v1 Posts")
        })

        v2Mux := http.NewServeMux()

        v2Mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "v2 Profile")
        })

        v2Mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "v2 Posts")
        })

        mux.Handle("/v1/", http.StripPrefix("/v1", v1Mux))
        mux.Handle("/v2/", http.StripPrefix("/v2", v2Mux))

        loggedHandler := loggingMiddleware(mux)

        if err := http.ListenAndServe(":3001", loggedHandler); err != nil {
                fmt.Println(err)
        }
}

func loggingMiddleware(handler http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                log.Printf("Got a %s request for: %v\n", r.Method, r.URL)
                handler.ServeHTTP(w, r)
                log.Printf("Handler finished processing request")
        })
}