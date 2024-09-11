import { ReactNode } from "react";
import useModal from "../../utils/hooks/useModal";
import clsx from "clsx";

export const Popover = ({
  children,
  title,
  toggleOnHover = false,
  toggleOnClick = false,
  position,
  className = "",
}: {
  title?: ReactNode;
  children?: ReactNode;
  toggleOnHover?: boolean;
  toggleOnClick?: boolean;
  position?:
    | "top-left"
    | "top-right"
    | "top-center"
    | "bottom-center"
    | "bottom-right"
    | "bottom-left"
    | undefined;
  className?: string;
}) => {
  const { isShow, toggle } = useModal();
  return (
    <div
      {...(toggleOnClick && { onClick: toggle })}
      {...(toggleOnHover && {
        onMouseEnter: () => !isShow && toggle(),
        onMouseLeave: () => isShow && toggle(),
      })}
      className={clsx("flex relative flex-col", className)}
    >
      <div>{title}</div>
      <div
        className={clsx("absolute w-max p-3", !isShow && "hidden", {
          "bottom-0 right-full": position == "top-left",
          "bottom-full left-full": position == "top-right",
          "top-full right-0 !px-0": position == "bottom-left",
          "top-full left-0 !px-0": position == "bottom-right",
          "top-full right-1/2 translate-x-1/2":
            position == "bottom-center" || position == undefined,
          "bottom-full right-1/2 translate-x-1/2": position == "top-center",
        })}
      >
        {children}
      </div>
    </div>
  );
};
