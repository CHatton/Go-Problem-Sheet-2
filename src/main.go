package main

import (
	"./guessinggameui"
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var ui *guessinggameui.GuessingGameUI
var gameOver bool

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // seed rand so we get different values each time
	ui = guessinggameui.New("Guess a number between 1 and 20", "", "You haven't guessed yet!")
	port := getPort()
	gameOver = false // used as a flag to indicate if the player has guessed the correct number.

	http.HandleFunc("/guess/", guessingGame)

	http.HandleFunc("/newgame/", func(w http.ResponseWriter, r *http.Request) {
		ui = guessinggameui.New("Guess a number between 1 and 20", "", "You haven't guessed yet!")
		gameOver = false
		guessingGame(w, r)
	})

	// I consulted this article http://jessekallhoff.com/2013/04/14/go-web-apps-serving-static-files/
	// on how to serve specifc html pages. Not just index in the specified folder.
	// serves index.html in the html folder.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/index.html") // serve a static html file.
	})

	http.ListenAndServe(":"+port, nil) // will listen on the port specifed in the command line arguments.
}

func guessingGame(w http.ResponseWriter, r *http.Request) {
	// I consuled this question https://stackoverflow.com/questions/15407719/in-gos-http-package-how-do-i-get-the-query-string-on-a-post-request
	// on how to get the values received from a POST request, rather than a GET request.

	// from the documentation https://golang.org/pkg/net/http/#Request.ParseForm
	r.ParseForm() // populates r.Form, we can access the query parameters of a POST request there.

	var usersGuess string

	if userHasValidGuess(r) { // valid meaning it is a valid integer.
		usersGuess = r.FormValue("guess") // for a POST request.
		ui.Guess = string(usersGuess)     // update the struct that will be used in the template
	}

	if targetCookie, err := getTargetCookie(r); err == nil { // cookie exists, so we have a target value to guess already.
		// error from strconv.Atoi should always be nil unless we provide an invalid value in the Cookie in our code
		// so we ignore it
		targetAsInt, _ := strconv.Atoi(targetCookie.Value)
		userGuessAsInt, conversionErr := strconv.Atoi(usersGuess)
		if gameOver {
			ui.DisplayMessage = "Game over! Click new game to start again."
		} else if conversionErr != nil { // the user input was a valid integer
			ui.DisplayMessage = "Guess a number!"
		} else if userGuessAsInt < targetAsInt {
			ui.DisplayMessage = "You need to guess higher!"
		} else if userGuessAsInt > targetAsInt {
			ui.DisplayMessage = "You need to guess lower!"
		} else {
			gameOver = true // display will stop updating with the users' guess until they hit new game again.
			ui.DisplayMessage = "You guessed the number correctly!"
			target, _ := getTargetCookie(r)                       // update existing cookie instead of adding one with duplicate name
			target.Value = strconv.Itoa(rand.Intn(20) + 1)        // generate a new random number
			target.Expires = time.Now().Add(365 * 24 * time.Hour) // reset expiry date
			http.SetCookie(w, target)                             // save the cookie so this new target is saved
		}
	} else { // no cookie yet or the cookie expired.
		// I consulted this article https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/06.1.html
		// to learn how to set cookies in the response
		randomNum := rand.Intn(20) + 1                     // number between 1 - 20
		expiration := time.Now().Add(365 * 24 * time.Hour) // expires in 1 year
		cookie := http.Cookie{Name: "target", Value: strconv.Itoa(randomNum), Expires: expiration}
		http.SetCookie(w, &cookie)
	}

	// create the template from the template file.
	guessTemplate := template.Must( // program will throw an error if there is a problem here
		template.ParseFiles("./html/guess.tmpl"))
	guessTemplate.Execute(w, ui)
}

func getTargetCookie(r *http.Request) (*http.Cookie, error) {
	cookies := r.Cookies() // all the cookies in the response
	for _, cookie := range cookies {
		if cookie.Name == "target" {
			return cookie, nil
		}
	}
	return nil, errors.New("Cookie does not exist.")
}

func userHasValidGuess(r *http.Request) bool {
	guess := r.FormValue("guess")
	_, err := strconv.Atoi(guess)
	return err == nil // the guess can be converted into an int.
}

func hasCookies(r *http.Request) bool {
	return len(r.Cookies()) != 0
}

func getPort() string {
	// take in all the arguments bar the name of the file.
	args := os.Args[1:]
	var port string
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
