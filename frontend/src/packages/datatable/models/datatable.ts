// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

import { reactive, computed, unref } from 'vue';
import { TableInterface, TableParams } from '../interfaces';
import { FilterSet } from './filters';


export default class SwDatatableModel<T = Record<string, any>> {
  state = reactive<TableInterface<T>>({ columns: {}, rows: [] });
  _initialState = reactive<TableInterface<T>>({ columns: {}, rows: [] });
  filterset = new FilterSet();
  idCol: string | null = null;

  constructor({ columns, rows, idCol }: TableParams<T> = {}) {
    this._initialState = reactive<TableInterface<T>>({
      columns: columns ?? {},
      rows: rows ?? [],
    });
    this.state = reactive<TableInterface<T>>({
      columns: columns ?? {},
      rows: rows ?? [],
    });
    if (idCol) {
      this.idCol = idCol;
    }
  }

  get numCols(): number {
    return Object.keys(this.state.columns).length
  }

  addExcludeFilter(col: string, value: any) {
    this.filterset.appendExcludeFilterValue(col, value)
    //console.log(this.filterset.exclude)
    this.filter();
  };

  removeExcludeFilter(col: string, value: any) {
    this.filterset.removeExcludeFilterValue(col, value)
    //console.log(this.filterset.exclude)
    this.filter();
  }

  filter() {
    //console.log("FILTERSET", JSON.stringify(this.filterset, null, "  "))
    const rows = this._initialState.rows.filter((row: Record<string, any>) => {
      for (const col of Object.keys(this.filterset.exclude)) {
        //console.log("COL", col)
        for (const exval of this.filterset.exclude[col]) {
          //console.log("FVAL", values)
          //console.log(row[col], "==", exval);
          if (row[col] == exval) {
            return false
          }
        }
      }
      return true
    });
    this.state.rows = [...rows]
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