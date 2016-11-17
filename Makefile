# Name of the binary
BINARY=pastestuff

# Build Information
VERSION=0.0.1
BUILD=$(shell date -u '+%Y-%m-%d')

# Some Paths
CWD=$(shell pwd)
GOAPP=$(CWD)/server
SERVER=github.com/sstutz/pastestuff/server
ELMAPP=$(CWD)/client
DISTDIR=$(CWD)/dist
RELEASE=$(BINARY).$(BUILD).tar.gz

# interpolate variable values
LDFLAGS=-ldflags "-X $(SERVER)/helpers.Version=$(VERSION) -X $(SERVER)/helpers.Build=$(BUILD) -X $(SERVER)/helpers.AppName=$(BINARY)"

.PHONY: client server

all: client server release checksums

deps:
	go get -u github.com/GeertJohan/go.rice
	go get -u github.com/GeertJohan/go.rice/rice
	go get -u github.com/wellington/wellington/wt
	cd $(GOAPP) && glide install
	cd $(ELMAPP) && elm-package install -y

server:
	cd $(GOAPP)/assets/ && rice embed-go
	GOOS=linux	GOARCH=amd64	go build $(LDFLAGS) -o $(DISTDIR)/linux/amd64/$(BINARY)		$(SERVER)/pastestuff

release:
	tar zcvf $(DISTDIR)/linux/amd64/$(RELEASE)	LICENSE README.md -C $(DISTDIR)/linux/amd64/	$(BINARY)

client:
	cd $(ELMAPP) && mkdir -p dist/js  && elm-make src/Main.elm --output dist/js/$(BINARY).js
	cd $(ELMAPP) && mkdir -p dist/css && wt -b dist/css compile assets/scss/style.scss
	cp -r client/assets/images/ client/dist/

checksums:
	cd $(DISTDIR)/linux/amd64/	&& sha256sum $(RELEASE) > $(BINARY).sha256

clean:
	if [ -d $(ELMAPP)/dist ]; then rm -rf $(ELMAPP)/dist; fi
	if [ -f $(GOAPP)/assets/rice-box.go ]; then rm -rf $(GOAPP)/assets/rice-box.go; fi
	if [ -d $(DISTDIR) ]; then rm -rf $(DISTDIR); fi

todo:
	grep --color=always -rwn "TODO" $(GOAPP) $(ELMAPP)/src

fixme:
	grep --color=always -rwn "FIXME" $(GOAPP) $(ELMAPP)/src
