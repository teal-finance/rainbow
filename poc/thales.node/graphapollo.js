import { ApolloClient, InMemoryCache, gql } from '@apollo/client'
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


const client = new ApolloClient({
  uri: APIURL,
  cache: new InMemoryCache(),
})

client
  .query({
    query: gql(tokensQuery),
  })
  .then((data) => console.log('Subgraph data: ', data))
  .catch((err) => {
    console.log('Error fetching data: ', err)
  })