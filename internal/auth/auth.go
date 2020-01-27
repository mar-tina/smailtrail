package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

var GmailConfig *oauth2.Config

func getClient(config *oauth2.Config) *http.Client {
	// tokFile := "token.json"
	// tok, err := tokenFromFile(tokFile)
	// if err != nil {
	// 	tok = getTokenFromWeb(config)
	// 	saveToken(tokFile, tok)
	// }

	tok := getTokenFromWeb(config)

	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf("Go to the following link in your browser and then type the authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("unable to read the authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential to %v\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}

	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

//Authorize oauth setup to access Gmail
func Authorize(configFile string) *gmail.Service {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file %v", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope, gmail.GmailComposeScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret %v", err)
	}

	client := getClient(config)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail Client")
	}

	return srv
}

//Processing Auth through http function helpers
func GetAuthURL(configFile string) string {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file %v", err)
	}

	GmailConfig, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope, gmail.GmailComposeScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret %v", err)
	}

	authURL := GmailConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	return authURL

}

func GetSrvFromAuthCode(code string) *gmail.Service {
	tok, err := GmailConfig.Exchange(context.TODO(), code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	client := GmailConfig.Client(context.Background(), tok)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail Client")
	}

	return srv
}
