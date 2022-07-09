// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

const allQuery = `{
  rows {
    date
    expiry
    provider
    asset
    strike
    call {
      bid {
        px
        size
      }
      ask {
        px
        size
      }
    }
    put {
      bid {
        px
        size
      }
      ask {
        px
        size
      }
    }
  }
}`;

export { allQuery }
