// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

type ProviderType = "PsyOptions" | "Deribit" | "Opyn";

// contracts for the json

interface OrderContract {
  px: string;
  size: string;
}

interface LimitContract {
  bid: OrderContract;
  ask: OrderContract;
}

interface CallPutContract {
  provider: ProviderType
  asset: string;
  date: string;
  expiry: string;
  call: LimitContract;
  strike: number;
  put: LimitContract;
}

// internal types

type OrderType = {
  px: string;
  size: string;
}

type CallType = {
  callBidPrice: string;
  callBidSize: string;
  callAskPrice: string;
  callAskSize: string;
}

type PutType = {
  putBidPrice: string;
  putBidSize: string;
  putAskPrice: string;
  putAskSize: string;
}

type RowType = {
  provider: ProviderType;
  asset: string;
  date: Date;
  expiry: string;
  call: CallType;
  strike: number;
  put: PutType;
}

type OptionsJsonDataset = {
  rows: Array<CallPutContract>;
}

type OptionsTable = Record<string, number | string | Date>;

export {
  ProviderType,
  LimitContract,
  CallPutContract,
  CallType,
  PutType,
  RowType,
  OptionsJsonDataset,
  OrderType,
  OptionsTable,
}
