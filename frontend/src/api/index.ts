import Api from "./model";

const addr = import.meta.env.VITE_API_ADDR as string || "http://localhost"
const port = import.meta.env.VITE_API_PORT as string || "8080"

const api = new Api({
  serverUrl: port == "" ? addr : addr + ":" + port,
  verbose: true,
});

export default api;