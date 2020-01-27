package smailclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mar-tina/smailtrail/internal/auth"
	"github.com/mar-tina/smailtrail/internal/models"
	"google.golang.org/api/gmail/v1"
)

type ISmailClient interface {
	ListLabels() ([]string, error)
	ListMessages()
	NewSmailClient(creds string)
	InitSmailClient(srv *gmail.Service)
}

type SmailClient struct {
	srv *gmail.Service
}

func (smail *SmailClient) InitSmailClient(srv *gmail.Service) {
	smail.srv = srv
}

func (smail *SmailClient) NewSmailClient(creds string) {
	smail.srv = auth.Authorize(creds)
}

func (smail *SmailClient) ListLabels() ([]string, error) {
	r, err := smail.srv.Users.Labels.List("me").Do()
	if err != nil {
		log.Printf("ERROR: Failed to read labels %v", err.Error())
	}

	var labels []string
	log.Println("There are the labels")
	for _, l := range r.Labels {
		// fmt.Printf("- %s\n", l.Name)
		labels = append(labels, fmt.Sprintf("%s", l.Name))
	}

	log.Printf("These are the labels %v", labels)

	return labels, err
}

func (smail *SmailClient) ListMessages() {
	r, err := smail.srv.Users.Messages.List("me").MaxResults(1).IncludeSpamTrash(true).Do()
	if err != nil {
		log.Printf("ERROR: Failed to read messages %v", err.Error())
	}

	log.Println("These are the messages", r.Messages)

	for _, msg := range r.Messages {
		s, err := smail.srv.Users.Messages.Get("me", msg.Id).Do()
		if err != nil {
			log.Printf("ERROR: Failed to get message %v", err.Error())
		}

		newmsg := models.Message{}
		jsonBytes, err := s.Payload.MarshalJSON()

		err = json.Unmarshal(jsonBytes, &newmsg)
		if err != nil {
			log.Printf("Failed to unmarshal DATA %v", err.Error())
		}

		body := newmsg.Parts[0].Body.Data

		data, _ := base64.URLEncoding.DecodeString(body)
		html := string(data)
		fmt.Println(html)

	}

}
