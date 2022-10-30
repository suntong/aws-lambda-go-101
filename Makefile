.PHONY: build clean deploy gomodgen

FILES := $(wildcard */*.go)
DIRS := $(patsubst %/.,%,$(wildcard */.))

#TARGS := $(patsubst %/, bin/%,$(dir $(FILES)))
TARGS := $(sort $(patsubst %/, bin/%,$(dir $(FILES))))

#$(info info: Go source directories are: $(DIRS)) # displays the contents for debugging
$(info info: Target files are: $(TARGS)) # displays the contents for debugging

build: $(TARGS)

bin/%: %/
#	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/*.go
	env GO111MODULE=on GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o $@ $*/*.go
	ls -lh $@

clean:
	rm -rf ./bin/*

deploy: clean build
	serverless deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
