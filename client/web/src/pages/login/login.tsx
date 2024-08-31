import { Link } from "react-router-dom";
import { Label } from "../../componnets/base/label/label";
import { Button } from "../../componnets/base/button/button";
import { Input } from "../../componnets/base/input/input";
import { BackgroundGradientAnimation } from "../../componnets/shared/background-blob/background-blob";

export function Dashboard() {
  return (
    <div className="w-full h-screen max-h-screen lg:grid lg:grid-cols-2 ">
      <div className="flex items-center justify-center py-12">
        <div className="mx-auto grid w-[350px] gap-6">
          <div className="grid gap-2 text-center">
            <h1 className="text-3xl font-bold">Login</h1>
            <p className="text-balance text-muted-foreground">
              Enter your email below to login to your account
            </p>
          </div>
          <div className="grid gap-4">
            <div className="grid gap-2">
              <div className="flex items-center">
                <Label htmlFor="email">Email</Label>
              </div>
              <Input
                id="email"
                type="email"
                placeholder="m@example.com"
                required
              />
            </div>
            <div className="grid gap-2">
              <div className="flex items-center">
                <Label htmlFor="password">Password</Label>
                {/* <Link
                  to="/forgot-password"
                  className="ml-auto inline-block text-sm underline"
                >
                  Forgot your password?
                </Link> */}
              </div>
              <Input id="password" type="password" required />
            </div>
            <Button type="submit" className="w-full">
              Login
            </Button>
            <Button variant="outline" className="w-full">
              Login with Google
            </Button>
          </div>
          <div className="mt-4 text-center text-sm">
            Don&apos;t have an account?{" "}
            {/* <Link to="#" className="underline">
              Sign up
            </Link> */}
          </div>
        </div>
      </div>
      <div className="hidden bg-muted lg:block">
        {/* <img
          src="/placeholder.svg"
          alt="Image"
          width="1920"
          height="1080"
          className="h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
        /> */}
        <BackgroundGradientAnimation containerClassName="w-full h-full" />
      </div>
    </div>
  );
}
