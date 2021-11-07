// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

export interface TableParams<T> {
  columns?: Record<string, string>;
  rows?: Array<T>;
  idCol?: string;
}

export interface TableInterface<T> {
  columns: Record<string, string>;
  rows: Array<T>;
}