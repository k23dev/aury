DIR_FRONTEND:=./frontend
DIR_ASSETS:=frontend/public/assets/
DIR_ASSETS_WASM:=${DIR_ASSETS}WASM/
wasm:
	mkdir -p ./${DIR_ASSETS_WASM}
	cd ./wasm && GOOS=js GOARCH=wasm go build -o ../${DIR_ASSETS_WASM}game.wasm -ldflags "-w"
	cp "${GOROOT}/misc/wasm/wasm_exec.js" ./${DIR_ASSETS_WASM}
frontend:
	rm -rf ./frontend
	npx create-vite frontend --template svelte-ts
	npm install -D svelte @sveltejs/vite-plugin-svelte typescript
	# Install tailwindcss
	npm install -D tailwindcss postcss autoprefixer
	npx tailwindcss init -p
	npm install -D flowbite-svelte flowbite
dev: wasm
	cd ${DIR_FRONTEND} && npm run dev
wasm-build: wasm
	cd ${DIR_FRONTEND} && npm run build

