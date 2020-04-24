package main

import (
	"os"
	"fmt"
	"flag"
	"net"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

var usage = `Usage: ping host`

func checkErr(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func main(){
	// all flags
	sudo := flag.Bool("sudo", false, "")
	flag.Usage = func () {
		fmt.Println(usage)
	}
	flag.Parse()
	fmt.Println( *sudo, flag.NArg())

	if flag.NArg() == 0{
		flag.Usage()
		os.Exit(1)
	}

	// resolve the host
	host := flag.Arg(0)
	resolvedHost, err := net.ResolveIPAddr("ip" ,host)
	checkErr(err)

	// new connecion
	c, err := icmp.ListenPacket("udp", "0.0.0.0")
	checkErr(err)

	// var c icmp.PacketConn

	// new message
	message := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 8,
		Checksum: 1234,
		Body: &icmp.Echo{
			ID: os.Getpid(),
			Seq: 1,
			Data: []byte("You there?"),
		},
	}
	binaryMessage, err := message.Marshal(nil)
	checkErr(err)
	
	// send message
	_,err = c.WriteTo(binaryMessage, &net.UDPAddr{
		IP: resolvedHost.IP,
		Zone: resolvedHost.Zone,
	})
	checkErr(err)


	// read message
	rb := make([]byte, 1500)
	n, peer, err := c.ReadFrom(rb)
	checkErr(err)

	rm, err := icmp.ParseMessage(58, rb[:n])
	checkErr(err)

	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		fmt.Printf("got reflection from %v", peer)
	default:
		fmt.Printf("got %+v; want echo reply", rm)

	}

}

// TODO
// 1. resolve hostname, ip
// 2. icmp request
// 3. inturrupt

// Understanding
// open a raw socket to destination, while true or no inturrupt
// package echo request
// send,  with data field = timestamp
// filter out echo replies from incoming icmp messages
// filter on basis of Process id
// get rtt time