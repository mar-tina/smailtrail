package dbclient

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

type IBadgerClient interface {
	OpenBadgerDB()
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
