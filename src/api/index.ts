import Api from "./model";


const api = new Api({
  serverUrl: "http://teal.finance:33322",
  //serverUrl: "http://localhost:8080",
  verbose: true,
});

export default api;