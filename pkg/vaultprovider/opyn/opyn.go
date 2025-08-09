// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package opyn

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spewerspew/spew"
	"github.com/teal-finance/emo"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

const (
	Osqth             = "0xf1B99e3E573A1a9C5E6B2Ce818b617F0E664E86B" // Squeeth token address 0squth=Opyn Squeeth
	ControllerAddress = "0x64187ae08781B09368e6253F9E94951243A493D5" // controller address
	InfuraRPC         = "https://mainnet.infura.io/v3/c1b94bff90754066a81d195ddc337ff3"
)

var log = emo.NewZone("Opyn")

type VaultsProvider struct{}

func (VaultsProvider) Name() string {
	return "Opyn"
}

type Vault struct {
	Operator         common.Address
	NftCollateralId  uint32
	CollateralAmount *big.Int
	ShortAmount      *big.Int
}

// GetClient returns a rpc client for mainnet.
func GetClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(InfuraRPC)
	if err != nil {
		log.Print("client", err)
		return &ethclient.Client{}, err
	}
	return client, nil
}

func GetVault(v int64, client *ethclient.Client) (Vault, error) {
	controller, err := NewController(common.HexToAddress(ControllerAddress), client)
	if err != nil {
		log.Print("controller", err)
		return Vault{}, err
	}
	vault, err := controller.Vaults(B, big.NewInt(v))
	if err != nil {
		log.Print("Vaults", err)
		return Vault{}, nil
	}
	/*return Vault{
		Operator:         vault.Operator,
		NftCollateralId:  vault.NftCollateralId,
		CollateralAmount: vault.CollateralAmount,
		ShortAmount:      vault.ShortAmount},
	err*/

	return vault, err
}

// HasNFT if the vault has a Uniswap LP NFT (!=0).
func (v Vault) HasNFT() bool {
	return v.NftCollateralId != 0
}

// Twap: returns the Time Weighted Average Price of ETH for the last 420 blocks
// This is the period used for liquidation.
func Twap(controller *Controller) float64 {
	index, err := controller.GetIndex(B, 420)
	if err != nil {
		log.Print("ERR: ", err)
		return 0
	}
	spew.Dump(index)
	index.Sqrt(index)
	spew.Dump(index)
	return rainbow.ToFloat(index, 5)
}

var B *bind.CallOpts = &bind.CallOpts{}
