package stdout

import (
	"fmt"

	"github.com/prodatalab/cbp"
	ba "github.com/prodatalab/msg/bytearray"
)

const (
	name = "stdout"
)

var (
	c   *cbp.Component
	err error
	msg []byte
)

// Init the component
func Init(urlStrings []string) {
	c, err = cbp.NewComponent(name)
	if err != nil {
		panic(err)
	}
	for _, u := range urlStrings {
		c.AddSocket(u)
	}
}

// Run this component
func Run() {
	c.Run()
	b := ba.ByteArray{}
	for {
		msg := c.Recv()
		b.UnmarshalMsg(msg)
		if b.Value != nil {
			fmt.Println("Value:", string(b.Value))
		} else {
			fmt.Println("Where is my value?")
		}
	}
}
