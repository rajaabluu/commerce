import { Axios } from "axios";

const interceptor = {
  setup: (api: Axios) => {
    api.interceptors.request.use(
      (config) => {
        const token = localStorage.getItem("accessToken");
        if (!!token) config.headers.Authorization = `Bearer ${token}`;
        return config;
      },
      (err) => {
        throw new Error(err);
      }
    );
  },
};

export default interceptor;
