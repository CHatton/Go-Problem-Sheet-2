package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

/*
func handler(w http.ResponseWriter, r *http.Request) {
	htmlPage := `<!DOCTYPE html>
                    <html>
                        <body>
                            <h1>Guessing Game</h1>
                        </body>
                    </html>`

	fmt.Fprintf(w, htmlPage)
}
*/

func main() {
	port := getPort()
	http.Handle("/", http.FileServer(http.Dir("./res")))
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
		fmt.Println("Using default port value of 7777")
		port = "7777"
	}
	return port
}
