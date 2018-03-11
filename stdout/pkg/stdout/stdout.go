package stdout

import (
	"fmt"

	"github.com/prodatalab/cbp"
	ba "github.com/prodatalab/msg/bytearray"
	"github.com/sirupsen/logrus"
)

const (
	name = "stdout"
)

var (
	c   *cbp.Component
	err error
	msg []byte
	log = logrus.New()
)

// log.Level = logrus.DebugLevel

// Init the component
func Init(urlStrings []string) {
	c, err = cbp.NewComponent(name)
	if err != nil {
		log.Panic(err)
	}
	for _, u := range urlStrings {
		err = c.AddSocket(u)
		if err != nil {
			log.Panic(err)
		}
	}
	log.Info("stdout component created", "\n")
}

// Run this component
func Run() {

	err = c.Run()
	if err != nil {
		log.Panic(err)
	}
	b := ba.ByteArray{}
	b.Version = 1
	b.Type = 1
	for {
		msg := c.Recv()
		b.UnmarshalMsg(msg)
		if b.Value != nil {
			fmt.Println(string(b.Value))
		} else {
			log.Panic("b.Value is nil")
		}
	}
}
