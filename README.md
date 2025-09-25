<p align="center">
   <img width="150" src="icon.png" alt="" align="center">
</p>
<p align="center">
   A tool for easy version incrementation.
</p>



Upon the first launch in a new project, the application creates a `.version` file (any filename can be specified using
the `-o` flag) in which the version `0.0.0` is recorded. If the project already contains a version file, it will update
the version number in it.

## Usage

```text
version [-M | -m | -p | -b] [-o]
```

## Options

```text
-M             Increment major version
-m             Increment minor version
-p             Increment patch version
-b             Batch incrementation options: (M)ajor, (m)inor, and (p)atch
-o             File for storing the version number (default: "./.version")
-h, --help     Help
-v, --version  Version
```

## Examples

```shell
version -M
# 0.0.0 -> 1.0.0

version -m
# 0.0.0 -> 0.1.0

version -p
# 0.0.0 -> 0.0.1

version -b "M"
# 0.0.0 -> 1.0.0

version -b "m"
# 0.0.0 -> 0.1.0

version -b "p"
# 0.0.0 -> 0.0.1

version -b "Mmp"
# 0.0.0 -> 1.1.1

version -b "mMp"
# 0.0.0 -> 1.0.1

version -b "mmmmp"
# 0.0.0 -> 0.4.1

version -b "Mpppppppppppppp"
# 0.0.0 -> 1.0.14

version -b "ppppppppppppppM"
# 0.0.0 -> 1.0.0

version -b "MMMMMMmmmmmmmmmmppppppppppppp"
# 0.0.0 -> 7.5.13
```




