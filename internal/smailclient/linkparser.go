package smailclient

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mar-tina/smailtrail/internal/dbclient"
	"github.com/mar-tina/smailtrail/internal/models"
)

var DBClient dbclient.IBadgerClient

func LoadBody(headers []models.Part, docBody string) {
	var unsubLink string
	resultChannel := make(chan string)

	var fromVal string
	go func() {
		x := returnSenderValue(headers)
		resultChannel <- x
	}()

	body := strings.NewReader(docBody)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	links := doc.Find("a")

	for i := range links.Nodes {
		single := links.Eq(i)

		if single.Text() == "Unsubscribe" || single.Text() == "unsubscribe" {

			ret, exists := single.Attr("href")
			if !exists {
				log.Printf("Link Href does not exist %v", err.Error())
			}

			log.Printf("This is the selection %v \n", ret)

			unsubLink = ret
		}

	}

	fromVal = <-resultChannel

	DBClient.SaveSubscription(unsubLink, fromVal)
	log.Printf("From %v All the links %v", fromVal, unsubLink)
}

func returnSenderValue(headers []models.Part) string {
	var name string
	for _, val := range headers {
		if val.Name == "From" {
			name = val.Value
		}
	}

	return name
}
