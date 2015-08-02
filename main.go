package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	PORT = flag.String("port", "8000", "listen on specified port")
)

// HandleIndex handles incoming request to "/"
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

// HandleReceive handles incoming POST request to "/receive". Each incoming
// email handled by custom logic based on request parameter.
func HandleReceive(w http.ResponseWriter, r *http.Request) {
	rbody := bufio.NewReader(r.Body)
	_, err := rbody.WriteTo(os.Stdin)
	if err != nil {
		http.Error(w, "failed to write request body", http.StatusInternalServerError)
	}
	fmt.Fprintf(os.Stdin, "%s", "\n")
}

func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", HandleIndex).Methods("GET")
	r.HandleFunc("/receive", HandleReceive).Methods("POST")

	http.Handle("/", r)
	log.Printf("start listening on :%s\n", *PORT)
	log.Fatal(http.ListenAndServe(":"+*PORT, nil))
}
