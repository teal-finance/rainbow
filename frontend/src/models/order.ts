export default class Order {
  price: number;
  quantity: number;
  quoteCurrency: string;

  constructor({ Price, Quantity, QuoteCurrency }: { Price: number, Quantity: number, QuoteCurrency: string }) {
    this.price = Price;
    this.quantity = Quantity;
    this.quoteCurrency = QuoteCurrency;
  }
}