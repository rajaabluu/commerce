import { Bars3Icon, MagnifyingGlassIcon } from "@heroicons/react/24/outline";
import { Link } from "react-router-dom";

export const items = [
  {
    name: "Products",
    link: "/products",
  },
  {
    name: "Categories",
    link: "/categories",
  },
];

export default function TopBar() {
  return (
    <nav className="flex items-center gap-3 h-fit sticky top-0 px-4 py-3 border-b border-b-slate-200">
      <div className="flex  items-center gap-4 h-fit flex-shrink">
        <Bars3Icon className="size-8 md:hidden" />
        <h1>Store</h1>
      </div>
      <div className="flex gap-3 flex-grow lg:ml-6">
        <div className="flex gap-4 text-slate-700 items-center">
          {items.map((link) => (
            <Link to={link.link}>{link.name}</Link>
          ))}
        </div>
        <div className="p-2 lg:py-1.5 flex rounded-full items-center lg:gap-3 ms-auto border border-slate-300">
          <MagnifyingGlassIcon className="size-5 lg:pl-1 lg:size-6 text-slate-500" />
          <div className="text-sm max-lg:hidden flex gap-2 items-center focus:outline-none text-slate-500 min-w-5 pr-3">
            <span>Search </span>
            <div className="flex font-medium pl-1 gap-2 items-center *:bg-slate-200 *:rounded text-xs *:px-2 *:py-0.5">
              <span>CTRL</span>
              <span>K</span>
            </div>
          </div>
        </div>
        <div className="">
          <img
            src="https://picsum.photos/200"
            className="rounded-full size-9 object-cover"
            alt=""
          />
        </div>
      </div>
    </nav>
  );
}
