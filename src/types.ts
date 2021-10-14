export type SideType = "PUT" | "CALL";

export type OptionData = {
  Name: string;
  Asset: string;
  Bid: [{
    Price: number,
    Quantity: number,
    QuoteCurrency: string,
  }];
  Ask: [{
    Price: number,
    Quantity: number,
    QuoteCurrency: string,
  }];
  Chain: string;
  ExchangeType: string;
  Expiry: string;
  Layer: string;
  Provider: string;
  Strike: number;
  Type: SideType;
};