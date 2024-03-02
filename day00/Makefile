build:
	go build ./cmd/main.go

run: build
	./main

test:
	go test internal/anscombe/anscombe.go internal/anscombe/anscombe_test.go

deps:
	go mod tidy

clean:
	@rm -rf main
