class BaseFilter {
  col: string;
  value: any;

  constructor(col: string, value: any) {
    this.col = col;
    this.value = value;
  }
}

export class IncludeFilter extends BaseFilter { }

export class ExcludeFilter extends BaseFilter { }