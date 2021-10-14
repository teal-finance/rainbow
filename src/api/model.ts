import ApiParams from "./interface";


class Api {
  serverUrl: string;
  verbose: boolean;
  onError: (response: Response) => void;
  private csrfToken = "";
  private credentials: RequestCredentials = "include";

  constructor({ serverUrl, verbose, onError }: ApiParams) {
    this.serverUrl = serverUrl;
    this.verbose = verbose;
    this.onError = onError ?? function (resp) { console.log("Api error response: " + resp) }
  }

  set newCsrfToken(t: string) {
    this.csrfToken = t;
  }

  async post<T>(uri: string, payload: Array<any> | Record<string, any> | FormData, multipart: boolean = false) { // eslint-disable-line
    const opts = this._postPutHeader(payload, "post", multipart);
    let url = this.serverUrl + uri;
    if (this.verbose) {
      console.log("Post to", url, opts.body?.toString())
      console.log(JSON.stringify(opts, null, "  "))
    }
    const response = await fetch(url, opts);
    if (!response.ok) {
      this.onError(response);
      throw new Error(response.statusText);
    }
    return (await response.json()) as T;
  }

  async patch<T>(uri: string, payload: Array<any> | Record<string, any>) { // eslint-disable-line
    const opts = this._postPutHeader(payload, "patch");
    let url = this.serverUrl + uri;
    if (this.verbose) {
      console.log("Patch to", url)
      console.log(JSON.stringify(opts, null, "  "))
    }
    const response = await fetch(url, opts);
    if (!response.ok) {
      this.onError(response);
      throw new Error(response.statusText);
    }
    return (await response.json()) as T;
  }

  async put<T>(uri: string, payload: Array<any> | Record<string, any>) { // eslint-disable-line
    const opts = this._postPutHeader(payload, "put");
    let url = this.serverUrl + uri;
    if (this.verbose) {
      console.log("Put to", url)
      console.log(JSON.stringify(opts, null, "  "))
    }
    const response = await fetch(url, opts);
    if (!response.ok) {
      this.onError(response);
      throw new Error(response.statusText);
    }
    return (await response.json()) as T;
  }

  async get<T>(uri: string): Promise<T> {
    let url = this.serverUrl + uri;
    if (this.verbose) {
      console.log("Get", url)
    }
    const response = await fetch(url, this._header("get"));
    if (!response.ok) {
      this.onError(response);
      throw new Error(response.statusText);
    }
    return (await response.json()) as T;
  }

  async delete(uri: string): Promise<void> { // eslint-disable-line
    const response = await fetch(this.serverUrl + uri, this._header("delete"));
    if (!response.ok) {
      this.onError(response);
      throw new Error(response.statusText);
    }
  }

  private _header(method: string = "get"): RequestInit {
    const h = {
      method: method,
      headers: {
        "Content-Type": "application/json",
      },
      mode: "no-cors"
    } as RequestInit;
    if (this.credentials !== null) {
      h.credentials = this.credentials as RequestCredentials;
    }
    if (this.csrfToken !== null) {
      h.headers = {
        "Content-Type": "application/json",
        "X-CSRFToken": this.csrfToken
      }
    }
    return h;
  }

  private _postPutHeader(payload: Array<any> | Record<string, any> | FormData, method = "post", multipart: boolean = false): RequestInit { // eslint-disable-line
    const pl = multipart ? payload as FormData : JSON.stringify(payload);
    const r: RequestInit = {
      method: method,
      mode: "cors",
      body: pl
    };
    if (!multipart) {
      r.headers = {
        "Content-Type": "application/json"
      }
    }
    if (this.credentials !== null) {
      r.credentials = this.credentials as RequestCredentials
    }
    if (this.csrfToken !== null) {
      if (multipart) {
        r.headers = {
          "X-CSRFToken": this.csrfToken
        }
      } else {
        r.headers = {
          "Content-Type": "application/json",
          "X-CSRFToken": this.csrfToken
        }
      }

    }
    return r;
  }
}

export default Api;