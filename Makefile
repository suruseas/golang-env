.PHONY: build build-wasm web

build:
	GOOS=js GOARCH=wasm go build -o ./cmd/web/static/main.wasm ./cmd/hello/main.go

download-wasm-exec-js:
	ash -c "cd cmd/web/static && curl -LJO https://raw.githubusercontent.com/golang/go/go1.20.5/misc/wasm/wasm_exec.js"

copy-wasm-exec-js:
	find /usr/local/go/ -name wasm_exec.js | xargs -i cp -p {} cmd/web/static

web:
	go run ./cmd/web/main.go

