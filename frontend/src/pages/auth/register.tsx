import googleLogo from "/img/google.png";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Input } from "../../components/form/form_control";
import Button from "../../components/buttons/button";
import { Link } from "react-router-dom";
import { useGoogleLogin } from "@react-oauth/google";
import { useMutation } from "@tanstack/react-query";
import api from "../../utils/api/api";
import Loader from "../../components/loader/loader";

const Schema = z
  .object({
    name: z.string().min(1, { message: "Name is required" }),
    email: z.string().min(1, { message: "Email is required" }),
    password: z.string().min(1, { message: "Password is required" }),
    password_confirmation: z.string(),
  })
  .refine(
    ({ password, password_confirmation }) => password == password_confirmation,
    {
      message: "Password confirmation does not match",
      path: ["password_confirmation"],
    }
  );

type Fields = z.infer<typeof Schema>;

const defaultValue: Fields = {
  name: "",
  email: "",
  password: "",
  password_confirmation: "",
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

  const googleLogin = useGoogleLogin({
    onSuccess: (data) => handleGoogleLogin(data),
    onError: (err) => console.log(err),
  });

  const { mutate: handleGoogleLogin, isPending: googleLoginPending } =
    useMutation({
      mutationFn: async (data: { access_token: string }) => {
        try {
          const res = await api.post(
            "/auth/google",
            {},
            {
              headers: {
                Authorization: `Bearer ${data.access_token}`,
              },
            }
          );
          if (res.status == 200) return res.data;
        } catch (err: any) {
          console.log(err.response);
          if (err.response && err.response.data)
            throw new Error(err.response.data.message);
          else throw new Error(err);
        }
      },
      onSuccess: (data) => {
        localStorage.setItem("accessToken", data.access_token);
      },
      onError: (err) => console.log(err),
    });

  return (
    <div className="flex h-screen">
      <div className="w-full sm:w-[70%] md:w-[60%] lg:w-1/2 xl:w-2/5 px-6 sm:px-8 py-8 flex flex-col items-center justify-center">
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
                className="px-3 py-1.5 !bg-slate-100 border border-slate-300 rounded-md"
                key={i}
                name={field.name}
                placeholder={field.placeholder}
                type="text"
                register={register}
                error={errors[field.name as keyof Fields]}
              />
            ))}
            <div className="flex flex-col gap-1">
              {googleLoginPending ? (
                <div className="flex justify-center mt-8">
                  <Loader className="size-8" />
                </div>
              ) : (
                <>
                  <Button
                    type="button"
                    onClick={() => googleLogin()}
                    className="rounded-md mt-8 border border-slate-300"
                    fullWidth
                    variant="medium"
                  >
                    <div className="flex justify-center gap-2.5 items-center">
                      <img src={googleLogo} className="size-5" alt="" />
                      <h1>Sign Up with Google</h1>
                    </div>
                  </Button>
                  <Button
                    type="submit"
                    className="bg-indigo-600 text-white rounded-md mt-1"
                    fullWidth
                    variant="medium"
                  >
                    Sign Up
                  </Button>
                </>
              )}
            </div>
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
