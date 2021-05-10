# Version
Go cli tool for auto increment app version


- Creates a file with a number of current version of the application.
- Each run, the version increases according to the `-v` setting.



## Usage


```shell
version -v "Mmp" -o "./app.v"
```


### -v
**Type**: `string`   
**Default**: `p`   

Increment settings: `M`ajor; `m`inor and `p`atch

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



### -o
**Type**: `string`   
**Default**: `./.version`   

File for storing the version of your app



## Build example (for Windows)

```
for /f %%i in ('bin\version.exe') do set VERSION=%%i

set GOOS=windows
set GOARCH=amd64
go build -ldflags "-X main.version=%VERSION%" -o "build/app.exe"

```