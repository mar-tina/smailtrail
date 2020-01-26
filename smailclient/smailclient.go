package smailclient

import (
	"fmt"
	"log"

	"github.com/mar-tina/smailtrail/auth"
	"google.golang.org/api/gmail/v1"
)

type ISmailClient interface {
	ListLabels()
	NewSmailClient(creds string)
}

type SmailClient struct {
	srv *gmail.Service
}

func (smail *SmailClient) NewSmailClient(creds string) {
	smail.srv = auth.Authorize(creds)
}

func (smail *SmailClient) ListLabels() {
	r, err := smail.srv.Users.Labels.List("me").Do()
	if err != nil {
		log.Printf("ERROR: Failed to read labels %v", err.Error())
	}

	log.Println("There are the labels")
	for _, l := range r.Labels {
		fmt.Printf("- %s\n", l.Name)
	}
}
