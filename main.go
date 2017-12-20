package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli"
	"github.com/valyala/gorpc"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gochain"
	app.Usage = ""

	app.Commands = []cli.Command{
		{
			Name:    "daemon",
			Aliases: []string{},
			Usage:   "start the daemon",
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
		{
			Name:    "keys",
			Aliases: []string{},
			Usage:   "generate a new keypair",
			Action:  generate_keys,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "view, v",
					Usage: "View your existing public key",
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
			action := Action{}
			proto.Unmarshal(request.([]byte), &action) // returns an err, but meh
			switch action.Type {
			case Action_SEND:
				log.Printf("Obtained request %v from the client %s\n", action, clientAddr)
			case Action_VIEW:
				log.Printf("Obtained request %v from the client %s\n", action, clientAddr)
			case Action_LIST:
				log.Printf("Obtained request %v from the client %s\n", action, clientAddr)
			case Action_BALANCE:
				log.Printf("Obtained request %v from the client %s\n", action, clientAddr)
			default:
				return errors.New("Invalid action")
			}
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

func generate_keys(c *cli.Context) error {
	keyfile := os.Getenv("HOME") + "/.gochain"
	if c.Bool("view") {
		data, err := ioutil.ReadFile(keyfile)
		keys := rsa.PrivateKey{}
		if err == nil {
			err = json.Unmarshal(data, &keys)
			key_bytes := keys.PublicKey.N.Bytes()
			fmt.Printf("%v\n", keys.PublicKey.E)
			fmt.Printf("%v\n", base64.StdEncoding.EncodeToString(key_bytes))
		}
		return err
	}

	if _, err := os.Stat(keyfile); err == nil {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("A keyfile already exists. Are you sure? (y/n): ")
		text, _ := reader.ReadString('\n')
		if text != "y" {
			return nil
		}
	}

	keys, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	js, _ := json.Marshal(keys)
	err = ioutil.WriteFile(keyfile, js, 0600)
	return err
}
