
build:
	CGO_ENABLED=0 go build -a -ldflags '-s' .

install:
	go install -v .

setup-root-gotools:
	sudo GOPATH=/tmp/tmpGOPATH go get -u -v code.google.com/p/go.tools/cmd/cover
	sudo GOPATH=/tmp/tmpGOPATH go get -u -v code.google.com/p/go.tools/cmd/vet
	sudo rm -rf /tmp/tmpGOPATH

setup-gotools:
	go get -u -v github.com/golang/lint/golint

test-style:
	go vet ./...
	-golint ./...

test: test-style
	go test -v -cover ./...
