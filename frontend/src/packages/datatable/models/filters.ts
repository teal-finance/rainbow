export class FilterSet {
  exclude: Record<string, Set<any>> = {};

  appendExcludeFilterValue(col: string, value: any) {
    if (this._hasExcludeFilter(col)) {
      this.exclude[col].add(value);
    } else {
      this.exclude[col] = new Set([value])
    }
  }

  removeFilterValue(col: string, value: any) {
    if (this._hasExcludeFilter(col)) {
      this.exclude[col].delete(value)
    }
  }

  _hasExcludeFilter(col: string): boolean {
    return Object.keys(this.exclude).length > 0
  }
}

