all: clean
	cp /usr/local/go/misc/wasm/wasm_exec.js .
	GOOS=js GOARCH=wasm go build -o main.wasm main.go
	rm -rf site
	mkdir site
	cp *.css site/
	cp *.js site/
	cp *.wasm site/
	cp *.html site/
	cp *.woff site/

clean:
	rm -rf main.wasm

run:
	go run server.go

.PHONY: all
