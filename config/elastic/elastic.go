package elastic

import (
	"context"
	"go-simple/config/env"
	"net"
	"net/http"
	"time"

	"sync"

	"github.com/labstack/gommon/log"
	"github.com/olivere/elastic/v7"
)

var (
	connection *elastic.Client
	mutex      sync.Mutex
)

func NewElastic() (c *elastic.Client, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	if connection == nil {
		c, err = newConnection()
	}

	return c, nil
}

func newConnection() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(env.Conf.ElasticHost),
		elastic.SetHttpClient(&http.Client{Transport: &http.Transport{ /* updated elastic to many connection start */
			MaxIdleConns:       100,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
		}}), /* updated elastic to many connection end */
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetHealthcheckInterval(12*time.Second),
	)

	if err != nil {
		log.Errorf("got an error while connecting elasticsearch server, error: ", err)
	}

	return client, err
}

func ValidIndex(e *elastic.Client, idx string) error {

	ctx := context.Background()

	// validation index
	isExist, err := e.IndexExists(idx).Do(ctx)
	if err != nil {
		log.Error("IndexCreate", err)
		return err
	}

	if !isExist {
		_, err := e.CreateIndex(idx).Do(ctx)

		if err != nil {
			log.Errorf("IndexCreate", err)
			return err
		}
	}

	return nil
}
