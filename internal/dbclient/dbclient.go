package dbclient

import (
	"log"
	"strings"

	"github.com/asdine/storm"
	"github.com/mar-tina/smailtrail/internal/models"
)

type IStormClient interface {
	OpenStormDB(dbname string)
	SaveSubscription(link, from, date string) error
	FetchSubscriptions(take, limit int) ([]models.Sub, error)
}

type StormClient struct {
	stormDB *storm.DB
}

func (sc *StormClient) OpenStormDB(dbname string) {
	var err error
	sc.stormDB, err = storm.Open(dbname)
	if err != nil {
		log.Printf("DB Failed to open %v", err.Error())
	}

	log.Printf("DB setup Done")
}

func (sc *StormClient) SaveSubscription(link, from, date string) error {
	subID := strings.Replace(strings.Split(from, "<")[0], " ", "", -1)
	sub := models.Sub{
		ID:     subID,
		Sender: from,
		Link:   link,
		Date:   date,
	}

	//Check if key already exists
	_, err := sc.stormDB.GetBytes("Sub", subID)
	if err == storm.ErrNotFound {
		err := sc.stormDB.Save(&sub)
		if err != nil {
			return err
		}
	}

	err = sc.stormDB.UpdateField(&models.Sub{ID: subID}, "Date", date)
	if err != nil {
		log.Printf("DB update failed %v", err.Error())
		return err
	}

	return nil
}

func (sc *StormClient) FetchSubscriptions(take, skip int) ([]models.Sub, error) {
	var subs []models.Sub
	err := sc.stormDB.All(&subs, storm.Limit(take), storm.Skip(skip))
	if err != nil {
		return nil, err
	}

	return subs, err
}
