package main

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/kryptn/confg/parser"
	"io/ioutil"
	"log"
)

var logLevel = logErr

func main() {
	log.Print("Starting")
	settings, err := getSettings()
	if err != nil {
		log.Fatal(err)
	}
	logLevel = settings.verbosity
	log.Print("Got settings")

	parsed, err := parser.ParsedFromFile(settings.inputFile)
	if err != nil {
		log.Fatal(err)
	}
	if parsed != nil {
		log.Printf("parsed backends: %+v", parsed.Backends)
		log.Printf("parsed keys: %+v", parsed.Keys)
	}

	err = parsed.Parse()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Parsed settings")

	confg, err := confgFromParsed(parsed)
	if confg != nil {
		log.Printf("confg rendered: %+v", confg.Rendered)
	}

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(confg.Rendered); err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(settings.outputFile, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(buf)

}
