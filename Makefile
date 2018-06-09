NAME = graffity

BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
COMMIT = $(shell git rev-parse --short HEAD)
BUILDTIME = $(shell date +%Y-%m-%dT%T%z)
LASTTAG = $(shell git describe --tags --abbrev=0 --dirty)
GOPATH = $(shell echo "$$GOPATH")
LD_OPTS = -ldflags="-X main.branch=${BRANCH} -X main.commit=${COMMIT} -X main.lasttag=${LASTTAG} -X main.buildtime=${BUILDTIME} -w "

all:  build run

all-with-deps: setup deps
	cd cmd && GOOS=linux GOARCH=amd64 go build $(LD_OPTS)  -o $(NAME) .

run:
	cd $(GOPATH)/src/github.com/Appscrunch/Multy-back/cmd && rm -rf multy && cd .. && make build  && cd cmd && ./$(NAME) && ../


	
build:
	cd node-streamer/btc/ && protoc --go_out=plugins=grpc:. *.proto && cd ../../cmd/ && go build $(LD_OPTS) -o $(NAME) . && cd -

# Show to-do items per file.
todo:
	@grep \
		--exclude-dir=vendor \
		--exclude-dir=node_modules \
		--exclude=Makefile \
		--text \
		--color \
		-nRo -E ' TODO:.*|SkipNow|nolint:.*' .
.PHONY: todo

