wasm:
	cd ./game/wasm && GOOS=js GOARCH=wasm go build -o ../../assets/json.wasm -ldflags "-w"
	cp "${GOROOT}/misc/wasm/wasm_exec.js" ./build/assets/
dev: wasm
	cd ./cmd/server && go run .