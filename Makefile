UNAME:=$(shell uname|sed 's/.*/\u&/')
OS:=$(shell echo $(GOOS)| sed 's/.*/\u&/')
PKG=$(shell basename $$(pwd))

ifeq  ($(BITBUCKET_BUILD_NUMBER),)
TYPE:="Local"
else
TYPE:=$(BITBUCKET_BUILD_NUMBER)
endif
ifeq  ($(FILENAME),)
FNAME:=$(PKG)
else
FNAME:=$(FILENAME)
endif


main: *.go clean
ifeq ($(GOOS),)
	@printf "OS not specified, defaulting to: \e[33m$(UNAME)\e[39m\n"
else
	@printf "OS specified: \e[33m$$(echo $$GOOS | sed 's/.*/\u&/' )\e[39m\n"
endif
	@echo "Building..."
	@if [ ! -n  "$$BITBUCKET_BUILD_NUMBER" ]; then export BITBUCKET_BUILD_NUMBER=$(TYPE); fi;\
	export GOARCH=amd64; \
	export CGO_ENABLED=0;\
	export GitCommit=`git rev-parse HEAD | cut -c -7`;\
	export BuildTime=`date -u +%Y%m%d.%H%M%S`;\
	export Authors=`git log -50 --format='%aN' | sort -u | sed "s@root@@"  | tr '\n' ';' | sed "s@;;@;@g" | sed "s@;@; @g" | sed "s@\(.*\); @\1@" | sed "s@[[:blank:]]@SpAcE@g"`;\
	export GitTag=$$(TAG=`git tag --contains $$(git rev-parse HEAD) | sort -R | tr '\n' ' '`; if [ "$$(printf "$$TAG")" ]; then printf "$$TAG"; else printf "undefined"; fi);\
	go build -ldflags "-X main.BuildNo=$$BUILD_NUMBER -X main.GitCommit=$$GitCommit -X main.Tag=$$GitTag -X main.BuildTime=$$BuildTime -X main.Authors=$$Authors" -o $(FNAME) bin/main.go
	@printf "\e[32mSuccess!\e[39m\n"
clean:  
	@printf "Cleaning up \e[32mmain\e[39m...\n"
	rm -f main $(FNAME) || rm -rf main $(FNAME)

install: clean main
	mv $(FNAME) "$$GOPATH/bin/$(PKG)"

vet: 
	@echo "Running go vet..."
	@go vet || (printf "\e[31mGo vet failed, exit code $$?\e[39m\n"; exit 1)
	@printf "\e[32mGo vet success!\e[39m\n"

