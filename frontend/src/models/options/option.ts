import CallOrder from "./call_order";
import PutOrder from "./put_order";
import { PutOrderType, CallOrderType, ProviderType, OptionContract } from "./types";

export default class Option {
  provider: ProviderType;
  put: PutOrderType;
  call: CallOrderType;
  strike: number;


  constructor(data: OptionContract) {
    this.provider = data.provider;
    this.put = new PutOrder(data.put);
    this.call = new CallOrder(data.call);
    this.strike = data.strike;
  }

  toRow(): Record<string, number | string> {
    return {
      provider: this.provider,
      ...this.put,
      strike: this.strike,
      ...this.call,
    }
  }
}
