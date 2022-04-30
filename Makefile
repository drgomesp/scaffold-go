NAME = scaffold-go
VERSION = $$(git describe --tags --always || echo "unknown version")
BUILD_TIME = $(shell date -u)

GO_BUILD = CGO_ENABLED=0 go build -ldflags "'-X main.Version=$(VERSION)' '-X main.Build=$(BUILD_TIME)'"
build: clean
	@echo "building $(VERSION). ($(BUILD_TIME))"
	$(GO_BUILD) -o build/$(NAME) ./...

clean:
	rm -rf build/ && mkdir build/
