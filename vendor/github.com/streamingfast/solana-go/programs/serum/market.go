// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package serum

import (
	"math/big"

	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/token"
)

type MarketMeta struct {
	Address    solana.PublicKey `json:"address"`
	Name       string           `json:"name"`
	Deprecated bool             `json:"deprecated"`
	QuoteMint  token.Mint
	BaseMint   token.Mint

	Market Market
}

func (m *MarketMeta) baseSplTokenMultiplier() *big.Int {
	return solana.DecimalsInBigInt(uint32(m.BaseMint.Decimals))
}

func (m *MarketMeta) quoteSplTokenMultiplier() *big.Int {
	return solana.DecimalsInBigInt(uint32(m.QuoteMint.Decimals))
}

func (m *MarketMeta) PriceLotsToNumber(price *big.Int) *big.Float {
	ratio := I().Mul(I().SetInt64(int64(m.Market.GetQuoteLotSize())), m.baseSplTokenMultiplier())
	numerator := F().Mul(F().SetInt(price), F().SetInt(ratio))
	denomiator := F().Mul(F().SetFloat64(float64(m.Market.GetBaseLotSize())), F().SetInt(m.quoteSplTokenMultiplier()))
	v := divideBnToNumber(numerator, denomiator)
	return v
}

func (m *MarketMeta) BaseSizeLotsToNumber(size *big.Int) *big.Float {
	numerator := I().Mul(size, I().SetInt64(int64(m.Market.GetBaseLotSize())))
	denomiator := m.baseSplTokenMultiplier()
	return F().Quo(F().SetInt(numerator), F().SetInt(denomiator))
}

func (m *MarketMeta) PriceNumberToLots(price *big.Int) *big.Float {
	numerator := I().Mul(price, m.quoteSplTokenMultiplier())
	numerator = I().Mul(numerator, big.NewInt(int64(m.Market.GetBaseLotSize())))
	denomiator := I().Mul(m.baseSplTokenMultiplier(), I().SetInt64(int64(m.Market.GetQuoteLotSize())))
	return F().Quo(F().SetInt(numerator), F().SetInt(denomiator))
}

type OpenOrdersMeta struct {
	OpenOrders OpenOrders
}

type SideLayoutType string

const (
	SideLayoutTypeUnknown SideLayoutType = "UNKNOWN"
	SideLayoutTypeBid     SideLayoutType = "BID"
	SideLayoutTypeAsk     SideLayoutType = "ASK"
)

type SideLayout uint32

func (s SideLayout) getSide() SideLayoutType {
	switch s {
	case 0:
		return SideLayoutTypeBid
	case 1:
		return SideLayoutTypeAsk
	}
	return SideLayoutTypeUnknown
}

type OrderType uint32

const (
	OrderTypeLimit = OrderType(iota)
	OrderTypeImmediateOrCancel
	OrderTypePostOnly
)
