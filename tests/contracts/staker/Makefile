.PHONY: build all

build:
	solc --overwrite --abi --bin --storage-layout staker.sol --bin-runtime -o .
	abigen  --bin=Staker.bin --abi=Staker.abi --pkg=staker --out=staker.go