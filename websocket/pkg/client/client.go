package client

import (
	"net/url"

	ws "github.com/gorilla/websocket"
	"github.com/prodatalab/cbp"
	ba "github.com/prodatalab/msg/bytearray"
	"github.com/sirupsen/logrus"
)

const (
	name = "websocket-client"
)

type (
	// Value blah
	Value struct {
		Sockets []string
		WSURL   string
	}
)

var (
	// Val blah
	Val   = Value{}
	c     *cbp.Component
	err   error
	msg   []byte
	wsURL string
	conn  *ws.Conn
	log   = logrus.New()
	addr  *url.URL
)

// log.Level = logrus.DebugLevel

func initialize() {
	c, err = cbp.NewComponent(name)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("cbp component created")
	for _, u := range Val.Sockets {
		c.AddSocket(u)
	}
	addr, err = url.Parse(Val.WSURL)
	if err != nil {
		log.Error(err)
	}
	u := addr.Scheme + "://" + addr.Host
	conn, _, err = ws.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("websocket client created")
}

// Run this component
func Run() {
	initialize()
	err = c.Run()
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	go func() {
		b := ba.ByteArray{}
		b.Version = 1
		b.Type = 1
		for {
			_, msg, err = conn.ReadMessage()
			if err != nil {
				log.Error(err)
				return
			}
			b.Value = msg
			msg, err = b.MarshalMsg(nil)
			if err != nil {
				log.Error(err)
			}
			c.Send(msg)
		}
	}()
	b := ba.ByteArray{}
	b.Version = 1
	b.Type = 1
	for {
		msg := c.Recv()
		if msg == nil {
			log.Error("msg is nil.. breaking from loop's iteration")
			continue
		}
		b.UnmarshalMsg(msg)
		if b.Value != nil {
			err = conn.WriteMessage(ws.TextMessage, b.Value)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
