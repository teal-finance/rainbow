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