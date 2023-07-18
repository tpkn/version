# Version
Go cli tool for auto increment app version


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
-v "Mmp"
0.0.0 -> 1.1.1
```
```
-v "mmmmp"
0.0.0 -> 0.4.1
```
```
-v "Mpppppppppppppp"
0.0.0 -> 1.0.14
```
```
-v "MMMMMMmmmmmmmmmmppppppppppppp"
0.0.0 -> 6.10.13
```

