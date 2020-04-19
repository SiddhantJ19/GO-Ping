package main

import (
	"log"
	"os"
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	ipv4 bool
	ipv6 bool
)


func main(){
	app := &cli.App{
		Flags: []cli.Flag {
			&cli.BoolFlag{ // ipv4 flag
				Name: "4",
				Value: true,
				Usage: "Use IPv4 only",
				Destination: &ipv4,
				Required: false,
			},
			&cli.BoolFlag{ // ipv6 flag
				Name: "6",
				Value: false,  // By default ipv6 is not set
				Usage: "Use IPv6 only",
				Destination: &ipv6,
			},
		},
		Action: func(c *cli.Context) error {
			if ipv6 == true {
				ipv4 = false // overides behaviour
			}else{

			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}