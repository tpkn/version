package version

type Args struct {
	IncrementMajor   bool
	IncrementMinor   bool
	IncrementPatch   bool
	IncrementOptions string
	FilePath         string
	Help             bool
	Version          bool
}
