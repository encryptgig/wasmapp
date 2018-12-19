all: clean
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

clean:
	rm -rf main.wasm

run:
	go run server.go

.PHONY: all
