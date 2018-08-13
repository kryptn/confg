package main

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/kryptn/confg/containers"
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

	confg := containers.Confg{
		Backends: parsed.Backends,
		Keys:     parsed.Keys,
	}

	log.Printf("made confg")

	ok, errs := confg.Validate()
	if errs != nil {
		for _, err := range errs {
			log.Print(err)
		}
	}
	if !ok {
		log.Fatal("Due to above errors, cannot continue")
	}

	GatherAllKeys(confg)

	err = confg.ReduceKeys()
	if err != nil {
		log.Printf("error when reducing %v", err)
	}

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(confg.Reduced); err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(settings.outputFile, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\nconfig output: \n%s", buf)

}
