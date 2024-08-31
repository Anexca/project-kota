import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "../../componnets/base/card/card";
import { Label } from "../../componnets/base/label/label";

import { Button } from "../../componnets/base/button/button";
import { BackgroundGradientAnimation } from "../../componnets/shared/background-blob/background-blob";
import { paths } from "../../routes/route.constant";
import { Link } from "react-router-dom";
import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import ControlledInput from "../../componnets/base/controlled-input";
import { useToast } from "../../hooks/use-toast";
import useSessionStore from "../../store/auth-store";
import { LoginSchema, LoginType } from "../../validation-schema/auth";
import { supabase } from "../../supabase/client";
import GoogleIcon from "../../assets/svg/google-icon";

export function RegisterPage() {
  const { setSession } = useSessionStore();
  const { toast } = useToast();
  const { handleSubmit, control } = useForm({
    defaultValues: {
      email: "",
      password: "",
    },
    resolver: yupResolver(LoginSchema),
  });
  const onSumbit = async (formData: LoginType) => {
    const { email, password } = formData;
    const { data, error } = await supabase.auth.signInWithPassword({
      email,
      password,
    });
    if (error) {
      toast({ title: error.message || "Something went wrong." });
      return;
    }
    if (data) {
      toast({ title: "Successfully logged in." });
      setSession(data.session);
    }
  };
  return (
    <div className="w-full h-screen max-h-screen lg:grid lg:grid-cols-2 ">
      <div className="flex items-center justify-center py-12">
        <Card className="max-w-[20rem]">
          <CardHeader>
            <CardTitle className="text-xl flex">
              Register
              <Link className="ml-auto" to={`/${paths.HOMEPAGE}`}>
                <i className="fa-solid fa-xmark"></i>
              </Link>
            </CardTitle>
            <CardDescription>
              Enter your information to create an account
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form
              onSubmit={handleSubmit(onSumbit)}
              noValidate
              className="grid gap-4"
            >
              <div className="grid gap-2">
                <div className="flex items-center">
                  <Label htmlFor="email">Email</Label>
                </div>
                <ControlledInput
                  control={control}
                  name="email"
                  inputProps={{
                    placeholder: "m@example.com",
                  }}
                />
              </div>
              <div className="grid gap-2">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
                </div>
                <ControlledInput
                  control={control}
                  inputProps={{ type: "password" }}
                  name="password"
                />
              </div>
              <Button type="submit" className="w-full">
                Register
              </Button>
              <Button variant="outline" className="w-full">
                <GoogleIcon className="mr-2" /> Login with Google
              </Button>
            </form>
            <div className="mt-4 text-center text-sm">
              Already have an account?{" "}
              <Link to={`/${paths.LOGIN}`} className="underline">
                Login
              </Link>
            </div>
          </CardContent>
        </Card>
      </div>
      <div className="hidden bg-muted lg:block">
        <BackgroundGradientAnimation containerClassName="w-full h-full" />
      </div>
    </div>
  );
}
