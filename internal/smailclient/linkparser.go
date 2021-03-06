package smailclient

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mar-tina/smailtrail/internal/models"
)

func ParseBody(headers []models.Part, docBody string) error {
	var unsubLink string
	fromChannel := make(chan string)
	dateChannel := make(chan string)

	var fromVal string
	var dateVal string

	go func() {
		from, date := returnSenderandDateValue(headers)
		fromChannel <- from
		dateChannel <- date
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

			unsubLink = ret
		}

	}

	fromVal = <-fromChannel
	dateVal = <-dateChannel

	err = StormDBClient.SaveSubscription(unsubLink, fromVal, dateVal)
	if err != nil {
		log.Printf("DB subscription save failed %v", err.Error())
		return err
	}

	return nil

}

func checkIfATagContainsUnsubscribe(link string) bool {
	return strings.Contains(link, "unsubscribe") || strings.Contains(link, "manage email preferences") || strings.Contains(link, "manage your notifications") || strings.Contains(link, "manage your email settings") || strings.Contains(link, "email subscription preferences") || strings.Contains(link, "communication settings")
}

func checkIfHrefContainsUnsubscribe(link *goquery.Selection) bool {
	val, exists := link.Attr("href")
	if exists {
		return strings.Contains(strings.ToLower(val), "unsubscribe")
	}
	return false
}

func returnSenderandDateValue(headers []models.Part) (name string, date string) {

	for _, val := range headers {
		if val.Name == "From" {
			name = val.Value
		}
		if val.Name == "Date" {
			date = val.Value
		}
	}

	return name, date
}
