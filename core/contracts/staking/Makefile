.PHONY: build all

build:
	solc --overwrite --abi --bin staker.sol -o staker
	abigen  --bin=staker/Staker.bin --abi=staker/Staker.abi --pkg=staker --out=staker/staker.go