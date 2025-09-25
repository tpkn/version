package main

const Help = `                                          /▒▒
                                         |__/
  /▒▒    /▒▒ /▒▒▒▒▒▒   /▒▒▒▒▒▒   /▒▒▒▒▒▒▒ /▒▒  /▒▒▒▒▒▒  /▒▒▒▒▒▒▒
 |  ▒▒  /▒▒//▒▒__  ▒▒ /▒▒__  ▒▒ /▒▒_____/| ▒▒ /▒▒__  ▒▒| ▒▒__  ▒▒
  \  ▒▒/▒▒/| ▒▒▒▒▒▒▒▒| ▒▒  \__/|  ▒▒▒▒▒▒ | ▒▒| ▒▒  \ ▒▒| ▒▒  \ ▒▒
   \  ▒▒▒/ | ▒▒_____/| ▒▒       \____  ▒▒| ▒▒| ▒▒  | ▒▒| ▒▒  | ▒▒
 /▒▒\  ▒/  |  ▒▒▒▒▒▒▒| ▒▒       /▒▒▒▒▒▒▒/| ▒▒|  ▒▒▒▒▒▒/| ▒▒  | ▒▒ v%v
|__/ \_/    \_______/|__/      |_______/ |__/ \______/ |__/  |__/ https://tpkn.me

A tool for easy version incrementation.

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
  version -M
  # 0.0.0 -> 1.0.0

  version -m
  # 0.0.0 -> 0.1.0

  version -p
  # 0.0.0 -> 0.0.1

  version -b "p"
  # 0.0.0 -> 0.0.1

  version -b "Mmp"
  # 0.0.0 -> 1.1.1

  version -b "mMp"
  # 0.0.0 -> 1.0.1
`
