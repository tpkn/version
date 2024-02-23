package models

import (
	"fmt"
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

// Stringify ...
func (v *Version) Stringify() string {
	return fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Patch)
}
