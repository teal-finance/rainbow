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
    this.call = new CallOrder(data.call);
    this.strike = data.strike;
    this.put = new PutOrder(data.put);
  }

  toRow(): Record<string, number | string | Date> {
    return {
      provider: this.provider,
      asset: this.asset,
      expiry: this.expiry.toLocaleDateString(),
      ...this.call,
      strike: this.strike,
      ...this.put,
    }
  }
}
