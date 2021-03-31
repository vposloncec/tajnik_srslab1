# Tajnik 🔐 

* Simple Cli tool for storing all your passwords.

* Made as an assignment for Computer Security course @FER

## Installation
Clone this repo with `git clone git@github.com:vposloncec/srs-lab1.git`

* [**Go**](https://golang.org/cmd/go/) cli tool should be installed (run `go` to check). 

  **If you have `go` you can skip this part.**

  The tool can be found and installed using your operating system's package manager or just
  position yourself inside this repo and run:
  ```shell
  make install-go
  # refresh your shell env variables with:
  source ~/.(your_shell_rc_file)

  ```
  This script should install the tool and set needed env variables for all go-based tools to work.

#### Install using 
```shell
make
```

## Build without install
```shell
make build
```
this will create an executable in `./bin` folder


## Documentation
`tajnik help`
```
Simple password (credentials) manager

Usage:
  tajnik [command]

Available Commands:
  get         Get credentials for desired address
  help        Help about any command
  init        Initializes tajnik using the provided master password
  put         Put (add) password to the manager's database

Flags:
  -h, --help           help for tajnik
      --mfile string   Master file where data is stored(default is $HOME/.tajnik/master_file)

Use "tajnik [command] --help" for more information about a command.
```
