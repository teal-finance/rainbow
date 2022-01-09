// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

interface ApiParams {
  url: string,
  verbose: boolean,
  onError?: (response: Response) => void
}

export default ApiParams