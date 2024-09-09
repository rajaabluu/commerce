import axios from "axios";
import { API_BASE_URL } from "../../constant/constant";
import interceptor from "./interceptor";

const api = axios.create({ baseURL: API_BASE_URL });
interceptor.setup(api);

export default api;
