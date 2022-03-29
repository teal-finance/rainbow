import { serverUrl } from '@/const';
import { OptionsJsonDataset } from '@/models/options/types';
import { GraphQLClient, gql } from 'graphql-request'
import { allQuery } from './queries';

const graphQLClient = new GraphQLClient(serverUrl + "v0/graphql", {
  //credentials: 'include',
  mode: 'cors',
});

async function query(): Promise<OptionsJsonDataset> {
  const data = await graphQLClient.request(gql`query ${allQuery}`);
  //console.log("Q", JSON.stringify(data, undefined, 2))
  return data as OptionsJsonDataset
}

export { query }
