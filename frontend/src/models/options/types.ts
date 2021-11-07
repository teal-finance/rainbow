type ProviderType = "PsyOptions" | "Deribit" | "Opyn";

// contracts for the json

interface OrderFrameContract {
  px: number;
  size: number;
}

interface OrderContract {
  bid: OrderFrameContract;
  ask: OrderFrameContract;
}

interface OptionContract {
  provider: ProviderType
  asset: string;
  expiry: string;
  put: OrderContract;
  call: OrderContract;
  strike: number;
}

// internal types

type OrderFrameType = {
  px: number;
  size: number;
}

type PutOrderType = {
  putBidPrice: number;
  putBidSize: number;
  putAskPrice: number;
  putAskSize: number;
}

type CallOrderType = {
  callBidPrice: number;
  callBidSize: number;
  callAskPrice: number;
  callAskSize: number;
}

type OptionType = {
  provider: ProviderType;
  asset: string;
  expiry: Date;
  put: PutOrderType;
  call: CallOrderType;
  strike: number;
}

type OptionsJsonDataset = {
  provider: ProviderType
  asset: string;
  expiry: string;
  rows: Array<OptionContract>;
}

type OptionsTable = Record<string, number | string | Date>;

export {
  ProviderType,
  OrderContract,
  OptionContract,
  PutOrderType,
  CallOrderType,
  OptionType,
  OptionsJsonDataset,
  OrderFrameType,
  OptionsTable,
}