.PHONY: test
test:
	go test -v

.PHONY: build
build: test
	go build -o bin/chaoscmd

.PHONY: clean
clean:
	go clean
	rm -f bin/*