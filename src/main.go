package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Guessing Game")
}

func main() {
	port := getPort()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}

func getPort() string {
	// take in all the arguments bar the name of the file.
	args := os.Args[1:]
	port := ""
	if len(args) == 1 { // user provided one command line argument
		if _, err := strconv.Atoi(args[0]); err == nil {
			port = args[0] // can be successfully parsed as a number.
		} else { // invalid command line argument provided.
			fmt.Println("Please provide a valid port number.\nExiting...")
			os.Exit(0)
		}
	} else {
		fmt.Println("Using default port value of 8080")
		port = "8080"
	}
	return port
}
