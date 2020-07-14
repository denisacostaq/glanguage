package main

import (
	"fmt"
	"github.com/denisacostaq/glanguage/src"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

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
	}
	err := app.Run(os.Args)
	if err != nil {
		log.WithError(err).Errorln("unable to start the application")
	}
}