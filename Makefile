BIN := tmpl
GO := go
TAGS := yaml
VERSION := $(shell version.sh)
PREFIX := /usr/local/bin

.PHONY: all clean install uninstall

all: build

build: $(BIN)

$(BIN): $(wildcard *.go)
	@$(GO) build -ldflags "-X main.Version=$(VERSION)" -tags $(TAGS) -o $(BIN) .

clean:
	@rm -f $(BIN)

install: build
	@install -m 700 $(BIN) $(PREFIX)/$(BIN)

uninstall:
	@rm -f $(PREFIX)/$(BIN)
