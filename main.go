package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli"
	"github.com/valyala/gorpc"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "RPC Test"
	app.Usage = ""

	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{},
			Usage:   "start the RPC server",
			Action:  serve,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:   "port, p",
					Value:  12345,
					Usage:  "Server port to bind to",
					EnvVar: "PORT",
				},
			},
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
				cli.IntFlag{
					Name:   "port, p",
					Value:  12345,
					Usage:  "Server port to bind to",
					EnvVar: "PORT",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err.Error())
		os.Exit(1)
	}
}

func serve(c *cli.Context) error {
	s := &gorpc.Server{
		Addr: ":" + fmt.Sprintf("%d", c.Int("port")),
		Handler: func(clientAddr string, request interface{}) interface{} {
			cool_data := Test{}
			proto.Unmarshal(request.([]byte), &cool_data) // returns an err, but meh
			log.Printf("Obtained request %v from the client %s\n", cool_data.Msg, clientAddr)
			return request
		},
	}
	err := s.Serve()
	return err
}

func run(c *cli.Context) error {
	client := &gorpc.Client{
		Addr: "127.0.0.1:" + fmt.Sprintf("%d", c.Int("port")),
	}
	client.Start()

	message := Test{
		Msg: c.String("data"),
	}
	message_bytes, _ := proto.Marshal(&message)

	resp, err := client.Call(message_bytes)
	cool_data := Test{}
	proto.Unmarshal(resp.([]byte), &cool_data)
	fmt.Printf("%s", cool_data.Msg)
	return err
}
