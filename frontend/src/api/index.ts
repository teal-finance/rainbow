// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

import Api from "./model";

const addr = import.meta.env.VITE_ADDR as string || "http://localhost:8090"

const api = new Api({
  url: addr + import.meta.env.BASE_URL,
  verbose: true,
});

export default api;
