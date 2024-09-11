import { useQuery } from "@tanstack/react-query";
import { createContext } from "react";
import api from "../../utils/api/api";
import { AxiosError } from "axios";

export const AuthContext = createContext<{ auth: any; authLoading: boolean }>({
  auth: null,
  authLoading: false,
});

export const AuthContextProvider = () => {
  const { data: auth, isLoading: authLoading } = useQuery({
    queryKey: ["auth"],
    queryFn: async () => {
      try {
        const res = await api.get("/auth/me");
        if (res.status == 200) return res.data;
      } catch (err: any) {
        if (err instanceof AxiosError)
          throw new Error(err.response?.data.message);
        else throw new Error(err);
      }
    },
  });

  return (
    <AuthContext.Provider
      value={{ auth: auth, authLoading: authLoading }}
    ></AuthContext.Provider>
  );
};
