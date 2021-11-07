import { OrderContract } from "./types";

export default class PutOrder {
  putBidPrice: number;
  putBidSize: number;
  putAskPrice: number;
  putAskSize: number;

  constructor(data: OrderContract) {
    this.putBidPrice = data.bid.px;
    this.putBidSize = data.bid.size;
    this.putAskPrice = data.ask.px;
    this.putAskSize = data.ask.size;
  }
}