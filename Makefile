.PHONY: build clean install test

# build directive
build:
	go install

# test directive
test:
	go test ./...
	make -C testdata all
