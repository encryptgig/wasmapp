GOROOT = ${GOROOT}

all: clean
	cp GOROOT/misc/wasm/wasm_exec.js .
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

clean:
	rm -rf main.wasm

run:
	go run server.go

.PHONY: all
