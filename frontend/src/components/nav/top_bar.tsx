import { Bars2Icon, MagnifyingGlassIcon } from "@heroicons/react/24/outline";
import { Link } from "react-router-dom";

export const items = [
  {
    name: "Products",
    link: "/products",
    type: "link",
  },
  {
    name: "Hot Sales",
    link: "/#hot-sales",
    type: "link",
  },
  {
    name: "Categories",
    items: [
      {
        name: "Gaming",
        link: "/product?categories[]=gaming",
      },
      {
        name: "Gaming",
        link: "/product?categories[]=gaming",
      },
      {
        name: "Gaming",
        link: "/product?categories[]=gaming",
      },
      {
        name: "Gaming",
        link: "/product?categories[]=gaming",
      },
    ],
    type: "submenu",
  },
];

export default function TopBar() {
  return (
    <nav className="flex items-center justify-between gap-3 h-fit sticky top-0 px-4 md:px-6 lg:px-8 xl:px-10 py-3 border-b border-b-slate-200">
      <div className="flex  items-center gap-4 h-fit flex-shrink">
        <div className="flex flex-col gap-1.5 justify-center border border-slate-300 rounded-full py-1 px-1.5 size-10 cursor-pointer md:hidden">
          <Bars2Icon className="size-full" />
        </div>
        <h1 className="max-md:hidden font-medium md:font-semibold md:text-lg">
          {" "}
          <span className="text-purple-600">Flake</span>Shop
        </h1>
      </div>
      <div className="flex max-md:hidden lg:pl-14 gap-6 text-slate-600 items-center">
        {items.map(
          (link, i) =>
            link.type == "link" && (
              <Link key={i} to={link?.link as string}>
                {link.name}
              </Link>
            )
        )}
      </div>
      <div className="flex gap-4 lg:flex-grow ">
        <div className="p-2 lg:py-1.5 flex rounded-full items-center lg:gap-2 ms-auto border border-slate-300">
          <MagnifyingGlassIcon className="size-[1.35rem] lg:pl-1 lg:size-6 text-slate-500" />
          <div className="text-sm max-lg:hidden flex gap-2 items-center focus:outline-none text-slate-500 min-w-5 pr-2">
            <span>Search </span>
            <div className="flex font-semibold pl-1 gap-1.5 items-center *:border *:border-slate-300 *:bg-slate-200 *:rounded *:text-[0.65rem] *:px-1.5">
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
