package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/dgraph-io/badger"
	"github.com/mar-tina/smailtrail/internal/models"
)

type IBadgerClient interface {
	OpenBadgerDB(dbname string)
	SaveSubscription(link, from string) error
}

type BadgerClient struct {
	badgerDB *badger.DB
}

func (bc *BadgerClient) OpenBadgerDB(dbname string) {

	var err error
	bc.badgerDB, err = badger.Open(badger.DefaultOptions(dbname))
	if err != nil {
		log.Printf("DB Open Failed with err %s\n", err.Error())
	}

	fmt.Println("DB Setup Done")
}

func (bc *BadgerClient) SaveSubscription(link, from string) error {
	sub := models.Subscription{}
	fromVal := strings.Replace(strings.Split(from, "<")[0], " ", "", -1)
	sub.Sender = from
	sub.Link = link
	jsonBytes, _ := json.Marshal(sub)
	err := bc.badgerDB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(fromVal), jsonBytes)
		if err != nil {
			log.Printf("Insert subscription failed %v", err.Error())
			return err
		}
		return nil
	})

	return err
}
