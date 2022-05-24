type ProviderType = "PsyOptions" | "Deribit" | "Opyn";

// contracts for the json

interface OrderContract {
  px: string;
  size: string;
  iv: number;
}

interface LimitContract {
  bid: OrderContract;
  ask: OrderContract;
}

interface CallPutContract {
  provider: ProviderType
  asset: string;
  date: string;
  expiry: string;
  call: LimitContract;
  strike: number;
  put: LimitContract;
}

// internal types

type OrderType = {
  px: string;
  size: string;
  iv: number;
}

type CallType = {
  callBidIv: number;
  callBidPrice: string;
  callBidSize: string;
  callAskIv: number;
  callAskPrice: string;
  callAskSize: string;
}

type PutType = {
  putBidIv: number;
  putBidPrice: string;
  putBidSize: string;
  putAskIv:number;
  putAskPrice: string;
  putAskSize: string;
}

type RowType = {
  provider: ProviderType;
  asset: string;
  date: Date;
  expiry: string;
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