package version

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"version/version/utils"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

// IncMajor increments major version, zeroing minor and patch versions
func (v *Version) IncMajor() {
	v.Major++
	v.Minor = 0
	v.Patch = 0
}

// IncMinor increments minor version, zeroing patch version
func (v *Version) IncMinor() {
	v.Minor++
	v.Patch = 0
}

// IncPatch increments patch version
func (v *Version) IncPatch() {
	v.Patch++
}

// Stringify returns version object as formatted string
func (v *Version) Stringify() string {
	return fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Patch)
}

func Update(args *Args) (string, error) {
	var err error
	// Default version
	var current_version = "0.0.0"

	if (args.IncrementMajor || args.IncrementMinor || args.IncrementPatch) && args.IncrementOptions != "" {
		return "", errors.New("you can't use a targeted version incrementation (-M, -m, -p) and batch incrementation (-b) at the same time")
	}

	if (args.IncrementMajor && args.IncrementMinor) || (args.IncrementMinor && args.IncrementPatch) || (args.IncrementMajor && args.IncrementPatch) {
		return "", errors.New("targeted version incrementation (-M, -m, -p) resets lower version tier (-M '1.24.6' => '2.0.0'")
	}

	// Check if '-o' file exists
	if utils.FileExists(args.FilePath) {
		current_version, err = utils.ReadFileString(args.FilePath)
		if err != nil {
			return "", err
		}
	}

	if args.IncrementMajor {
		args.IncrementOptions = "M"
	}
	if args.IncrementMinor {
		args.IncrementOptions = "m"
	}
	if args.IncrementPatch {
		args.IncrementOptions = "p"
	}

	updated_version, err := changeVersion(current_version, args.IncrementOptions)
	if err != nil {
		return "", err
	}

	err = utils.WriteFile(args.FilePath, updated_version)
	if err != nil {
		return "", err
	}

	return updated_version, nil
}

// changeVersion increases the current version or returns the default one
func changeVersion(current_version, options string) (string, error) {
	if !isValidVersion(current_version) {
		return "", errors.New("version should match '0.0.0' pattern")
	}

	v := Version{}

	parts := strings.Split(current_version, ".")
	for i := 0; i < 3; i++ {
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			return "", err
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

	// Change version according to '-b' argument
	for _, c := range strings.Split(options, "") {
		switch c {
		case "M":
			v.IncMajor()
		case "m":
			v.IncMinor()
		case "p":
			v.IncPatch()
		default:
			return "", fmt.Errorf("unknown incrementation flag '%v'", c)
		}
	}

	return v.Stringify(), nil
}

// isValidVersion checks the passed string for compliance with the version template
func isValidVersion(v string) bool {
	rule := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	return rule.MatchString(v)
}
