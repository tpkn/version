package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"version/version"
)

var v = "0.0.0"
var help = fmt.Sprintf(Help, v)

func main() {
	args := version.Args{}
	flag.BoolVar(&args.IncrementMajor, "M", false, "Increment major version")
	flag.BoolVar(&args.IncrementMinor, "m", false, "Increment minor version")
	flag.BoolVar(&args.IncrementPatch, "p", false, "Increment patch version")
	flag.StringVar(&args.IncrementOptions, "b", "", "Batch increment options: (M)ajor, (m)inor, and (p)atch")
	flag.StringVar(&args.FilePath, "o", "./.version", "File for storing the version number")
	flag.BoolVar(&args.Help, "h", false, "Help")
	flag.BoolVar(&args.Help, "help", false, "Alias for -h")
	flag.BoolVar(&args.Version, "v", false, "Version")
	flag.BoolVar(&args.Version, "version", false, "Alias for -v")
	flag.Parse()

	log.SetFlags(0)
	log.SetPrefix("version: ")

	if args.Help {
		fmt.Println(help)
		os.Exit(0)
	}

	if args.Version {
		fmt.Println(v)
		os.Exit(0)
	}

	_, err := version.Update(&args)
	if err != nil {
		log.Fatalln(err)
	}
}
