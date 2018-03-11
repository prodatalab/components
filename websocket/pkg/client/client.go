package client

import (
	"io"
	"net/url"

	ws "github.com/gorilla/websocket"
	"github.com/onrik/logrus/filename"
	"github.com/pkg/errors"
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
	Val  = Value{}
	c    *cbp.Component
	err  error
	msg  []byte
	conn *ws.Conn
	log  = logrus.New()
	addr *url.URL
)

func initialize() error {
	logrus.AddHook(filename.NewHook())
	// log.Level = logrus.DebugLevel
	c, err = cbp.NewComponent(name)
	if err != nil {
		return errors.Wrap(err, "NewComponent failed")
	}
	log.Info("cbp component created")
	for _, u := range Val.Sockets {
		err = c.AddSocket(u)
		if err != nil {
			return errors.Wrap(err, "AddSocket failed")
		}
	}
	addr, err = url.Parse(Val.WSURL)
	if err != nil {
		return errors.Wrap(err, "Parse failed")
	}
	u := addr.Scheme + "://" + addr.Host
	conn, _, err = ws.DefaultDialer.Dial(u, nil)
	if err != nil {
		return errors.Wrap(err, "Dial failed")
	}
	return nil
}

// Run this component
func Run() error {
	err = initialize()
	if err != nil {
		return errors.Wrap(err, "initialize failed")
	}
	err = c.Run()
	if err != nil {
		return errors.Wrap(err, "Run failed")
	}
	defer Close(conn)
	go func() {
		b := ba.ByteArray{}
		b.Version = 1
		b.Type = 1
		for {
			_, msg, err = conn.ReadMessage()
			if err != nil {
				log.Error("ReadMessage failed", err)
			}
			b.Value = msg
			msg, err = b.MarshalMsg(nil)
			if err != nil {
				log.Error("MarshalMsg failed", err)
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
				log.Error("WriteMessage failed", err)
			}
		}
	}
	return nil
}

// Close handles deferred Close funcs that return an error
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err, "Close failed")
	}
}
