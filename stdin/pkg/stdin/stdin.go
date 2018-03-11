package stdin

import (
	"bufio"
	"fmt"
	"os"

	"github.com/prodatalab/cbp"
	ba "github.com/prodatalab/msg/bytearray"
	"github.com/sirupsen/logrus"
)

const (
	name = "stdin"
)

var (
	c   *cbp.Component
	err error
	msg []byte
	log = logrus.New()
)

// log.Level = logrus.DebugLevel

// Init is called to initialize the component.
func Init(urlStrings []string) {
	c, err = cbp.NewComponent(name)
	if err != nil {
		log.Panic(err)
	}
	for _, u := range urlStrings {
		c.AddSocket(u)
	}
	log.Info("stdin component created", "\n")
}

// Run this component
func Run() {

	scanner := bufio.NewScanner(os.Stdin)
	err = c.Run()
	if err != nil {
		log.Panic(err)
	}
	b := ba.ByteArray{}
	b.Version = 1
	b.Type = 1
	fmt.Printf("stdin:> ")
	for scanner.Scan() {
		fmt.Printf("\nstdin:> ")
		b.Value = []byte(scanner.Text())
		if scanner.Err() != nil {
			log.Panic(err)
		}
		msg, err = b.MarshalMsg(nil)
		if err != nil {
			log.Panic(err)
		}
		c.Send(msg)
	}
	if err = scanner.Err(); err != nil {
		log.Panic(err)
	}
}
