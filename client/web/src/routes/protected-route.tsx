import { PropsWithChildren, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useToast } from "../hooks/use-toast";
import useSessionStore from "../store/auth-store";
import useUserProfileStore from "../store/user-info-store";
import { paths } from "./route.constant";

const ProtectedRoute = ({ children }: PropsWithChildren) => {
  const navigate = useNavigate();
  const { toast } = useToast();
  const { session, loadSession, subscribeToAuthChanges } = useSessionStore();
  const { profile, notificationShow, getProfile, disableNotification } =
    useUserProfileStore();

  const [isLoading, setIsloading] = useState(true);
  const updateUserProfile = async () => {
    if (!profile.phone_number && !profile.first_name && !profile.last_name) {
      try {
        const userProfile = await getProfile();
        if (
          !userProfile.phone_number &&
          !userProfile.first_name &&
          !userProfile.last_name &&
          notificationShow
        ) {
          toast({
            title: "Suggestion",
            description: "Plese complete your user profile in profile section.",
            variant: "default",
            duration: 5000,
          });
        }
      } catch (_) {
        toast({
          title: "Something went wrong",
          description:
            "Sorry we are not able to get your profile at this moment.",
          variant: "destructive",
          duration: 5000,
        });
      }
      disableNotification();
    }
  };
  const checkSessionToken = async () => {
    setIsloading(true);
    if (session) {
      setIsloading(false);
      await updateUserProfile();
      return;
    }
    const response = await loadSession();
    if (response) {
      setIsloading(false);
      await updateUserProfile();
      return;
    }
    navigate(`/${paths.LOGIN}`);
  };
  useEffect(() => {
    if (!session && !isLoading) {
      navigate(`/${paths.LOGIN}`);
    }
  }, [session, isLoading]);
  useEffect(() => {
    checkSessionToken();
    return subscribeToAuthChanges();
  }, []);
  return session ? children : null;
};

export default ProtectedRoute;
