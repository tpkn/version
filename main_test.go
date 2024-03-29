package main

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
)

func TestChangeVersion(t *testing.T) {
	steps := []struct {
		label   string
		version string
		must_be string
	}{
		{"Empty", "", "0.0.0"},
		{"Only major", "MMMM", "4.0.0"},
		{"Only minor", "mm", "0.2.0"},
		{"Only patch", "p", "0.0.1"},
		{"Mmp", "Mmp", "1.1.1"},
		{"mMp", "mMp", "1.0.1"},
		{"mmmmp", "mmmmp", "0.4.1"},
		{"Mpppppppppppppp", "Mpppppppppppppp", "1.0.14"},
		{"ppppppppppppppM", "ppppppppppppppM", "1.0.0"},
		{"MMMMMMmmmmmMmmmmmppppppppppppp", "MMMMMMmmmmmMmmmmmppppppppppppp", "7.5.13"},
	}
	
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.TabIndent)
	
	for _, k := range steps {
		result := changeVersion("0.0.0", k.version)
		if result != k.must_be {
			t.Errorf("%v: %v != %v", k.label, result, k.must_be)
		}
		fmt.Fprintln(w, k.label, "\t", result, "\t")
	}
	w.Flush()
}

func TestIsValidVersion(t *testing.T) {
	steps := []struct {
		label   string
		version string
		must_be bool
	}{
		{"Valid version", "0.0.0", true},
		{"Huge but valid version", "123.767.990", true},
		{"Empty version", "", false},
		{"Invalid version 1", "1.2.3.4.5", false},
		{"Invalid version 2", "1-.0.0", false},
	}
	
	for _, k := range steps {
		result := isValidVersion(k.version)
		if result != k.must_be {
			t.Errorf("%v: %v != %v", k.label, result, k.must_be)
		} else {
			t.Logf("[OK] %v: %v", k.label, k.version)
		}
	}
}
