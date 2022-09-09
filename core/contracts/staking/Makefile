.PHONY: build all

build:
	solc --overwrite --abi --bin Staker.sol -o staker
	abigen  --bin=staker/staker.bin --abi=staker/staker.abi --pkg=staker --out=staker/staker.go