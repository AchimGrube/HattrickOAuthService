package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	_ "github.com/joho/godotenv/autoload"
)

const (
	listeningAddress string = "[::]:4222"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	http.HandleFunc("/authorize", authorize)

	http.HandleFunc("/authenticate", authenticate)

	fmt.Printf("listening at %s ...", listeningAddress)
	http.ListenAndServe(listeningAddress, nil)
}

func authorize(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "keep-alive")

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	config := oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    "oob",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: "https://chpp.hattrick.org/oauth/request_token.ashx",
			AuthorizeURL:    "https://chpp.hattrick.org/oauth/authorize.aspx",
			AccessTokenURL:  "https://chpp.hattrick.org/oauth/access_token.ashx",
		},
	}

	requestToken, _, _ := config.RequestToken()
	authorizationURL, _ := config.AuthorizationURL(requestToken)
	http.Redirect(w, r, authorizationURL.String(), http.StatusTemporaryRedirect)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "keep-alive")

	// TODO receive pin from client app
	// pin := r.Header.Get("PIN")
}
