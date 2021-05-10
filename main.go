package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	
	"version/models"
	"version/utils"
)

var env = "development"
var version = "0.0.0"
var description = fmt.Sprintf(`Version, v%v | https://tpkn.me`, version)

func main() {
	cli := models.CLI{}
	flag.StringVar(&cli.About,"about", "", description)
	flag.StringVar(&cli.FilePath,"o", "./.version", "Path to file where you want to store your app version")
	flag.StringVar(&cli.Options,"v", "p", "Increment settings")
	flag.BoolVar(&cli.Verbose, "vb", true, "Verbose output")
	flag.Parse()
	
	if env == "development" {
		cli.FilePath = "./version"
		cli.Options = "pppmmM"
	}
	
	// Default version
	current_version := "0.0.0"
	
	// Check if '-o' file exists
	if utils.FileExists(cli.FilePath) {
		current_version = utils.ReadFile(cli.FilePath)
	}
	
	new_version := changeVersion(current_version, cli.Options)
	utils.WriteFile(cli.FilePath, new_version)
	
	// Print version to STDOUT
	if cli.Verbose {
		fmt.Print(new_version)
	}
}

func changeVersion(current_version, options string) string {
	if !isValidVersion(current_version) {
		current_version = "0.0.0"
	}
	
	v := models.Version{0, 0, 0}
	
	parts := strings.Split(string(current_version), ".")
	for i := 0; i < 3; i++ {
		n, _ := strconv.Atoi(parts[i])
		switch i {
		case 0: v.Major = n
		case 1: v.Minor = n
		case 2: v.Patch = n
		}
	}
	
	// Change version according to '-v' argument
	options_list := strings.Split(options, "")
	for _, c := range options_list {
		switch c {
		case "M": v.IncMajor()
		case "m": v.IncMinor()
		case "p": v.IncPatch()
		}
	}
	
	// Stringify version
	stringified := fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Patch)
	
	return stringified
}

func isValidVersion(v string) bool {
	rule := regexp.MustCompile(`^\d{1,}\.\d{1,}\.\d{1,}$`)
	return rule.MatchString(v)
}
