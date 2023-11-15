package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	
	"version/models"
	"version/utils"
)

var version = "0.0.0"
var help = fmt.Sprintf(`version (v%v) | https://tpkn.me

Go cli tool for auto incrementing app version

Usage:
  version -v "Mmp" [-options]

Options:
  -v          Increment options: [M]ajor, [m]inor, and [p]atch (default: "p")
  -o          File for storing the version number (default: "./.version")
  -h, --help  Help

Examples:
  -v "p"
  0.0.0 -> 0.0.1

  -v "Mmp"
  0.0.0 -> 1.1.1

  -v "mMp"
  0.0.0 -> 1.0.1
`, version)

func main() {
	cli := models.CLI{}
	flag.StringVar(&cli.IncrementOptions, "v", "p", "Increment options: [M]ajor, [m]inor, and [p]atch")
	flag.StringVar(&cli.FilePath, "o", "./.version", "File for storing the version number")
	flag.BoolVar(&cli.Help, "h", false, "Help")
	flag.BoolVar(&cli.Help, "help", false, "Alias for --help")
	flag.Parse()
	
	log.SetFlags(0)
	log.SetPrefix("Error: ")
	
	if cli.Help {
		fmt.Println(help)
		os.Exit(0)
	}
	
	// Default version
	current_version := "0.0.0"
	
	// Check if '-o' file exists
	if utils.FileExists(cli.FilePath) {
		current_version = utils.ReadFile(cli.FilePath)
	}
	
	new_version := changeVersion(current_version, cli.IncrementOptions)
	utils.WriteFile(cli.FilePath, new_version)
}

// changeVersion increases the current version or returns the default one
func changeVersion(current_version, inc_options string) string {
	if !isValidVersion(current_version) {
		log.Fatalln("version should match '0.0.0' pattern")
	}
	
	v := models.Version{}
	
	parts := strings.Split(current_version, ".")
	for i := 0; i < 3; i++ {
		n, err := strconv.ParseUint(parts[i], 10, 32)
		if err != nil {
			log.Fatalln(err)
		}
		
		switch i {
		case 0:
			v.Major = n
		case 1:
			v.Minor = n
		case 2:
			v.Patch = n
		}
	}
	
	// Change version according to '-v' argument
	options_list := strings.Split(inc_options, "")
	for _, c := range options_list {
		switch c {
		case "M":
			v.IncMajor()
		case "m":
			v.IncMinor()
		case "p":
			v.IncPatch()
		}
	}
	
	return v.Stringify()
}

// isValidVersion checks the passed string for compliance with the version template
func isValidVersion(v string) bool {
	rule := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	return rule.MatchString(v)
}
