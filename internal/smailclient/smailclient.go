package smailclient

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/mar-tina/smailtrail/internal/auth"
	"github.com/mar-tina/smailtrail/internal/models"
	"google.golang.org/api/gmail/v1"
)

type ISmailClient interface {
	ListLabels() ([]string, error)
	ListMessages(nextPageToken string) (models.GmailMsg, []models.Message, error)
	InitSmailClient(credFile, authcode string) error
	IndividualTrail(id string)
}

type SmailClient struct {
	srv *gmail.Service
}

func (smail *SmailClient) InitSmailClient(credFile, code string) error {

	smail.srv = auth.Authorize(credFile, code)
	if smail.srv == nil {
		return errors.New("Failed to init client")
	}

	return nil
}

func (smail *SmailClient) ListLabels() ([]string, error) {

	var labels []string
	r, err := smail.srv.Users.Labels.List("me").Do()
	if err != nil {
		return nil, err
	}

	for _, l := range r.Labels {
		labels = append(labels, fmt.Sprintf("%s", l.Name))
	}

	return labels, nil
}

func (smail *SmailClient) ListMessages(nextPageToken string) (models.GmailMsg, []models.Message, error) {
	var msgList models.GmailMsg
	var allMessages []models.Message

	r, err := smail.srv.Users.Messages.List("me").MaxResults(5).PageToken(nextPageToken).Do()
	if err != nil {
		log.Printf("ERROR: Failed to read messages %v", err.Error())
	}

	jsonBytes, _ := r.MarshalJSON()

	err = json.Unmarshal(jsonBytes, &msgList)
	if err != nil {
		log.Printf("Failed to unmarshal list DATA %v", err.Error())
	}

	for _, msg := range r.Messages {

		newmsg := models.Message{}

		s, err := smail.srv.Users.Messages.Get("me", msg.Id).Do()
		if err != nil {
			log.Printf("ERROR: Failed to fetch message %v", err.Error())

		}

		jsonBytes, _ := s.Payload.MarshalJSON()
		err = json.Unmarshal(jsonBytes, &newmsg)

		if err != nil {
			log.Printf("Failed to unmarshal DATA %v", err.Error())
		}

		for i := 0; i < len(newmsg.Parts); i++ {
			part, _ := base64.URLEncoding.DecodeString(newmsg.Parts[i].Body.Data)
			newmsg.Parts[i].Body.Data = string(part)
		}

		if len(newmsg.Parts) >= 1 {
			LoadBody(newmsg.Headers, newmsg.Parts[1].Body.Data)
		}

		allMessages = append(allMessages, newmsg)

	}

	return msgList, allMessages, nil
}

func (smail *SmailClient) IndividualTrail(id string) {

	s, err := smail.srv.Users.Threads.Get("me", id).Do()
	if err != nil {
		log.Println("Could not get thread %v", err.Error())
	}

	threadBytes, _ := s.MarshalJSON()
	log.Println("The thread", string(threadBytes))
}
