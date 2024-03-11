wasm:
	cd ./game/wasm && GOOS=js GOARCH=wasm go build -o ../../build/assets/game.wasm -ldflags "-w"
	cp "${GOROOT}/misc/wasm/wasm_exec.js" ./build/assets/
dev: wasm
	cd ./server && go run .