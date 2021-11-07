import { OrderContract } from "./types";

export default class CallOrder {
  callBidPrice: number;
  callBidSize: number;
  callAskPrice: number;
  callAskSize: number;

  constructor(data: OrderContract) {
    this.callBidPrice = data.bid.px;
    this.callBidSize = data.bid.size;
    this.callAskPrice = data.ask.px;
    this.callAskSize = data.ask.size;
  }
}