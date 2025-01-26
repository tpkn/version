package main

const Help = `version (v%v) | https://tpkn.me

Tool for automatic version incrementation.

Usage:
  version [-M | -m | -p | -b] [-o]

Options:
  -M             Increment major version
  -m             Increment minor version
  -p             Increment patch version
  -b             Batch incrementation options: (M)ajor, (m)inor, and (p)atch
  -o             File for storing the version number (default: "./.version")
  -h, --help     Help
  -v, --version  Version

Examples:
  -M
  0.0.0 -> 1.0.0

  -m
  0.0.0 -> 0.1.0

  -p
  0.0.0 -> 0.0.1

  -b "p"
  0.0.0 -> 0.0.1

  -b "Mmp"
  0.0.0 -> 1.1.1

  -b "mMp"
  0.0.0 -> 1.0.1
`
