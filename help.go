package main

const Help = `version (v%v) | https://tpkn.me

Go cli tool for auto incrementing app version

Usage:
  version -v "Mmp" [-options]

Options:
  -v          Increment options: [M]ajor, [m]inor, and [p]atch (default: "p")
  -o          File for storing the version number (default: "./.version")
  -h, --help  Help

Examples:
  -v "p"
  0.0.0 -> 0.0.1

  -v "Mmp"
  0.0.0 -> 1.1.1

  -v "mMp"
  0.0.0 -> 1.0.1
`
