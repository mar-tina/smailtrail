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

func getNewClient(config *oauth2.Config, tok *oauth2.Token) *http.Client {
	tokFile := "token.json"

	saveToken(tokFile, tok)

	return config.Client(context.Background(), tok)
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
func Authorize(configFile, authcode string) *gmail.Service {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file %v", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope, gmail.GmailComposeScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authcode)
	if err != nil {
		log.Printf("Unable to retrieve token from web %v", err)
	}

	client := getNewClient(config, tok)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail Client")
	}

	return srv
}

//Processing Auth through http function helpers
func GetAuthURL(configFile string) (string, error) {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file %v", err)
		return "", err
	}

	GmailConfig, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope, gmail.GmailComposeScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret %v", err)
		return "", err
	}

	authURL := GmailConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	return authURL, nil

}

func GetSrvFromAuthCode(code string) *gmail.Service {
	log.Println("Calling srvfromauthcode")
	tok, err := GmailConfig.Exchange(context.TODO(), code)
	if err != nil {
		log.Printf("Unable to retrieve token from web %v", err)
	}

	log.Println("Calling srvfromauthcode step 2")
	client := GmailConfig.Client(context.Background(), tok)

	srv, err := gmail.New(client)
	if err != nil {
		log.Println("Unable to retrieve Gmail Client")
	}

	return srv
}
