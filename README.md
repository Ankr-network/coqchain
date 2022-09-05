## `CoqChain`

this chain fork from `go-ethereum` and support some advanced features.

## Status

it still work in progress, coming soon.

## Features:

- [x] supply 0 gas fee contract, it is convenient for enterprise applications
- [x] lower and fixed gas fee
- [x] Fixed governance token supply
- [x] Recycle governance tokens
- [x] support to prune state data automatically
- [x] support to slash evil validator
- [x] support to remove evil validator automatically
- [x] every sync node should have some reward
- [ ] support original bridge, which can interact with `BSC` and `Ethereum` chain
- [x] speed up transactions
- [ ] implement special storage layer for store chain data better and more efficient
- [ ] replace `Merkle Tree` with `Verkle tree`
- [ ] etc...

## Bridge Architect



  



## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.
