// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

import { LimitContract } from "./types";

export default class PutLimit {
  putBidPrice: string;
  putBidSize: string;
  putAskPrice: string;
  putAskSize: string;

  constructor(data: LimitContract) {
    this.putBidPrice = data.bid.px;
    this.putBidSize = data.bid.size;
    this.putAskPrice = data.ask.px;
    this.putAskSize = data.ask.size;
  }
}
