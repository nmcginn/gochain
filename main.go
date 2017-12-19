package main

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/valyala/gorpc"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "RPC Test"
	app.Usage = ""

	app.Commands = []cli.Command{{
		Name:    "serve",
		Aliases: []string{},
		Usage:   "start the RPC server",
		Action:  serve,
		Flags:   []cli.Flag{},
	},
		{
			Name:    "run",
			Aliases: []string{},
			Usage:   "run an RPC client",
			Action:  run,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "data, d",
					Value: "foobar",
					Usage: "The data to send to the RPC server",
				},
			},
		}}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err.Error())
		os.Exit(1)
	}
}

func serve(c *cli.Context) error {
	s := &gorpc.Server{
		Addr: ":12345",
		Handler: func(clientAddr string, request interface{}) interface{} {
			log.Printf("Obtained request %v from the client %s\n", request, clientAddr)
			return request
		},
	}
	err := s.Serve()
	return err
}

func run(c *cli.Context) error {
	client := &gorpc.Client{
		Addr: "127.0.0.1:12345",
	}
	client.Start()

	resp, err := client.Call(c.String("data"))
	fmt.Printf("%s", resp.(string))
	return err
}
