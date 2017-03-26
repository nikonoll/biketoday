package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	var city string
	var start string
	var destination string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "c",
			Value:       "Munich, Germany",
			Usage:       "current city",
			Destination: &city,
		},
		cli.StringFlag{
			Name:        "s",
			Value:       "Goethestr 2",
			Usage:       "current city",
			Destination: &start,
		},
		cli.StringFlag{
			Name:        "d",
			Value:       "Arcisstr 1",
			Usage:       "Destination adress",
			Destination: &destination,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			city = c.Args()[0]
			start = c.Args()[0]
			destination = c.Args()[0]
		}
		// geovalide adresses

		// call forecast.io on citiiy

		// display weather with wegoii?

		// call maps api for distance and travel time

		// google compare to alternative travel time (bus etc.)

		fmt.Println("These args were passed", start, destination, city)
		return nil
	}

	app.Run(os.Args)
}
