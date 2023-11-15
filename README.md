# Version
Go cli tool for auto incrementing app version


## Usage

```shell
version -v "Mmp" [-options]
```


## Options

```text
-v      Increment options: [M]ajor, [m]inor, and [p]atch (default: "p")
-o      File for storing the version number (default: "./.version")
-p      Print version to Stdout (default: false)
--help  Help
```


## Examples
```
-v "M"
0.0.0 -> 1.0.0


-v "m"
0.0.0 -> 0.1.0


-v "p"
0.0.0 -> 0.0.1


-v "Mmp"
0.0.0 -> 1.1.1


-v "mMp"
0.0.0 -> 1.0.1


-v "mmmmp"
0.0.0 -> 0.4.1


-v "Mpppppppppppppp"
0.0.0 -> 1.0.14


-v "ppppppppppppppM"
0.0.0 -> 1.0.0


-v "MMMMMMmmmmmmmmmmppppppppppppp"
0.0.0 -> 7.5.13
```




