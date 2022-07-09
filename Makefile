.PHONY: build
build:
	go build -o server

.PHONY: run
run: build
	./server