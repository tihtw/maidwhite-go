package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tihtw/maidwhite-go"
)

func main() {
	ctx := context.Background()
	// conf := &oauth2.Config{
	// 	ClientID:     "YOUR_CLIENT_ID",
	// 	ClientSecret: "YOUR_CLIENT_SECRET",
	// 	Scopes:       []string{"SCOPE1", "SCOPE2"},
	// 	Endpoint: oauth2.Endpoint{
	// 		AuthURL:  "https://provider.com/o/oauth2/auth",
	// 		TokenURL: "https://provider.com/o/oauth2/token",
	// 	},
	// }

	clientID := os.Getenv("MAIDWHITE_CLIENT_ID")
	clientSecret := os.Getenv("MAIDWHITE_CLIENT_SECRET")
	clientRedirectURL := os.Getenv("MAIDWHITE_REDRIECT_URL")
	fmt.Println("Got client id: " + clientID)
	fmt.Println("Got client secret: " + clientSecret)
	fmt.Println("Got client redirect url: " + clientRedirectURL)
	if clientID == "" {
		fmt.Println("use export MAIDWHITE_CLIENT_ID= to set env")
		return
	}
	if clientSecret == "" {
		fmt.Println("use export MAIDWHITE_CLIENT_SECRET= to set env")
		return
	}

	client := maidwhiteclient.NewClient()
	client.SetClientID(clientID)
	client.SetClientSecret(clientSecret)
	client.SetRedirectURL(clientRedirectURL)

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := client.AuthCodeURL("state")
	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	fmt.Printf("code: ")
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	fmt.Println("code:", code)
	tok, err := client.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Access token:", tok.AccessToken)
	u, err := client.GetUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello World", u.DisplayName)

}
