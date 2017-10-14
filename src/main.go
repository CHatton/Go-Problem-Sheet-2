package main

import (
	"./message"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func main() {
	msg := message.New("Guess a number between 1 and 20")
	port := getPort()
	// I consulted this article http://jessekallhoff.com/2013/04/14/go-web-apps-serving-static-files/
	// on how to server specifc html pages. Not just index in the specified folder.
	http.HandleFunc("/guess/", func(w http.ResponseWriter, r *http.Request) {

		// create the template from the template file.
		guessTemplate := template.Must( // program will throw an error if there is a problem here
			template.ParseFiles("./html/guess.tmpl"))

		guessTemplate.Execute(w, msg)
	})
	// serves index.html in the res folder.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/index.html")
	})

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
