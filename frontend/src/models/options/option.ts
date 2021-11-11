import CallLimit from "./call_order";
import PutLimit from "./put_order";
import { CallType, PutType, ProviderType, CallPutContract } from "./types";

export default class Option {
  provider: ProviderType;
  asset: string;
  expiry: Date;
  call: CallType;
  strike: number;
  put: PutType;

  constructor(data: CallPutContract) {
    this.provider = data.provider;
    this.asset = data.asset;
    this.expiry = new Date(data.expiry);
    this.call = new CallLimit(data.call);
    this.strike = data.strike;
    this.put = new PutLimit(data.put);
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
