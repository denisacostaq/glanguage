package main

import (
	"fmt"
	"github.com/denisacostaq/glanguage/src"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	var portArg string
	app := &cli.App{
		Name: "glanguage",
		HelpName: "glanguage",
		Usage: "Translate English to the Gophers's language",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Listen port",
				Destination: &portArg,
				EnvVars: []string{"GLANGUAGE_PORT"},
			},
		},
		Action: func(c *cli.Context) error {
			port := 8080
			if len(portArg) > 0 {
				var err error
				if port, err = strconv.Atoi(portArg); err != nil {
					return err
				}
				if port < 1 || port > 65535 {
					return fmt.Errorf("the number %d is not in the valid range from 1 to 65535", port)
				}
				if port <= 1023 {
					fmt.Println("Warning, system port range")
				}
			}
			s := src.NewServer(uint16(port))
			return s.Start()
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}