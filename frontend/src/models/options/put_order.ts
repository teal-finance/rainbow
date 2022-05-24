import { LimitContract } from "./types";

export default class PutLimit {
  putBidIv: number;
  putBidPrice: string;
  putBidSize: string;
  putAskIv:number;
  putAskPrice: string;
  putAskSize: string;

  constructor(data: LimitContract) {
    this.putBidIv=data.bid.iv;
    this.putBidPrice = data.bid.px;
    this.putBidSize = data.bid.size;
    this.putAskIv =data.ask.iv;
    this.putAskPrice = data.ask.px;
    this.putAskSize = data.ask.size;
  }
}