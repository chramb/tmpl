BIN := tmpl
GO := go
TAGS := yaml

.PHONY: all clean

all: build

build: $(BIN)

$(BIN): $(wildcard *.go)
	@$(GO) build -tags $(TAGS) -o $(BIN) .

clean:
	@rm -f tmpl
