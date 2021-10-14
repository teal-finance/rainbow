import Api from "./model";


const api = new Api({
  serverUrl: "http://teal.finance:33322",
  verbose: true,
  onError: (err) => console.log(err)
});

export default api;