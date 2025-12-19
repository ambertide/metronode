wasm:
	echo "Generating WASM Library"
	GOOS=wasip1 GOARCH=wasm go build -C pkg/wasm/ -buildmode=c-shared -o ../../dist/wasm/metronode.wasm
