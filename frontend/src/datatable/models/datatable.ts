// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

import { reactive } from 'vue';
import { TableInterface, TableParams } from '../interfaces';
import { FilterSet } from './filters';


export default class SwDatatableModel<T = Record<string, any>> {
  state = reactive<TableInterface<T>>({ columns: {}, rows: [] });
  _initialState = reactive<TableInterface<T>>({ columns: {}, rows: [] });
  _filtersets: Record<string, FilterSet> = {};

  constructor({ columns, rows }: TableParams<T> = {}) {
    this._initialState = reactive<TableInterface<T>>({
      columns: columns ?? {},
      rows: rows ?? [],
    });
    this.state = reactive<TableInterface<T>>({
      columns: columns ?? {},
      rows: rows ?? [],
    });
  }

  get hasData(): boolean {
    return this._initialState.rows.length > 0;
  }

  filterExclude(col: string, value: any) {
    this.state.rows = this.state.rows.filter((row: Record<string, any>) => {
      if (row[col] != value) {
        return true
      }
    });
  }

  filterInclude(col: string, value: any) {
    const rows = this._initialState.rows.filter((row: Record<string, any>) => {
      if (row[col] == value) {
        return true
      }
    });
    this.state.rows = [...this.state.rows, ...rows]
  }

  setColumnsFromData(): void {
    const row = this.state.rows[0]; //eslint-disable-line
    console.log("Setting columns from row names", Object.keys(row))
    const cols = {} as Record<string, string>;
    Object.keys(row).forEach((k) => {
      let _v = k.charAt(0).toUpperCase() + k.slice(1);
      _v = _v.replace(/([A-Z])/g, " $1");
      _v = _v.replace("_", " ");
      cols[String(k)] = _v;
    });
    this.state.columns = cols;
    console.log("Cols", this.state.columns)
  }
}