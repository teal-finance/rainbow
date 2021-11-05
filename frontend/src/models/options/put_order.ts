import { OrderContract } from "./types";

export default class PutOrder {
  putBid: number;
  putAsk: number;
  putIv: number;
  putLastPrice: number;
  putChange: number;
  putVolume: number;
  putOpen: number;

  constructor(data: OrderContract) {
    this.putBid = data.bid;
    this.putAsk = data.ask;
    this.putIv = data.iva;
    this.putLastPrice = data.last_price;
    this.putChange = data.change;
    this.putVolume = data.volume;
    this.putOpen = data.open;
  }
}