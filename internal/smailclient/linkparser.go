package smailclient

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mar-tina/smailtrail/internal/models"
)

func ParseBody(headers []models.Part, docBody string) error {
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

		//Need to find a better way to do this.
		//FIX LATER
		if single.Text() == "Unsubscribe" || single.Text() == "unsubscribe" || single.Text() == "UNSUBSCRIBE" || strings.Contains(single.Text(), "Manage your Notifications") || checkIfHrefContainsUnsubscribe(single) {

			ret, exists := single.Attr("href")
			if !exists {
				log.Printf("Link Href does not exist %v", err.Error())
			}

			log.Printf("This is the selection %v \n", ret)

			unsubLink = ret
		}

	}

	fromVal = <-resultChannel

	err = DBClient.SaveSubscription(unsubLink, fromVal)
	if err != nil {
		log.Printf("DB subscription save failed %v", err.Error())
		return err
	}

	return nil

}

func checkIfHrefContainsUnsubscribe(link *goquery.Selection) bool {
	val, exists := link.Attr("href")
	if exists {
		return strings.Contains(val, "unsubscribe") || strings.Contains(val, "Unsubscribe")
	}
	return false
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
