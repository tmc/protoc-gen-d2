.PHONY: all buf-generate clean

all: buf-generate

buf-generate:
	# buf will generate the protobuf code
	go install ..
	buf generate
