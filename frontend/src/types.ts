// Copyright 2021-2022 Teal.Finance contributors
// This file is part of Teal.Finance/Rainbow,
// a screener fo DeFi options, under the MIT License.
// SPDX-License-Identifier: MIT

export type SideType = "PUT" | "CALL";

export type OptionData = {
  Name: string;
  Asset: string;
  QuoteCurrency: string,
  Bid: [{
    Price: number,
    Quantity: number,
  }] | null;
  Ask: [{
    Price: number,
    Quantity: number,
  }] | null;
  Chain: string;
  ExchangeType: string;
  Expiry: string;
  Layer: string;
  Provider: string;
  Strike: number;
  Type: SideType;
};
