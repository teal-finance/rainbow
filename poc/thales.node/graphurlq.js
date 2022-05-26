import { createClient } from 'urql';
const APIURL = 'https://api.thegraph.com/subgraphs/name/thales-markets/thales-optimism'
const tokensQuery = `
  query {
    {
        markets(first: 5) {
          id
          timestamp
          creator
          currencyKey
        }
        rangedMarkets(first: 5) {
          id
          timestamp
          currencyKey
          maturityDate
        }
      }
  }
`


const client = createClient({
  url: APIURL,
})

const data = await client.query(tokensQuery).toPromise()