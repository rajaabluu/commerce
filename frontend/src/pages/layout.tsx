import { Outlet } from "react-router-dom";
import TopBar from "../components/nav/top_bar";

export default function MainLayout() {
  return (
    <>
      <TopBar />
      <Outlet />
    </>
  );
}
