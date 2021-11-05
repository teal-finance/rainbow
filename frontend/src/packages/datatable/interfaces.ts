export interface TableParams<T> {
  columns?: Record<string, string>;
  rows?: Array<T>;
  idCol?: string;
}

export interface TableInterface<T> {
  columns: Record<string, string>;
  rows: Array<T>;
}