// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

import { serverUrl } from '@/const';
import { OptionsJsonDataset } from '@/models/options/types';
import { GraphQLClient, gql } from 'graphql-request'
import { classicOptionsQuery, exoticOptionsQuery } from './queries';

const graphQLClient = new GraphQLClient(serverUrl + "v0/graphql", {
  //credentials: 'include',
  mode: 'cors',
});

async function classicQuery(): Promise<OptionsJsonDataset> {
  const data = await graphQLClient.request(gql`query ${classicOptionsQuery}`);
  //console.log("Q", JSON.stringify(data, undefined, 2))
  return data as OptionsJsonDataset
}

async function exoticQuery(): Promise<OptionsJsonDataset> {
  const data = await graphQLClient.request(gql`query ${exoticOptionsQuery}`);
  //console.log("Q", JSON.stringify(data, undefined, 2))
  return data as OptionsJsonDataset
}

export { classicQuery, exoticQuery }
