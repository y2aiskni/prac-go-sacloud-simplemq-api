.PHONY: build
build:
	go build -o enqueue ./cmd/enqueue
	go build -o dequeue ./cmd/dequeue
	go build -o peek ./cmd/peek
	go build -o length ./cmd/length
