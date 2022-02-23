package lyra

import (
	"fmt"
	"log"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const optimismrpc = "https://mainnet.optimism.io"

type Provider struct{}

func (Provider) Name() string {
	return "Lyra"
}

func (Provider) Options() ([]rainbow.Option, error) {

	return []rainbow.Option{}, nil
}

func Try() error {
	client, err := ethclient.Dial(optimismrpc)
	if err != nil {
		log.Fatal(err)
	}
	sum := 0
	for i := 0; i < len(OptionMarkets); i++ {
		//viewer:=OptionMarketsViewer[i]
		addressMarket := common.HexToAddress(OptionMarkets[i])
		market, err := NewLyra(addressMarket, client)
		if err != nil {
			return err
		}
		addressViewer := common.HexToAddress(OptionMarketViewers[i])
		viewer, err := NewLyrap(addressViewer, client)
		if err != nil {
			return err
		}

		boards, err := market.GetLiveBoards(&bind.CallOpts{})
		if err != nil {
			return err
		}

		fmt.Println("GetLiveBoards ", boards)

		for _, j := range boards {
			boardListings, err := market.GetBoardListings(&bind.CallOpts{}, j)
			if err != nil {
				return err
			}
			fmt.Println("getBoardListings ", boardListings)
			sum += len(boardListings)
			for _, k := range boardListings {

				listings, err := market.OptionListings(&bind.CallOpts{}, k)
				if err != nil {
					return err
				}
				fmt.Println("OptionsListings ", listings)

				vlist, err := viewer.GetListingView(&bind.CallOpts{}, k)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("OptionsListings ", vlist)
				spew.Dump(listings)

				spew.Dump(vlist)
				amount := big.Int{}
				amount.SetString("1000000000000000000", 10) //1000000000000000000
				trade, err := viewer.GetPremiumForTrade(&bind.CallOpts{}, k, 0, true, &amount)
				if err != nil {
					log.Fatal(err)
				}
				spew.Dump(trade)
			}
		}

	}
	spew.Dump(sum)
	return nil
}
