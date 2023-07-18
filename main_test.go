package main

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
)

func TestChangeVersion(t *testing.T) {
	type step struct {
		label   string
		version string
		must_be string
	}
	
	steps := []step{
		{"Empty", "", "0.0.0"},
		{"Mmp", "Mmp", "1.1.1"},
		{"mmmmp", "mmmmp", "0.4.1"},
		{"Mpppppppppppppp", "Mpppppppppppppp", "1.0.14"},
		{"MMMMMMmmmmmMmmmmmppppppppppppp", "MMMMMMmmmmmMmmmmmppppppppppppp", "7.10.13"},
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
	type step struct {
		label   string
		version string
		must_be bool
	}
	
	steps := []step{
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
