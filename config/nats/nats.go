package nats

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"kumparan/config/env"

	"time"

	"github.com/labstack/gommon/log"
	"github.com/nats-io/nats.go"
)

// nats publish
func Publish(subj string, m interface{}) error {
	var (
		natsUrl string
	)
	// Connect Options.
	opts := []nats.Option{nats.Name("NATS")}
	if env.Conf.NatsHost == "" {
		natsUrl = nats.DefaultURL
	} else {
		natsUrl = fmt.Sprintf("nats://%s", env.Conf.NatsHost)
	}

	// Connect to NATS
	nc, err := nats.Connect(natsUrl, opts...)
	if err != nil {
		return err
	}
	defer nc.Close()

	msg, err := WriteMessage(m)
	if err != nil {
		return err
	}

	if err := nc.Publish(subj, msg); err != nil {
		return err
	}

	nc.Flush()

	if err := nc.LastError(); err != nil {
		return err
	}

	log.Printf("Published [%s] : '%s'\n", subj, msg)
	return nil
}

// nats subscription
func Subscription(subj string, c nats.MsgHandler) {
	var (
		natsUrl string
	)
	// Connect Options.
	opts := []nats.Option{nats.Name("NATS")}
	opts = SetupConnOptions(opts)

	if env.Conf.NatsHost == "" {
		natsUrl = nats.DefaultURL
	} else {
		natsUrl = fmt.Sprintf("nats://%s", env.Conf.NatsHost)
	}

	// Connect to NATS
	nc, err := nats.Connect(natsUrl, opts...)
	if err != nil {
		log.Fatal(err)
	}

	_, err = nc.Subscribe(subj, c)

	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

}

// option of nats
func SetupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}

func WriteMessage(m interface{}) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func ReadMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
