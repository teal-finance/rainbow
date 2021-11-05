import { OrderContract } from "./types";

export default class CallOrder {
  callBid: number;
  callAsk: number;
  callIv: number;
  callLastPrice: number;
  callChange: number;
  callVolume: number;
  callOpen: number;

  constructor(data: OrderContract) {
    this.callBid = data.bid;
    this.callAsk = data.ask;
    this.callIv = data.iva;
    this.callLastPrice = data.last_price;
    this.callChange = data.change;
    this.callVolume = data.volume;
    this.callOpen = data.open;
  }
}