.PHONY: test
test:
	go test -v

.PHONY: build
build: test
	go build -o bin/chaoscmd

.PHONY: image
image: build
	docker build . --rm -t docker.io/fbac/chaoscmd:latest

.PHONY: clean
clean:
	go clean
	rm -f bin/*