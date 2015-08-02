package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	PORT = flag.String("port", "8000", "listen on specified port")
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", HandleIndex)

	http.Handle("/", r)
	log.Printf("start listening on :%s\n", *PORT)
	log.Fatal(http.ListenAndServe(":"+*PORT, nil))

}
