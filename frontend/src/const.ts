const addr = import.meta.env.VITE_ADDR as string || "http://localhost:8090"
const serverUrl = addr + import.meta.env.BASE_URL

export { serverUrl }