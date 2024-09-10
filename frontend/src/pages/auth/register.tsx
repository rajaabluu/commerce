import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Input } from "../../components/form/form_control";
import Button from "../../components/buttons/button";
import { Link } from "react-router-dom";

const Schema = z
  .object({
    name: z.string().min(1, { message: "Name is required" }),
    email: z.string().min(1, { message: "Email is required" }),
    password: z.string().min(1, { message: "Password is required" }),
    password_confirmation: z.string(),
    address: z.string().min(1, { message: "Address must be not empty" }),
  })
  .refine(
    ({ password, password_confirmation }) => password == password_confirmation,
    { message: "Password confirmation does not match" }
  );

type Fields = z.infer<typeof Schema>;

const defaultValue: Fields = {
  name: "",
  email: "",
  password: "",
  password_confirmation: "",
  address: "",
};

const fields = [
  {
    name: "name",
    placeholder: "Name",
  },
  {
    name: "email",
    placeholder: "example@gmail.com",
  },
  {
    name: "password",
    placeholder: "Password",
  },
  {
    name: "password_confirmation",
    placeholder: "Password Confirmation",
  },
  {
    name: "address",
    placeholder: "Address",
  },
];

export default function RegisterPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Fields>({
    resolver: zodResolver(Schema),
    defaultValues: defaultValue,
  });

  return (
    <div className="flex h-screen">
      <div className="w-full sm:w-[70%] md:w-[60%] lg:w-1/2 xl:w-2/5 px-5 sm:px-6 py-8 flex flex-col items-center justify-center">
        <div className="text-center">
          <h1 className="text-indigo-600 text-3xl font-semibold">Sign Up</h1>
          <p className="text-slate-500 mt-1 px-4">
            Create your account and enjoy the shopping!
          </p>

          <form
            onSubmit={handleSubmit((data) => console.log(data))}
            className="flex flex-col gap-3 mt-12"
          >
            {fields.map((field, i) => (
              <Input
                className="px-3 py-2 rounded-md"
                key={i}
                name={field.name}
                placeholder={field.placeholder}
                type="text"
                register={register}
                error={errors[field.name as keyof Fields]}
              />
            ))}
            <Button
              type="submit"
              className="bg-indigo-600 text-white rounded-md mt-8"
              fullWidth
              variant="medium"
            >
              Submit
            </Button>
          </form>
          <p className="text-slate-500 text-sm mt-3">
            Already registered?{" "}
            <Link className="text-indigo-500" to={"/auth/login"}>
              Sign in Now
            </Link>
          </p>
        </div>
      </div>
      <div className="bg-indigo-700 flex-shrink flex-grow"></div>
    </div>
  );
}
