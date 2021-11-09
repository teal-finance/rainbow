// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

import { reactive, toRaw } from 'vue';
import { TableInterface, TableParams } from '../interfaces';
import { FilterSet } from './filterset';


export default class SwDatatableModel<T = Record<string, any>> {
  state = reactive<TableInterface<T>>({ columns: {}, rows: [] });
  _initialState = reactive<TableInterface<T>>({ columns: {}, rows: [] });
  filterset = new FilterSet();
  idCol: string | null = null;
  sortFilter: { col: string, direction: "asc" | "desc" } | null = null

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

  addSortFilterDesc(col: string) {
    this.sortFilter = { col: col, direction: "desc" };
    this.filterDesc(col)
  }

  addSortFilterAsc(col: string) {
    this.sortFilter = { col: col, direction: "asc" };
    this.filterAsc(col)
  }

  filterAsc(col: string) {
    const s = (a: any, b: any) => {
      if (a[col] < b[col]) {
        return -1;
      }
      if (a[col] > b[col]) {
        return 1;
      }
      return 0;
    }
    this.state.rows.sort(s);
  }

  filterDesc(col: string) {
    const s = (a: any, b: any) => {
      if (a[col] > b[col]) {
        return -1;
      }
      if (a[col] < b[col]) {
        return 1;
      }
      return 0;
    }
    this.state.rows.sort(s);
  }

  removeSortFilter() {
    console.log("Remove sort filter");
    this.sortFilter = null;
    this.filter();
  }

  addExcludeFilter(col: string, value: any) {
    //console.log("Add ex filter", col, value)
    this.filterset.appendExcludeFilterValue(col, toRaw(value))
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
        for (const exval of this.filterset.exclude[col].values) {
          if (row[col] == exval) {
            return false
          }
        }
      }
      return true
    });
    this.state.rows = [...rows]
    // apply sorting
    if (this.sortFilter != null) {
      if (this.sortFilter.direction == "asc") {
        this.filterAsc(this.sortFilter.col);
      } else {
        this.filterDesc(this.sortFilter.col);
      }
    }
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