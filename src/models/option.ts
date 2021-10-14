import Bid from "./bid";
import Ask from "./ask";
import { OptionData, SideType } from "@/types";

export default class Option {
  [key: string]: any;
  name: string;
  asset: string;
  bids = new Set<Bid>();
  asks = new Set<Ask>();
  chain: string;
  exchangeType: string;
  expiry: string;
  layer: string;
  provider: string;
  strike: number;
  type: SideType;

  constructor(data: OptionData) {
    this.name = data.Name.toString();
    this.asset = data.Asset.toString();
    for (const bid of data.Bid) {
      this.bids.add(new Bid(bid));
    }
    for (const ask of data.Ask) {
      this.asks.add(new Ask(ask));
    }
    this.chain = data.Chain;
    this.exchangeType = data.ExchangeType;
    this.expiry = data.Expiry;
    this.layer = data.Layer;
    this.provider = data.Provider;
    this.strike = data.Strike;
    this.type = data.Type;
  }
}