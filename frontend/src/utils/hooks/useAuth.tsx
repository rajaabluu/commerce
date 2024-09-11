import { useContext } from "react";
import { AuthContext } from "../../context/auth/context";

export default function useAuth() {
  const context = useContext(AuthContext);
  if (!context) throw new Error("No Auth Context Provided");
  return context;
}
