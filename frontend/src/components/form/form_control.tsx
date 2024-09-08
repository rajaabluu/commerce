import { EyeIcon, EyeSlashIcon } from "@heroicons/react/24/solid";
import clsx from "clsx";
import { useState } from "react";

export function Input({
  onChange,
  className,
  register,
  name,
  value,
  type,
  error,
}: {
  type: string;
  onChange?: () => void;
  name: string;
  register?: (field: any) => any;
  value?: string;
  className?: string;
  error?: any;
}) {
  const [hide, setHide] = useState(true);
  return (
    <div className="flex flex-col">
      <div className={clsx("flex items-stretch")}>
        <input
          id={name ?? null}
          type={type}
          {...(value && { value: value })}
          {...(name && { name: name })}
          {...(onChange && { onChange: onChange })}
          {...(register && name && { ...register(name) })}
          className={clsx("bg-slate-200 focus:outline-none", className)}
        />
        {type == "password" && (
          <div
            onClick={() => setHide((hide) => !hide)}
            className="*:size-5 flex *:m-auto items-center"
          >
            {hide ? <EyeSlashIcon /> : <EyeIcon />}
          </div>
        )}
      </div>
      {error && error.message && (
        <small className="text-red-500 mt-0.5">{error.message}</small>
      )}
    </div>
  );
}
