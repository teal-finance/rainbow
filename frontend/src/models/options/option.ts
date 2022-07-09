// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

import CallLimit from "./call_order";
import PutLimit from "./put_order";
import { CallType, PutType, ProviderType, CallPutContract } from "./types";

export default class Option {
  provider: ProviderType;
  asset: string;
  date: Date;
  expiry: string;
  call: CallType;
  strike: number;
  put: PutType;

  constructor(data: CallPutContract) {
    this.provider = data.provider;
    this.asset = data.asset;
    this.date = new Date(data.date);
    this.expiry = data.expiry;
    this.call = new CallLimit(data.call);
    this.strike = data.strike;
    this.put = new PutLimit(data.put);
  }

  toRow(): Record<string, number | string | Date> {
    return {
      provider: this.provider,
      asset: this.asset,
      date: this.date.toLocaleDateString(),
      expiry: this.expiry,
      ...this.call,
      strike: this.strike,
      ...this.put,
    }
  }
}
