import clsx from "clsx";
import { ReactNode } from "react";

export default function Button({
  variant = "medium",
  onClick,
  fullWidth = false,
  children,
  type = "button",
  className = "",
}: {
  onClick?: () => void;
  children?: ReactNode;
  variant?: "small" | "medium" | "large";
  fullWidth?: boolean;
  className?: string;
  type?: "button" | "submit" | "reset" | undefined;
}) {
  return (
    <button
      onClick={onClick}
      type={type}
      className={clsx(
        {
          "w-full": fullWidth,
          "py-1 px-2": variant == "small",
          "py-2 px-4": variant == "medium",
          "py-3 px-6": variant == "large",
        },
        className
      )}
    >
      {children}
    </button>
  );
}
