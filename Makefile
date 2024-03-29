.PHONY: all install clean build view
VERSION=v0.1.0

all: clean build

install:
	go get -u github.com/michalq/go-swagger/cmd/swagger

build:
	$(GOPATH)/bin/swagger generate client -f ./swagger.json -A airly-api-client

clean:
	rm -rf restapi
	rm -rf models

view:
	$(GOPATH)/bin/swagger serve ./swagger.json

release:
	$(GOPATH)/bin/git-chglog --next-tag $(VERSION) -o CHANGELOG.md
	git add CHANGELOG.md
	git commit -a -m "Release $(VERSION)"
	git tag $(VERSION)
