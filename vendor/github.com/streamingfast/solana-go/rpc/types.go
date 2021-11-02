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

package rpc

import (
	bin "github.com/streamingfast/binary"
	"github.com/streamingfast/solana-go"
)

type Context struct {
	Slot bin.Uint64
}

type RPCContext struct {
	Context Context `json:"context,omitempty"`
}

type GetBalanceResult struct {
	RPCContext
	Value bin.Uint64 `json:"value"`
}

type GetSlotResult bin.Uint64

type GetRecentBlockhashResult struct {
	RPCContext
	Value BlockhashResult `json:"value"`
}

type BlockhashResult struct {
	Blockhash     solana.PublicKey `json:"blockhash"` /* make this a `Hash` type, which is a copy of the PublicKey` type */
	FeeCalculator FeeCalculator    `json:"feeCalculator"`
}

type FeeCalculator struct {
	LamportsPerSignature bin.Uint64 `json:"lamportsPerSignature"`
}

type GetConfirmedBlockResult struct {
	Blockhash         solana.PublicKey      `json:"blockhash"`
	PreviousBlockhash solana.PublicKey      `json:"previousBlockhash"` // could be zeroes if ledger was clean-up and this is unavailable
	ParentSlot        bin.Uint64            `json:"parentSlot"`
	Transactions      []TransactionWithMeta `json:"transactions"`
	Rewards           []BlockReward         `json:"rewards"`
	BlockTime         bin.Uint64            `json:"blockTime,omitempty"`
}

type BlockReward struct {
	Pubkey   solana.PublicKey `json:"pubkey"`
	Lamports bin.Uint64       `json:"lamports"`
}

type TransactionWithMeta struct {
	Transaction *solana.Transaction `json:"transaction"`
	Meta        *TransactionMeta    `json:"meta,omitempty"`
}

type TransactionParsed struct {
	Transaction *ParsedTransaction `json:"transaction"`
	Meta        *TransactionMeta   `json:"meta,omitempty"`
}

type TransactionMeta struct {
	Err          interface{}  `json:"err"`
	Fee          bin.Uint64   `json:"fee"`
	PreBalances  []bin.Uint64 `json:"preBalances"`
	PostBalances []bin.Uint64 `json:"postBalances"`
}

type TransactionSignature struct {
	Err       interface{} `json:"err,omitempty"`
	Memo      string      `json:"memo,omitempty"`
	Signature string      `json:"signature,omitempty"`
	Slot      bin.Uint64  `json:"slot,omitempty"`
}

type GetAccountInfoResult struct {
	RPCContext
	Value *Account `json:"value"`
}

type Account struct {
	Lamports   bin.Uint64       `json:"lamports"`
	Data       solana.Data      `json:"data"`
	Owner      solana.PublicKey `json:"owner"`
	Executable bool             `json:"executable"`
	RentEpoch  bin.Uint64       `json:"rentEpoch"`
}

type GetProgramAccountsOpts struct {
	Commitment CommitmentType `json:"commitment,omitempty"`

	// Filter on accounts, implicit AND between filters
	Filters []RPCFilter `json:"filters,omitempty"`
}

type GetProgramAccountsResult []*KeyedAccount

type KeyedAccount struct {
	Pubkey  solana.PublicKey `json:"pubkey"`
	Account *Account         `json:"account"`
}

type GetConfirmedSignaturesForAddress2Opts struct {
	Limit  uint64 `json:"limit,omitempty"`
	Before string `json:"before,omitempty"`
	Until  string `json:"until,omitempty"`
}

type GetConfirmedSignaturesForAddress2Result []*TransactionSignature

type GetSignaturesForAddressOpts struct {
	Limit  uint64 `json:"limit,omitempty"`
	Before string `json:"before,omitempty"`
	Until  string `json:"until,omitempty"`
}

type GetSignaturesForAddressResult []*TransactionSignature

type RPCFilter struct {
	Memcmp   *RPCFilterMemcmp `json:"memcmp,omitempty"`
	DataSize bin.Uint64       `json:"dataSize,omitempty"`
}

type RPCFilterMemcmp struct {
	Offset int           `json:"offset"`
	Bytes  solana.Base58 `json:"bytes"`
}

// CommitmentType is the level of commitment desired when querying state.
// https://docs.solana.com/developing/clients/jsonrpc-api#configuring-state-commitment
type CommitmentType string

const (
	// CommitmentProcessed queries the most recent block which has reached 1 confirmation by the connected node
	CommitmentProcessed = CommitmentType("processed")
	// CommitmentConfirmed queries the most recent block which has reached 1 confirmation by the cluster
	CommitmentConfirmed = CommitmentType("confirmed")
	// CommitmentConfirmed queries the most recent block which has been finalized by the cluster
	CommitmentFinalized = CommitmentType("finalized")

	// The following are deprecated

	CommitmentMax          = CommitmentType("max")          // Deprecated as of v1.5.5
	CommitmentRecent       = CommitmentType("recent")       // Deprecated as of v1.5.5
	CommitmentRoot         = CommitmentType("root")         // Deprecated as of v1.5.5
	CommitmentSingle       = CommitmentType("single")       // Deprecated as of v1.5.5
	CommitmentSingleGossip = CommitmentType("singleGossip") // Deprecated as of v1.5.5
)

/// Parsed Transaction

type ParsedTransaction struct {
	Signatures []solana.Signature `json:"signatures"`
	Message    Message            `json:"message"`
}

type Message struct {
	AccountKeys     []*AccountKey `json:"accountKeys"`
	RecentBlockhash solana.PublicKey/* TODO: change to Hash */ `json:"recentBlockhash"`
	Instructions    []ParsedInstruction `json:"instructions"`
}

type AccountKey struct {
	PublicKey solana.PublicKey `json:"pubkey"`
	Signer    bool             `json:"signer"`
	Writable  bool             `json:"writable"`
}

type ParsedInstruction struct {
	Accounts  []solana.PublicKey `json:"accounts,omitempty"`
	Data      solana.Base58      `json:"data,omitempty"`
	Parsed    *InstructionInfo   `json:"parsed,omitempty"`
	Program   string             `json:"program,omitempty"`
	ProgramID solana.PublicKey   `json:"programId"`
}

type InstructionInfo struct {
	Info            map[string]interface{} `json:"info"`
	InstructionType string                 `json:"type"`
}

func (p *ParsedInstruction) IsParsed() bool {
	return p.Parsed != nil
}

type SendTransactionOptions struct {
	SkipPreflight       bool           // disable transaction verification step
	PreflightCommitment CommitmentType // preflight commitment level; default: "finalized"
}
