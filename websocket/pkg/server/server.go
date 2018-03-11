package server

import (
	"io"
	"net/http"
	"net/url"

	ws "github.com/gorilla/websocket"
	"github.com/onrik/logrus/filename"
	"github.com/pkg/errors"
	"github.com/prodatalab/cbp"
	ba "github.com/prodatalab/msg/bytearray"
	"github.com/sirupsen/logrus"
)

const (
	name = "websocket-server"
)

// Value blah
type (
	Value struct {
		Sockets []string
		WSURL   string
	}
)

var (
	// Val blah
	Val      = Value{}
	upgrader = ws.Upgrader{}
	log      = logrus.New()
	c        *cbp.Component
	err      error
	msg      []byte
	conn     *ws.Conn
	mt       int
	addr     *url.URL
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
	http.HandleFunc("/", home)
	if err != nil {
		return errors.Wrap(err, "HandleFunc failed")
	}
	return nil
}

// Run blah
func Run() error {
	err = initialize()
	if err != nil {
		return errors.Wrap(err, "initialize failed")
	}
	err = c.Run()
	if err != nil {
		return errors.Wrap(err, "Run failed")
	}
	log.Info("websocket server created")
	err = http.ListenAndServe(addr.Host, nil)
	if err != nil {
		return errors.Wrap(err, "ListenAndServe failed")
	}
	return nil
}

func home(w http.ResponseWriter, r *http.Request) {
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Upgrade failed", err)
	}
	defer Close(conn)
	go func() {
		b := ba.ByteArray{}
		b.Version = 1
		b.Type = 1
		cnt := 0
		for {
			cnt++
			mt, msg, err = conn.ReadMessage()
			if err != nil {
				log.Error("ReadMessage failed", err)
				continue
			}
			b.Value = msg
			msg, err = b.MarshalMsg(nil)
			if err != nil {
				log.Error("MarshalMsg failed", err)
				continue
			}
			c.Send(msg)
		}
	}()
	b := ba.ByteArray{}
	b.Version = 1
	b.Type = 1
	for {
		msg := c.Recv()
		b.Value, err = b.UnmarshalMsg(msg)
		if err != nil {
			log.Error("UnmarshalMsg failed", err)
			continue
		}
		if b.Value == nil {
			log.Warn("Value returned is nil")
			continue
		}
		err = conn.WriteMessage(ws.TextMessage, b.Value)
		if err != nil {
			log.Error("WriteMessage failed", err)
		}
	}
}

// Close handles deferred Close funcs that return an error
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err, "Close failed")
	}
}
