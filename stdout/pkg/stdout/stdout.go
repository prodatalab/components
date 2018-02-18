package stdout

import (
	"fmt"

	"github.com/prodatalab/cbp"
	msg "github.com/prodatalab/messages"
)

var (
	c   *cbp.Component
	err error
	b   msg.ByteArray
)

// Init the component
func Init(configSocketURL string, reportSocketURL string, dstreamURL string, dstreamTransportType string) {
	c, err = cbp.NewComponent("stdout")
	fmt.Printf("INFO: new component created: %s\n", c.Name())
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
	for {
		msg := c.Recv()
		fmt.Printf("MSGPACKED: %s\n", msg)
		//
		//
		//
		b.Unpack(msg)
		if b.Value != nil {
			fmt.Println(string(b.Value))
		} else {
			fmt.Println("Where is my value?")
		}
	}
}
