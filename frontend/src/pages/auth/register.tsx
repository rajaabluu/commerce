import { useForm } from "react-hook-form";
import { z } from "zod";

const Schema = z
  .object({
    name: z.string().min(1, { message: "Name is required" }),
    email: z.string().min(1, { message: "Email is required" }),
    password: z.string().min(1, { message: "Password is required" }),
    password_confirmation: z.string(),
    address: z.string().min(1),
    contact: z
      .string()
      .min(10)
      .refine((string) => !isNaN(Number(string)), {
        message: "Invalid Contact",
      }),
  })
  .refine(
    ({ password, password_confirmation }) => password == password_confirmation,
    { message: "Password confirmation does not match" }
  );

type Fields = z.infer<typeof Schema>;

export default function RegisterPage() {
  const {
    register,
    formState: { errors },
  } = useForm<Fields>();
  return (
    <div className="flex h-screen">
      <div className="w-full px-4 sm:px-6 py-8 flex flex-col items-center ">
        <div className="text-center">
          <h1 className="text-indigo-600 text-3xl font-semibold">Sign Up</h1>
          <p className="text-slate-500 mt-1">
            Create your account and enjoy the shopping!
          </p>
        </div>
      </div>
      <div className="bg-purple-700 flex-shrink flex-grow"></div>
    </div>
  );
}
