// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

import Order from "./order";
import { OptionData, SideType } from "@/types";

export default class Option {
  [key: string]: any;
  name: string;
  asset: string;
  bid = new Set<Order>();
  ask = new Set<Order>();
  bestBidPx: number = 0;
  bestAskPx: number = 0;
  bestBidQty: number = 0;
  bestAskQty: number = 0;
  chain: string;
  exchangeType: string;
  expiry: string;
  layer: string;
  provider: string;
  strike: number;
  type: SideType;

  constructor(data: OptionData) {
    this.name = data.Name.toString();
    this.asset = data.Asset.toString();
    if (data.Bid) {
      for (const order of data.Bid ?? new Set<Order>()) {
        this.bid.add(new Order(order));
      }
      this.bestBidPx = data.Bid[0].Price;
      this.bestBidQty = data.Bid[0].Quantity;
    }
    if (data.Ask) {
      for (const order of data.Ask) {
        this.ask.add(new Order(order));
      }
      this.bestAskPx = data.Ask[0].Price;
      this.bestAskQty = data.Ask[0].Quantity;
    }
    this.chain = data.Chain;
    this.exchangeType = data.ExchangeType;
    this.expiry = data.Expiry;
    this.layer = data.Layer;
    this.provider = data.Provider;
    this.strike = data.Strike;
    this.type = data.Type;
  }
}