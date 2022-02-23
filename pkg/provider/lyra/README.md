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