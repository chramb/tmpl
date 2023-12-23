BIN := tmpl
GO := go
TAGS := yaml
VERSION := $(shell version.sh)

.PHONY: all clean

all: build

build: $(BIN)

$(BIN): $(wildcard *.go)
	@$(GO) build -ldflags "-X main.Version=$(VERSION)" -tags $(TAGS) -o $(BIN) .

clean:
	@rm -f $(BIN)
