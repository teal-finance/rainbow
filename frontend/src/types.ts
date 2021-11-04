// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

export type SideType = "PUT" | "CALL";

export type OptionData = {
  Name: string;
  Asset: string;
  Bid: [{
    Price: number,
    Quantity: number,
    QuoteCurrency: string,
  }] | null;
  Ask: [{
    Price: number,
    Quantity: number,
    QuoteCurrency: string,
  }] | null;
  Chain: string;
  ExchangeType: string;
  Expiry: string;
  Layer: string;
  Provider: string;
  Strike: number;
  Type: SideType;
};
