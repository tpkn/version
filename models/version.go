package models

type Version struct {
	Major int
	Minor int
	Patch int
}

func (n *Version) IncMajor() {
	n.Major++
}

func (n *Version) IncMinor() {
	n.Minor++
}

func (n *Version) IncPatch() {
	n.Patch++
}
