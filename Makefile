#
# makefile of intri core
#

CROSS                 = aarch64-linux-gnu-
AR                    = $(CROSS)ar
CC                    = $(CROSS)gcc
STRIP                 = $(CROSS)strip
# GoBuildEnv=$(CC) GOSUMDB=off CGO_ENABLED=1 GOARCH=arm GOARM=7 GOOS=linux

NowDirName=$$(echo ${PWD} | awk -F '/' '{print $$NF}')
GO_CC                 = CC=$(CC)
GO_OPTS               = CGO_ENABLED=1 GOARCH=arm64 GOOS=linux
# GO_OPTS               = GOARCH=arm64 GOARM=8 GOOS=linux

all: intri-core

tag:
	@git describe --tags --abbrev=0 | awk -F. '{OFS="."; $$NF+=1; print $$0}' | xargs -t -I % sh -c 'git tag %'

hide:
	@printf "[_] Add build files into .gitignore\r"

	@sed -E -i '.bk' "s/(.*)#(.*)$(NowDirName)/$(NowDirName)/g" ./.gitignore; rm -f ./.gitignore.bk

	@if [[ $$(sed -n "/$(NowDirName)/p" ./.gitignore) = ""  ]]; \
	then \
	echo "# repository genarate binary\n$(NowDirName)\n\n" >> ./.gitignore; \
	fi

	@printf "[v] Add build files into .gitignore\n"

mod-tidy:
	@printf "[_] Run go mod tidy\r"
	@GOSUMDB=off go get -d github.com/Intrising/intri-utils@ltos5_feat_porting
	@GOSUMDB=off go get -d github.com/Intrising/intri-type@ltos5_feat_porting
	@GOSUMDB=off go mod tidy
	@printf "[v] Run go mod tidy\n"

clean:
	@printf "[_] Clear build files\r"
	@rm -f ./intri-core*
	@printf "[v] Clear build files\n"

build: main.go clean hide mod-tidy
	@printf "[_] Building binary\r"
	@$(GO_OPTS) go build -o $(NowDirName) $<
	@printf "[v] Building binary\n"

intri-core: main.go ./*/*.go
	@printf "[_] Building binary\r"
	$(GO_CC) $(GO_OPTS) go build -o $@ $<
	$(STRIP) $@
	@printf "[v] Building binary\n"

intri-core-test:${PWD}/testing/
	@printf "[_] Building binary\r"
	$(GO_CC) $(GO_OPTS) go test $< -c  -v -o $@  
	$(STRIP) $@
	@printf "[v] Building binary\n"