// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

import Api from "./model";

const addr = import.meta.env.VITE_API_ADDR as string || "http://localhost"
const port = import.meta.env.VITE_API_PORT as string || "8080"

const api = new Api({
  serverUrl: port == "" ? addr : addr + ":" + port,
  verbose: true,
});

export default api;