import { GraphQLClient, gql } from 'graphql-request'
import api from '.';
import { allQuery } from './queries';

const graphQLClient = new GraphQLClient(api.url + "graphql", {
  //credentials: 'include',
  mode: 'cors',
});

async function query() {
  const data = await graphQLClient.request(gql`query ${allQuery}`);
  console.log("Q", JSON.stringify(data, undefined, 2))
}

export { query }
