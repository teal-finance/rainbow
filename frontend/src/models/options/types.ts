type ProviderType = "PsyOptions" | "Deribit" | "Opyn";

// contracts for the json

interface OrderContract {
  px: number;
  size: number;
}

interface LimitContract {
  bid: OrderContract;
  ask: OrderContract;
}

interface CallPutContract {
  provider: ProviderType
  asset: string;
  expiry: string;
  call: LimitContract;
  strike: number;
  put: LimitContract;
}

// internal types

type OrderType = {
  px: number;
  size: number;
}

type CallType = {
  callBidPrice: number;
  callBidSize: number;
  callAskPrice: number;
  callAskSize: number;
}

type PutType = {
  putBidPrice: number;
  putBidSize: number;
  putAskPrice: number;
  putAskSize: number;
}

type RowType = {
  provider: ProviderType;
  asset: string;
  expiry: Date;
  call: CallType;
  strike: number;
  put: PutType;
}

type OptionsJsonDataset = {
  rows: Array<CallPutContract>;
}

type OptionsTable = Record<string, number | string | Date>;

export {
  ProviderType,
  LimitContract,
  CallPutContract,
  CallType,
  PutType,
  RowType,
  OptionsJsonDataset,
  OrderType,
  OptionsTable,
}