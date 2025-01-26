<p align="center">
   <img width="150" src="icon.png" alt="" align="center">
</p>
<p align="center">
   Tool for auto incrementation your app version.
</p>



## Usage

```shell
version [-M | -m | -p | -b] [-o]
```


## Options

```text
-M             Increment major version
-m             Increment minor version
-p             Increment patch version
-b             Batch increment options: (M)ajor, (m)inor, and (p)atch
-o             File for storing the version number (default: "./.version")
-h, --help     Help
-v, --version  Version
```


## Examples
```
-M
0.0.0 -> 1.0.0

-m
0.0.0 -> 0.1.0

-p
0.0.0 -> 0.0.1

-b "M"
0.0.0 -> 1.0.0

-b "m"
0.0.0 -> 0.1.0

-b "p"
0.0.0 -> 0.0.1

-b "Mmp"
0.0.0 -> 1.1.1

-b "mMp"
0.0.0 -> 1.0.1

-b "mmmmp"
0.0.0 -> 0.4.1

-b "Mpppppppppppppp"
0.0.0 -> 1.0.14

-b "ppppppppppppppM"
0.0.0 -> 1.0.0

-b "MMMMMMmmmmmmmmmmppppppppppppp"
0.0.0 -> 7.5.13
```




