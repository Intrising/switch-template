# makefile of intri event manager

CROSS = os-platform-xxx
GO_OPTS = GOARCH=arm64 GOOS=linux
NAME ?= switch-template
VERSION ?= $(shell git describe --tags)
IP ?= $(shell cat dtargets.lst)

all: build-prod

# project env setup
env:
	git submodule init
	git submodule update

serve:
	go run ./main.go

test:
	@echo "Start running tests"
	@go test -count=1 -cover ./...
	@echo "------------------------"
	@echo "[v] Test => pass all tests"

test-verbose:
	@go test -count=1 -cover -v ./...

tag:
	@printf "[v] Version => "
	@git describe --tags --abbrev=0 | awk -F. '{OFS="."; $$NF+=1; print $$0}' | xargs -t -I % sh -c 'git tag %'

clean:
	@printf "[_] Clean => binary files"\\r
	@rm -f ./$(NAME) ./$(NAME)-*
	@echo "[v] Clean => binary files"

build-test:
	@printf "[_] Build Test"\\r
	@$(GO_OPTS) go test -c ./...
	@echo "[v] Build Test"

build-local: clean
	@printf "[_] Build => $(GO_OPTS)"\\r
	@go build -o $(NAME)-$(VERSION)
	@echo "[v] Build => $(GO_OPTS)"

build-prod: clean
	@printf "[_] Build => $(GO_OPTS)"\\r
	@$(GO_OPTS) go build -o $(NAME) ./main.go
	-@$(CROSS)-strip $(NAME)
	-@say $(NAME) is successfully built.
	@echo "[v] Build => $(GO_OPTS)"

mod-tidy:
	@printf "[_] Run go mod tidy\r"
	-@go clean -modcache
	@go mod tidy
	@printf "[v] Run go mod tidy\n"