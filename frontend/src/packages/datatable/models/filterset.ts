import ExcludeFilter from "./exclude_filter"

export class FilterSet {
  exclude: Record<string, ExcludeFilter> = {};

  appendExcludeFilterValue(col: string, value: any) {
    if (!this._excludeFilterExists(col)) {
      //console.log("Creating exclude filter", col)
      this.exclude[col] = new ExcludeFilter()
    }
    this.exclude[col].addValue(value);
  }

  removeExcludeFilterValue(col: string, value: any) {
    if (this._excludeFilterExists(col)) {
      this.exclude[col].removeValue(value)
    }
    if (this.exclude[col].values.length == 0) {
      delete this.exclude[col]
      //console.log("Removing exclude filter", col)
    }
  }

  _excludeFilterExists(col: string): boolean {
    return Object.keys(this.exclude).includes(col)
  }
}

