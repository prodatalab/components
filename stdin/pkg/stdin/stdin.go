package stdin

import (
	"bufio"
	"fmt"
	"os"

	"github.com/prodatalab/cbp"
	ba "github.com/prodatalab/msg/bytearray"
)

const (
	name = "stdin"
)

var (
	c   *cbp.Component
	err error
	msg []byte
)

// Init is called to initialize the component.
func Init(urlStrings []string) {
	c, err = cbp.NewComponent(name)
	if err != nil {
		panic(err.Error())
	}
	for _, u := range urlStrings {
		c.AddSocket(u)
	}
}

// Run this component
func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	err = c.Run()
	if err != nil {
		fmt.Printf("ERROR: In stdin Run(): %s", err.Error())
		os.Exit(1)
	}
	b := ba.ByteArray{}
	b.Version = 1
	b.Type = 1
	for scanner.Scan() {
		b.Value = []byte(scanner.Text())
		if scanner.Err() != nil {
			panic(scanner.Err().Error())
		}
		msg, err = b.MarshalMsg(nil)
		if err != nil {
			panic(err)
		}
		c.Send(msg)
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
	}
}
