# Tajnik 

* Simple Cli tool for storing all your passwords.

* Made as an assignment for Computer Security course @FER

## Installation

* [Go](https://golang.org/cmd/go/) cli tool should be installed (run `go` to check)
  If you have `go` you can skip this part
  The tool can be found and installed using your operating system's package manager or just
  position yourself inside this repo and run:
  ```
  make go-install
  # refresh your shell env variables with:
  source ~/.(your_shell_rc_file)

  ```
  This script should install the tool and set needed env variables for all go-based tools to work.

* Run 

```
make
```

## Build without install

* Run 
```
make build
```
this will create an executable in `./bin` folder


## Documentation
```
tajnik --help
```
