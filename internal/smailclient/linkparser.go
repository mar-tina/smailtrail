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
		if checkIfATagContainsUnsubscribe(strings.ToLower(single.Text())) || checkIfHrefContainsUnsubscribe(single) {

			ret, exists := single.Attr("href")
			if !exists {
				log.Printf("Link Href does not exist %v", err.Error())
			}

			log.Printf("This is the selection %v \n", ret)

			unsubLink = ret
		}

	}

	fromVal = <-resultChannel

	err = StormDBClient.SaveSubscription(unsubLink, fromVal)
	if err != nil {
		log.Printf("DB subscription save failed %v", err.Error())
		return err
	}

	return nil

}

func checkIfATagContainsUnsubscribe(link string) bool {
	return strings.Contains(link, "unsubscribe") || strings.Contains(link, "manage email preferences") || strings.Contains(link, "manage your notifications") || strings.Contains(link, "manage your email settings")
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
