.PHONY: build all

build:
	solc --overwrite --abi --bin Store.sol -o store
	abigen  --bin=store/Store.bin --abi=store/Store.abi --pkg=store --out=store/Store.go