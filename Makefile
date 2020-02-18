export PATH := $(GOPATH)/bin:$(PATH)

all: fmt build

build: nslookgo

nslookgo:
	go build -o bin/nslookgo

fmt:
	go fmt .

clean:
	rm -rf bin/*