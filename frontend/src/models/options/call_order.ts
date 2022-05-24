import { LimitContract } from "./types";

export default class CallLimit {
  callBidIv:number;
  callBidPrice: string;
  callBidSize: string;
  callAskIv:number;
  callAskPrice: string;
  callAskSize: string;

  constructor(data: LimitContract) {
    this.callBidIv = data.bid.iv;
    this.callBidPrice = data.bid.px;
    this.callBidSize = data.bid.size;
    this.callAskIv = data.ask.iv;
    this.callAskPrice = data.ask.px;
    this.callAskSize = data.ask.size;
  }
}