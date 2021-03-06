// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

const assetsPresets = {
  'BTC and ETH and SOL': {
    'defaultValue': false,
    'BTC': true,
    'ETH': true,
    'SOL': true,
  },
  'BTC and ETH': {
    'defaultValue': false,
    'BTC': true,
    'ETH': true,
  },
}

const filterPresets: Record<string, any> = {
  'All': {
    assets: { 'defaultValue': true },
    providers: { 'defaultValue': true },
  },
  'BTC, ETH, SOL all providers': {
    assets: assetsPresets['BTC and ETH and SOL'],
    providers: { 'defaultValue': true }
  },
  'BTC and ETH on Deribit': {
    assets: assetsPresets['BTC and ETH'],
    providers: { 'defaultValue': false, 'Deribit': true }
  },
  'CEXes': {
    assets: { 'defaultValue': true },
    providers: { 'defaultValue': false, 'Deribit': true, 'Delta Exchange': true }
  },
  'DEXes': {
    assets: { 'defaultValue': true },
    providers: { 'defaultValue': false, 'Lyra': true, 'Opyn': true, 'Zeta': true, 'PsyOptions': true }
  },
  'BTC': {
    assets: { 'defaultValue': false, 'BTC': true },
    providers: { 'defaultValue': true },
  },
  'ETH': {
    assets: { 'defaultValue': false, 'ETH': true },
    providers: { 'defaultValue': true },
  },
  'SOL': {
    assets: { 'defaultValue': false, 'SOL': true },
    providers: { 'defaultValue': true },
  }
}

export default filterPresets;
