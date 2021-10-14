export interface TableParams<T> {
  columns?: Record<string, string>;
  rows?: Array<T>;
}

export interface TableInterface<T> {
  columns: Record<string, string>;
  rows: Array<T>;
}