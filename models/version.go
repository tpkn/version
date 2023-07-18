package models

import (
	"fmt"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

// IncMajor increments major version
func (v *Version) IncMajor() {
	v.Major++
}

// IncMinor increments minor version
func (v *Version) IncMinor() {
	v.Minor++
}

// IncPatch increments patch version
func (v *Version) IncPatch() {
	v.Patch++
}

// Stringify ...
func (v *Version) Stringify() string {
	return fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Patch)
}
