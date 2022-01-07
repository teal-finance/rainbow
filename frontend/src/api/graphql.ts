import { GraphQLClient, gql } from 'graphql-request'
import { allQuery } from './queries';

const addr = import.meta.env.VITE_ADDR as string || "http://localhost:8090";

const graphQLClient = new GraphQLClient(addr, {
  //credentials: 'include',
  mode: 'cors',
});

async function query() {
  const q = await graphQLClient.request(addr + "/graphql", gql`query ${allQuery}`);
  console.log("Q", q)
}

export { query }