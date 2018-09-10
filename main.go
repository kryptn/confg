package main

import (
	"github.com/kryptn/confg/containers"
	"github.com/kryptn/confg/gatherer"
	"github.com/kryptn/confg/outputter"
	"github.com/kryptn/confg/parser"
	"log"
)

func allConfgs(inputFiles []string) []*containers.Confg {
	var confgs []*containers.Confg
	for _, inputFile := range inputFiles {
		confg, err := parser.ConfgFromFile(inputFile)
		if err != nil {
			log.Printf("Error with [%s] input file: %v", inputFile, err)
		}
		confgs = append(confgs, confg)
	}
	return confgs
}

func main() {
	// get run settings from cli
	settings, err := GetSettings()
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("Settings: %+v", settings)

	// each input file is successively applied on top of the previous
	collected := (&containers.Confg{}).Overlay(allConfgs(settings.inputFiles)...)
	//log.Printf("collected: %+v", collected)

	// attempt to resolve each key
	gather := gatherer.NewGatherer(collected)
	resolved, err := gather.Resolve()
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("resolved: %+v", resolved)

	reduced, err := resolved.Reduce()
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("reduced: %+v", reduced)

	//for i, key := range reduced.Keys {
	//	fmt.Printf("\n%d: %+v\n", i, key)
	//}
	//
	//fmt.Print("\n")

	err = outputter.Output(settings.outputFile, reduced.Reduced)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%s\n", buf)

}
