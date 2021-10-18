class BaseFilter {
  col: string;
  value: any;

  constructor(col: string, value: any) {
    this.col = col;
    this.value = value;
  }
}

export class ExcludeFilter extends BaseFilter { }

export class FilterSet {
  col: string;
  exclude: Set<ExcludeFilter>;

  constructor(col: string, exclude: Set<ExcludeFilter>) {
    this.col = col;
    this.exclude = exclude;
  }
}

