GO=go
TEST_DIR=$(realpath test)
BIN_DIR=bin
APP=$(BIN_DIR)/app
STATIC_APP=$(BIN_DIR)/app_static
SOURCES=$(shell find . -name '*.go' -not -name '*_test.go' -not -path "./test/*" -not -path './scripts/*')

all: $(APP)

static: $(STATIC_APP)

format:
	find . -name '*.go' -not -path "./.cache/*" | xargs -n1 $(GO) fmt

check: format
	git diff
	git diff-index --quiet HEAD

lint:
	golangci-lint run --skip-dirs=test --deadline 3m0s

test: $(TEST_DIR)
	TEST_PATH=$(TEST_DIR) $(GO) test -cover ./... -coverprofile .coverage.txt
	cat .coverage.txt | grep "/pkg\|mode:" > .coverage.pkg
	$(GO) tool cover -func .coverage.pkg

clean:
	rm -rf $(BIN_DIR) $(PARSER) $(VALIDATOR)

$(BIN_DIR)/%: cmd/% $(SOURCES)
	$(GO) build -o $@ $</*.go
	strip -s $@

$(BIN_DIR)/%_static: cmd/% $(SOURCES)
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -installsuffix cgo -o $@ $</*.go
	strip -s $@

$(BIN_DIR):
	mkdir -p $@

.PHONY: all static format check lint clean