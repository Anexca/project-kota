import { PropsWithChildren, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import useSessionStore from "../store/auth-store";
import { paths } from "./route.constant";

const SignedInRoute = ({ children }: PropsWithChildren) => {
  const navigate = useNavigate();
  const { session, loadSession, subscribeToAuthChanges } = useSessionStore();
  const [isLoading, setIsloading] = useState(true);
  const checkSessionToken = async () => {
    setIsloading(true);
    if (session) {
      setIsloading(false);
      return;
    }
    const response = await loadSession();
    if (!response) {
      setIsloading(false);
      return;
    }
    navigate(`/${paths.QUESTION_PAPER}`);
  };
  useEffect(() => {
    if (session && !isLoading) {
      navigate(`/${paths.QUESTION_PAPER}`);
    }
  }, [session]);
  useEffect(() => {
    checkSessionToken();
    return subscribeToAuthChanges();
  }, []);
  return session ? null : children;
};

export default SignedInRoute;
