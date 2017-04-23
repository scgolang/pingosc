package main

import (
	"context"
	"log"
	"net"
	"runtime"

	"github.com/scgolang/osc"
)

func main() {
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	srv, err := osc.ListenUDPContext(context.Background(), "udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening at %s", srv.LocalAddr().String())

	dispatcher := osc.Dispatcher{
		"/ping": osc.Method(func(m osc.Message) error {
			return srv.SendTo(m.Sender, osc.Message{Address: "/pong"})
		}),
	}
	if err := srv.Serve(runtime.NumCPU(), dispatcher); err != nil {
		log.Fatal(err)
	}
}
