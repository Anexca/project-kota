import { PropsWithChildren, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import useSessionStore from "../store/auth-store";
import { paths } from "./route.constant";

const SignedInRoute = ({ children }: PropsWithChildren) => {
  const navigate = useNavigate();
  const { session } = useSessionStore();

  useEffect(() => {
    if (session) {
      navigate(`/${paths.HOMEPAGE}`);
    }
  }, [session]);

  return session ? null : children;
};

export default SignedInRoute;
