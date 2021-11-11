import { LimitContract } from "./types";

export default class PutLimit {
  putBidPrice: number;
  putBidSize: number;
  putAskPrice: number;
  putAskSize: number;

  constructor(data: LimitContract) {
    this.putBidPrice = data.bid.px;
    this.putBidSize = data.bid.size;
    this.putAskPrice = data.ask.px;
    this.putAskSize = data.ask.size;
  }
}