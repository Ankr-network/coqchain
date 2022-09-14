.PHONY: build all

build:
	solc --overwrite --abi --bin --storage-layout staker.sol --bin-runtime -o staker
	abigen  --bin=staker/Staker.bin --abi=staker/Staker.abi --pkg=staker --out=staker/staker.go