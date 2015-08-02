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
	// parse data form POST request
	if r.Form == nil {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("error:", err)
			http.Error(w, "error: parse form", http.StatusInternalServerError)
			return
		}
	}

	// create a new file
	f, err := os.Create(r.FormValue("token"))
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "error: create a new file", http.StatusInternalServerError)
		return
	}
	// write POST data to a file
	wrt := bufio.NewWriter(f)
	for k, values := range r.Form {
		for _, v := range values {
			_, err := wrt.WriteString(fmt.Sprintf("%s: %s\n", k, v))
			if err != nil {
				fmt.Println("error:", err)
				http.Error(w, "error: write to a file", http.StatusInternalServerError)
				return
			}
		}
	}
	err = wrt.Flush()
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "error: flushing buffer", http.StatusInternalServerError)
		return
	}
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
