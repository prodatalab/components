package server

import (
	"net/http"
	"net/url"

	ws "github.com/gorilla/websocket"
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
		log.Fatal(err)
	}
	http.HandleFunc("/", home)
	if err != nil {
		log.Error(err)
	}
}

// Run blah
func Run() {
	initialize()
	c.Run()
	log.Info("websocket server created")
	log.Fatal(http.ListenAndServe(addr.Host, nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()
	go func() {
		b := ba.ByteArray{}
		b.Version = 1
		b.Type = 1
		cnt := 0
		for {
			cnt++
			mt, msg, err = conn.ReadMessage()
			if err != nil {
				log.Error(err)
				break
			}
			b.Value = msg
			msg, err = b.MarshalMsg(nil)
			if err != nil {
				log.Error(err)
				break
			}
			c.Send(msg)
		}
	}()
	b := ba.ByteArray{}
	b.Version = 1
	b.Type = 1
	for {
		msg := c.Recv()
		b.UnmarshalMsg(msg)
		if b.Value != nil {
			err = conn.WriteMessage(ws.TextMessage, b.Value)
			if err != nil {
				log.Error(err)
				return
			}
		}
	}
}
