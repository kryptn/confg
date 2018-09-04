package main

import (
	"flag"
	"strings"
)

type InputFiles []string

func (i *InputFiles) String() string {
	return strings.Join(*i, ", ")
}

func (i *InputFiles) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type Settings struct {
	inputFiles InputFiles
	outputFile string

	priorityDelta int

	verbose     bool
	includeMeta bool
}

const (
	inputFileFlagsUsage = "Input file location (repeatable)"

	outputFileFlagsUsage   = "Output file location"
	outputFileFlagsDefault = "settings.toml"

	verboseFlagUsage = "add to stdout operation logs"

	metaFlagUsage = "include toml operation logs in a top-level `confg-meta` map"

	shorthand = " (shorthand)"
)

func (s *Settings) DeclareFlags() {
	flag.Var(&s.inputFiles, "file", inputFileFlagsUsage)
	flag.Var(&s.inputFiles, "f", inputFileFlagsUsage+shorthand)

	flag.StringVar(&s.outputFile, "out", outputFileFlagsDefault, outputFileFlagsUsage)
	flag.StringVar(&s.outputFile, "o", outputFileFlagsDefault, outputFileFlagsUsage+shorthand)

	flag.BoolVar(&s.verbose, "dry-run", false, verboseFlagUsage)

	flag.BoolVar(&s.includeMeta, "meta", false, metaFlagUsage)
}

func GetSettings() (*Settings, error) {
	settings := Settings{}
	settings.DeclareFlags()
	flag.Parse()

	return &settings, nil
}
