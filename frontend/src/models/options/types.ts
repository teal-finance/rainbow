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
  provider: "PsyOptions" | "Deribit"
}