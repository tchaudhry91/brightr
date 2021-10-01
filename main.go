package main

import (
	"log"
	"os"

	"github.com/tchaudhry91/brightr/backlight"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "brighr",
		Usage: "CLI Wrapper to change screen brightness values. This is an experimental tool. Please use at your own risk!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "device",
				Aliases: []string{"d"},
				Value:   "",
				Usage:   "Specify the device to operate on. Default: First Device",
			},
			&cli.IntFlag{
				Name:    "amount",
				Aliases: []string{"a"},
				Value:   10,
				Usage:   "Specify the amount to increase or decrease. Default: 10",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "backlights",
				Usage: "show eligible backglights",
				Action: func(c *cli.Context) error {
					backlights, err := backlight.GetBacklights()
					if err != nil {
						return err
					}
					println("Available Backlights:")
					for _, b := range backlights {
						println(b)
					}
					return nil
				},
			},
			{
				Name:    "up",
				Aliases: []string{"u"},
				Usage:   "Increase brightness",
				Action: func(c *cli.Context) error {
					backlights, err := backlight.GetBacklights()
					if err != nil {
						return err
					}
					device := backlights[0]
					if c.String("device") != "" {
						device = c.String("device")
					}
					amount := c.Int("amount")
					return backlight.IncreaseBrightness(device, amount)
				},
			},
			{
				Name:    "down",
				Aliases: []string{"d"},
				Usage:   "Decrease brightness",
				Action: func(c *cli.Context) error {
					backlights, err := backlight.GetBacklights()
					if err != nil {
						return err
					}
					device := backlights[0]
					if c.String("device") != "" {
						device = c.String("device")
					}
					amount := c.Int("amount")
					return backlight.DecreaseBrightness(device, amount)
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
