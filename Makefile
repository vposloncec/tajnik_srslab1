# Install go

all: install

install-go:
	wget -q -O - https://git.io/vQhTU | bash

install: check-go
	cd ./tajnik && go install
	@echo ""
	@echo "Cli tool successfully installed, run 'tajnik help' for more information"
	@echo "If command 'tajnik' is not found, make sure you have GOBIN in path"

build: 
	go build -o ./bin/tajnik ./tajnik/main.go

check-go:
ifeq (, $(shell which go 2> /dev/null))
	$(error No go cli tool found in PATH $(NEWLINE)\
		consider doing 'make install-go' or 'apt-get install golang-go' $(NEWLINE)\
		NOTE: doing 'make install-go' will automatically set needed env variables)
else
	@echo "go compiler found, continuing"
endif

define NEWLINE


endef
