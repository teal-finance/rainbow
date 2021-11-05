// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

import ApiParams from "./interface";

class Api {
  serverUrl: string;
  verbose: boolean;

  constructor({ serverUrl, verbose, onError }: ApiParams) {
    this.serverUrl = serverUrl;
    this.verbose = verbose;
  }

  async get<T>(uri: string): Promise<T> {
    let url = this.serverUrl + uri;
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