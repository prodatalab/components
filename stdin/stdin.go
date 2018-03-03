package stdin

import (
	"bufio"
	"fmt"
	"os"

	"github.com/prodatalab/cbp"
	ba "github.com/prodatalab/msg/bytearray"
)

var (
	c   *cbp.Component
	err error
	msg []byte
)

// c.AddSocket("to-stdout", "push", "tcp", "tcp://127.0.0.1:5555")

// Init is called to initialize the component.
func Init(configSocketURL string, reportSocketURL string, dstreamURL string, dstreamTransportType string) {
	c, err = cbp.NewComponent("stdin")
	// fmt.Printf("INFO: new component created: %s\n", c.Name())
	if err != nil {
		panic(err.Error())
	}
	if dstreamURL != "" && dstreamTransportType != "" {
		c.AddSocket("stdout", cbp.SocketType("push"), cbp.TransportType(dstreamTransportType), dstreamURL)
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
