// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

interface ApiParams {
  url: string,
  verbose: boolean,
  onError?: (response: Response) => void
}

export default ApiParams
