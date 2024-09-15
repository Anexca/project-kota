import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";
import GoogleIcon from "../../assets/svg/google-icon";
import { Button } from "../../componnets/base/button/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "../../componnets/base/card/card";
import ControlledInput from "../../componnets/base/controlled-input";
import { Label } from "../../componnets/base/label/label";
import { BackgroundGradientAnimation } from "../../componnets/shared/background-blob/background-blob";
import { useToast } from "../../hooks/use-toast";
import { paths } from "../../routes/route.constant";
import useSessionStore from "../../store/auth-store";
import { supabase } from "../../supabase/client";
import { LoginSchema, LoginType } from "../../validation-schema/auth";
import useUserProfileStore from "../../store/user-info-store";

export function Login() {
  const navigate = useNavigate();
  const { loadSession } = useSessionStore();
  const { getProfile } = useUserProfileStore();
  const { toast } = useToast();
  const { handleSubmit, control } = useForm({
    defaultValues: {
      email: "",
      password: "",
    },
    resolver: yupResolver(LoginSchema),
  });
  const onSumbit = async (formData: LoginType) => {
    try {
      const { email, password } = formData;

      const { data, error } = await supabase.auth.signInWithPassword({
        email,
        password,
      });
      if (error) {
        toast({
          title: error.message || "Something went wrong.",
          variant: "destructive",
          description:
            "Sorry there is some problem in proccessing your request.",
        });
        return;
      }
      if (data) {
        await loadSession();
        await getProfile();
        navigate(`/${paths.HOMEPAGE}`);
      }
    } catch (error) {
      toast({
        title: "Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in proccessing your request.",
      });
    }
  };
  const loginWithGoogle = async () => {
    const { data, error } = await supabase.auth.signInWithOAuth({
      provider: "google",
      options: {
        redirectTo: import.meta.env.VITE_OAUTH_GOOGLE_REDIRECT_URL,
        queryParams: {
          access_type: "offline",
        },
      },
    });

    if (error) {
      toast({
        title: error.message || "Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in proccessing your request.",
      });
      return;
    }
    if (data) {
      await loadSession();
      await getProfile();
    }
  };
  return (
    <div className="w-full h-screen max-h-screen lg:grid lg:grid-cols-2 ">
      <div className="flex items-center justify-center py-12">
        <Card className="max-w-[20rem]">
          <CardHeader>
            <CardTitle className="text-xl flex">
              Login
              <Link className="ml-auto" to={`/${paths.HOMEPAGE}`}>
                <i className="fa-solid fa-xmark"></i>
              </Link>
            </CardTitle>
            <CardDescription>
              Enter your email below to login to your account
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
                  <Link
                    to="/forgot-password"
                    className="ml-auto inline-block text-sm underline"
                  >
                    Forgot your password?
                  </Link>
                </div>
                <ControlledInput
                  control={control}
                  inputProps={{
                    type: "password",
                  }}
                  name="password"
                />
              </div>
              <Button type="submit" className="w-full">
                Login
              </Button>
              <Button
                type="button"
                onClick={loginWithGoogle}
                variant="outline"
                className="w-full"
              >
                <GoogleIcon className="mr-2" />
                Login with Google
              </Button>
            </form>
            <div className="mt-4 text-center text-sm">
              Don&apos;t have an account?{" "}
              <Link to={`/${paths.REGISTER}`} className="underline">
                Sign up
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
