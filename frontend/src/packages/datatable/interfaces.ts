// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

export interface TableParams<T> {
  columns?: Record<string, string>;
  rows?: Array<T>;
  idCol?: string;
}

export interface TableInterface<T> {
  columns: Record<string, string>;
  rows: Array<T>;
}
