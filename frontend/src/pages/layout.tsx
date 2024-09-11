import { Outlet } from "react-router-dom";
import Navbar from "../components/nav/navbar";

export default function MainLayout() {
  return (
    <>
      <Navbar />
      <main className="pt-4 px-4 sm:px-6 md:px-8 lg:px-10">
        <Outlet />
      </main>
    </>
  );
}
