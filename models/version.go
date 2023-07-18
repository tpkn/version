package models

import (
	"fmt"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v *Version) IncMajor() {
	v.Major++
}

func (v *Version) IncMinor() {
	v.Minor++
}

func (v *Version) IncPatch() {
	v.Patch++
}

func (v *Version) Stringer() string {
	return fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Patch)
}
