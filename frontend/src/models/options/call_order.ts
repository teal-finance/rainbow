// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

import { LimitContract } from "./types";

export default class CallLimit {
  callBidPrice: string;
  callBidSize: string;
  callAskPrice: string;
  callAskSize: string;

  constructor(data: LimitContract) {
    this.callBidPrice = data.bid.px;
    this.callBidSize = data.bid.size;
    this.callAskPrice = data.ask.px;
    this.callAskSize = data.ask.size;
  }
}
