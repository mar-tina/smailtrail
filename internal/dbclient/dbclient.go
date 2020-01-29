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
	FetchSubscriptions(key string) ([]models.Subscription, error)
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

func (bc *BadgerClient) FetchSubscriptions(key string) ([]models.Subscription, error) {
	fromVal := strings.Replace(strings.Split(key, "<")[0], " ", "", -1)
	allsubs := []models.Subscription{}
	err := bc.badgerDB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 30
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Seek([]byte(fromVal)); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				sub := models.Subscription{}
				err := json.Unmarshal(v, &sub)
				if err != nil {
					log.Printf("Failed to marshal sub %v", err.Error())
				}
				allsubs = append(allsubs, sub)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return allsubs, err
}
