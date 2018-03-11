package stdout

import (
	"fmt"

	"github.com/onrik/logrus/filename"
	"github.com/pkg/errors"
	"github.com/prodatalab/cbp"
	ba "github.com/prodatalab/msg/bytearray"
	"github.com/sirupsen/logrus"
)

const (
	name = "stdout"
)

// Value blah
type Value struct {
	Sockets []string
}

var (
	// Val blah
	Val = Value{}
	c   *cbp.Component
	err error
	msg []byte
	log = logrus.New()
)

func initialize() error {
	logrus.AddHook(filename.NewHook())
	// log.Level = logrus.DebugLevel
	c, err = cbp.NewComponent(name)
	if err != nil {
		return errors.Wrap(err, "NewComponent failed")
	}
	log.Info("stdout component created")
	for _, u := range Val.Sockets {
		err = c.AddSocket(u)
		if err != nil {
			return errors.Wrap(err, "AddSocket failed")
		}
	}
	return nil
}

// Run this component
func Run() {
	err = initialize()
	if err != nil {
		log.Fatal("initialize failed", err)
	}
	err = c.Run()
	if err != nil {
		log.Fatal("Run failed", err)
	}
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
			log.Warn("received Value is nil")
		}
		fmt.Println(string(b.Value))
	}
}
