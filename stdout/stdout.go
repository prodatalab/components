package stdout

import (
	"fmt"

	"github.com/prodatalab/cbp"
	ba "github.com/prodatalab/msg/bytearray"
)

var (
	c   *cbp.Component
	err error
	msg []byte
)

// Init the component
func Init(configSocketURL string, reportSocketURL string, dstreamURL string, dstreamTransportType string) {
	c, err = cbp.NewComponent("stdout")
	// fmt.Printf("INFO: new component created: %s\n", c.Name())
	if err != nil {
		fmt.Println(err.Error())
	}
	// c.AddConfigSocket(configSocketURL)
	// c.AddReportSocket("stdout", reportSocketURL)
	if dstreamURL != "" && dstreamTransportType != "" {
		c.AddSocket("stdin", cbp.SocketType("pull"), cbp.TransportType(dstreamTransportType), dstreamURL)
	}
}

// Run this component
func Run() {
	c.Run()
	b := ba.ByteArray{}
	for {
		msg := c.Recv()
		// fmt.Printf("MSGPACKED: %s\n", msg)
		//
		//
		//
		b.UnmarshalMsg(msg)
		if b.Value != nil {
			fmt.Println("Value:", string(b.Value))
		} else {
			fmt.Println("Where is my value?")
		}
	}
}
