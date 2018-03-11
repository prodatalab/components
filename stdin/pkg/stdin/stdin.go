package stdin

import (
	"bufio"
	"fmt"
	"os"

	"github.com/onrik/logrus/filename"
	"github.com/pkg/errors"
	"github.com/prodatalab/cbp"
	ba "github.com/prodatalab/msg/bytearray"
	"github.com/sirupsen/logrus"
)

const (
	name = "stdin"
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
	log.Info("stdin component created")
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
	scanner := bufio.NewScanner(os.Stdin)
	err = c.Run()
	if err != nil {
		log.Fatal("Run failed", err)
	}
	b := ba.ByteArray{}
	b.Version = 1
	b.Type = 1
	fmt.Printf("stdin:> ")
	for scanner.Scan() {
		fmt.Printf("\nstdin:> ")
		b.Value = []byte(scanner.Text())
		if scanner.Err() != nil {
			log.Error("scanner failed", err)
			continue
		}
		msg, err = b.MarshalMsg(nil)
		if err != nil {
			log.Error("MarshalMsg failed", err)
			continue
		}
		c.Send(msg)
	}
	if err = scanner.Err(); err != nil {
		log.Panic(err)
	}
}
