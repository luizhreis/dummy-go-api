package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", randomNumberHandler)
	r.HandleFunc("/healthcheck", healthcheckHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	s := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}

func randomNumberHandler(w http.ResponseWriter, r *http.Request) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Fprintf(w, "{data:{number:%d}}", rand.Intn(100))
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{data:{status:ok}}")
}
