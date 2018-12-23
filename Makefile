all: clean
	GOOS=js GOARCH=wasm go get -u github.com/rajnikant12345/engine
	cp ${GOROOT}/misc/wasm/wasm_exec.js .
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
