package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

type CliSettings struct {
	inputFile  string
	outputFile string

	dryRun    bool
	verbosity int
}

func (c *CliSettings) inputFileFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:        "file, f",
		Value:       "confg.toml",
		Usage:       "File location to use for config",
		EnvVar:      "CONFG_FILE",
		Destination: &c.inputFile,
	}
}

func (c *CliSettings) outputFileFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:        "output, o",
		Value:       "settings.toml",
		Usage:       "File location to save rendered settings",
		EnvVar:      "CONFG_OUTPUT",
		Destination: &c.outputFile,
	}
}

func (c *CliSettings) validateSettings(con *cli.Context) error {
	// Meant to validate the input data since
	// we're loading by reference
	return nil
}

func getSettings() (*CliSettings, error) {
	settings := &CliSettings{}
	app := cli.NewApp()
	app.Name = "Confg"
	app.Action = settings.validateSettings
	app.Flags = []cli.Flag{
		settings.inputFileFlag(),
		settings.outputFileFlag(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	return settings, nil
}
