// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

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