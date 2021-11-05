import CallOrder from "./call_order";
import PutOrder from "./put_order";
import { PutOrderType, CallOrderType, ProviderType, OptionContract } from "./types";

export default class Option {
  provider: ProviderType;
  asset: string;
  expiry: Date;
  put: PutOrderType;
  call: CallOrderType;
  strike: number;

  constructor(data: OptionContract) {
    this.provider = data.provider;
    this.asset = data.asset;
    this.expiry = new Date(data.expiry);
    this.put = new PutOrder(data.put);
    this.call = new CallOrder(data.call);
    this.strike = data.strike;
  }

  toRow(): Record<string, number | string | Date> {
    return {
      provider: this.provider,
      asset: this.asset,
      expiry: this.expiry.toLocaleDateString("en-US"),
      ...this.put,
      strike: this.strike,
      ...this.call,
    }
  }
}
