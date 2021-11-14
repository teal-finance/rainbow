type ProviderType = "PsyOptions" | "Deribit" | "Opyn";

// contracts for the json

interface OrderContract {
  px: string;
  size: string;
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
  strike: string;
  put: LimitContract;
}

// internal types

type OrderType = {
  px: string;
  size: string;
}

type CallType = {
  callBidPrice: string;
  callBidSize: string;
  callAskPrice: string;
  callAskSize: string;
}

type PutType = {
  putBidPrice: string;
  putBidSize: string;
  putAskPrice: string;
  putAskSize: string;
}

type RowType = {
  provider: ProviderType;
  asset: string;
  expiry: Date;
  call: CallType;
  strike: string;
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