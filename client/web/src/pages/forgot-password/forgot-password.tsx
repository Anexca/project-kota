import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";
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
import {
  ForgotPasswordSchema,
  ForgotPasswordType,
} from "../../validation-schema/auth";
import {
  Alert,
  AlertTitle,
  AlertDescription,
} from "../../componnets/base/alert/alert";

export function ForgotPassword() {
  const navigate = useNavigate();
  const { loadSession } = useSessionStore();
  const { toast } = useToast();
  const { handleSubmit, control } = useForm({
    defaultValues: {
      email: "",
    },
    resolver: yupResolver(ForgotPasswordSchema),
  });
  const onSumbit = async (formData: ForgotPasswordType) => {
    const { email } = formData;

    const { data, error } = await supabase.auth.resetPasswordForEmail(email);
    if (error) {
      toast({ title: error.message || "Something went wrong." });
      return;
    }
    if (data) {
      const session = await loadSession();
      session && navigate(`/${paths.HOMEPAGE}`);
    }
  };

  return (
    <div className="w-full h-screen max-h-screen lg:grid lg:grid-cols-2 ">
      <div className="flex items-center justify-center py-12">
        <Card className="max-w-sm">
          <CardHeader>
            <CardTitle className="text-xl">Forgot Password</CardTitle>
            <CardDescription>
              Enter your email address below to get password recovery email.
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

              <Button type="submit" className="w-full">
                Send Recovery Mail
              </Button>
            </form>
            <Alert variant={"success"} className="mt-4">
              <AlertTitle>Success ✌</AlertTitle>
              <AlertDescription>
                Please check you mail box for password reset mail.
              </AlertDescription>
            </Alert>

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
