package main

import (
	"fmt"
	"github.com/denisacostaq/glanguage/src"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
)

func init() {
	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// curl --request POST --header "Content-Type: application/json" http://localhost:8080/word --data '{"english-word": "apple"}'
// curl --request POST --header "Content-Type: application/json" http://localhost:8080/sentence --data '{"english-sentence": "hello world"}'
// curl http://localhost:8080/history
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
					log.WithField("port", portArg).Errorln("unable to get port as a number")
					return err
				}
				if port < 1 || port > 65535 {
					return fmt.Errorf("the number %d is not in the valid range from 1 to 65535", port)
				}
				if port <= 1023 {
					log.WithField("port", port).Warningln("this port is in the reserved system range")
				}
			}
			s := src.NewServer(uint16(port))
			return s.Start()
		},
		Authors: []*cli.Author{&cli.Author{
			Name: "Alvaro Denis",
			Email: "denisacostaq@gmail.co"},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.WithError(err).Errorln("unable to start the application")
	}
}