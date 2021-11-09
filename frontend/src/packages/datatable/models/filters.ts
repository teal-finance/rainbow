export class FilterSet {
  exclude: Record<string, Set<any>> = {};

  appendExcludeFilterValue(col: string, value: any) {
    if (this._hasExcludeFilter(col)) {
      this.exclude[col].add(value);
    } else {
      this.exclude[col] = new Set([value])
    }
  }

  removeExcludeFilterValue(col: string, value: any) {
    if (this._hasExcludeFilter(col)) {
      this.exclude[col].delete(value)
    }
  }

  _hasExcludeFilter(col: string): boolean {
    return col in Object.keys(this.exclude)
  }
}

