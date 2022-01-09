// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

import ApiParams from "./interface";

class Api {
  url: string;
  verbose: boolean;

  constructor({ url, verbose, onError }: ApiParams) {
    this.url = url;
    this.verbose = verbose;
  }

  async get<T>(uri: string, local = false): Promise<T> {
    let url = this.url + uri;
    if (local) {
      url = uri;
    }
    if (this.verbose) {
      console.log("Get", url)
    }
    const response = await fetch(url, this._header("get"));
    if (!response.ok) {
      console.log("API ERROR", response)
      throw new Error();
    }
    return (await response.json()) as T;
  }

  private _header(method: string = "get"): RequestInit {
    const h = {
      method: method,
      headers: {
        "Content-Type": "application/json",
      },
      mode: "cors"
    } as RequestInit;
    return h;
  }
}

export default Api;