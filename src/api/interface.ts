interface ApiParams {
  serverUrl: string,
  verbose: boolean,
  onError?: (response: Response) => void
}

export default ApiParams