type ProviderType = "PsyOptions" | "Deribit";

// contracts for the json

interface OrderContract {
  bid: number;
  ask: number;
  iva: number;
  last_price: number;
  change: number;
  volume: number;
  open: number;
}

interface OptionContract {
  put: OrderContract;
  call: OrderContract;
  strike: number;
  provider: ProviderType
}

// internal types

type PutOrderType = {
  putBid: number;
  putAsk: number;
  putIv: number;
  putLastPrice: number;
  putChange: number;
  putVolume: number;
  putOpen: number;
}

type CallOrderType = {
  callBid: number;
  callAsk: number;
  callIv: number;
  callLastPrice: number;
  callChange: number;
  callVolume: number;
  callOpen: number;
}

type OptionType = {
  put: PutOrderType;
  call: CallOrderType;
  strike: number;
  provider: ProviderType
}

type OptionsJsonDataset = {
  asset: string;
  expriry: string;
  table: Array<OptionContract>;
}

export { ProviderType, OrderContract, OptionContract, PutOrderType, CallOrderType, OptionType, OptionsJsonDataset }