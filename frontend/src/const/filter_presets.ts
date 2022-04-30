const assetsPresets = {
  'BTC and ETH only': {
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
  'BTC and ETH only all providers': {
    assets: assetsPresets['BTC and ETH only'],
    providers: { 'defaultValue': true }
  },
  'BTC and ETH on Opyn': {
    assets: assetsPresets['BTC and ETH only'],
    providers: { 'defaultValue': false, 'Opyn': true }
  },
  'SOL': {
    assets: { 'defaultValue': false, 'SOL': true },
    providers: { 'defaultValue': true },
  }
}

export default filterPresets;