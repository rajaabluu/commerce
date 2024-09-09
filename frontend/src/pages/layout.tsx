import { Outlet } from "react-router-dom";
import TopBar from "../components/nav/top_bar";

export default function MainLayout() {
  return (
    <>
      <TopBar />
      <main className="pt-4 px-4 sm:px-6 md:px-8 lg:px-10">
        <Outlet />
      </main>
    </>
  );
}
