// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

const classicOptionsQuery = `{
  rows(providers: ["Delta Exchange", "Deribit", "Thalex", "Aevo", "Lyra", "Synquote", "Rysk", "SDX"]) {
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

const exoticOptionsQuery = `{
  rows(providers: ["Thales::Optimism", "Thales::Polygon","Thales::Arbitrum","Thales::Bsc"]) {
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

export { classicOptionsQuery, exoticOptionsQuery }
