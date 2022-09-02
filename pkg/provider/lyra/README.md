# Avalon version 

The ABIs of Lyra Avalon are available here: https://github.com/lyra-finance/lyra-protocol/blob/master/deployments/mainnet-ovm/lyra.json
Alternative solution is to copy them from *(Optimistic) Etherscan*. This is what we use for the *quoter*. We use this (fork) contract: https://github.com/cryptohazard/lyra-quoter .
We copy the ABI here: https://optimistic.etherscan.io/address/0xea83ee73eB397c5974CB6b5220AE0A32fbE48B2B#code

```
abigen --abi=./ABI/OptionMarket.abi --pkg=lyra --out=OptionMarket.go
abigen --abi=./ABI/OptionMarketViewer.abi --pkg=lyrap --out=OptionMarketViewer.go
abigen --abi=./ABI/LyraRegistry.abi --pkg=lyrar --out=LyraRegistry.go
abigen --abi=./ABI/LyraQuoter.abi  --pkg=lyraq --out=LyraQuoter.go
```
Then correct the package of `OptionMarketViewer` from `lyrap`to `lyra`.
And then, in the same file, change `OptionMarketOptionMarketParameters` to `OptionMarketOptionMarketParam` because it becomes duplicates.

Then correct the package of `LyraRegistry` from `lyrar`to `lyra`.

Finally change `lyraq`to `lyra`for quoter.


# Old version (pre-avalon)

## Small tricks

The ABIs were taken from Etherscan because I couldn't make `solc` compile properly the solidity files and the file(lyra.json) on Lyra source code is not up to date (TODO: open an issue).
`OptionMarket.sol` and `OptionMarketViewer.sol` (smart contracts) have very similar  functions so I used:

```
abigen --abi=./ABI/OptionMarket.abi --pkg=lyra --out=lyra/OptionMarket.go
abigen --abi=./build/OptionMarketViewer.abi --pkg=lyrap --out=lyra/OptionMarketViewer.go
```

Then correct the package of `OptionMarketViewer` from `lyrap`to `lyra`.
There should be an easier way to do it but that's it for the time being.

## Prices

Prices will represent the total cost for 1 option.
When it is for options buying, the prices thus will include the fees of Lyra.
When it is for options selling, the prices will deduce the fees taken by Lyra from the premium.
