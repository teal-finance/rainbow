export type SideType = "PUT" | "CALL";

export type OptionData = {
  Name: string;
  Asset: string;
  Bid: [{
    Price: number,
    Quantity: number,
    QuoteCurrency: string,
  }] | null;
  Ask: [{
    Price: number,
    Quantity: number,
    QuoteCurrency: string,
  }] | null;
  Chain: string;
  ExchangeType: string;
  Expiry: string;
  Layer: string;
  Provider: string;
  Strike: number;
  Type: SideType;
};