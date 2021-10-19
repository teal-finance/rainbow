import Api from "./model";

const dns = import.meta.env.VITE_MAIN_DNS as string || "http://localhost"
const port = import.meta.env.VITE_MAIN_PORT as string || "8080"

const api = new Api({
  serverUrl: dns + ":" + port,
  verbose: true,
});

export default api;