import { PropsWithChildren, useEffect, useState } from "react";
import { useToast } from "../../../hooks/use-toast";
import useSessionStore from "../../../store/auth-store";
import useUserProfileStore from "../../../store/user-info-store";
import Loader from "../loder";

const SessionProvider = ({ children }: PropsWithChildren) => {
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
    try {
      setIsloading(true);
      if (session) {
        await updateUserProfile();
        setIsloading(false);
        return;
      }
      const response = await loadSession();
      if (response) {
        await updateUserProfile();
      }
      setIsloading(false);
    } catch (error) {
      toast({
        title: "Something went wrong",
        description:
          "Sorry we are not able to get your profile at this moment.",
        variant: "destructive",
        duration: 5000,
      });
    } finally {
      setIsloading(false);
    }
  };

  useEffect(() => {
    checkSessionToken();
    return subscribeToAuthChanges();
  }, []);
  if (isLoading)
    return (
      <div className="h-screen flex items-center justify-center">
        <Loader size={"large"} color={"info"} />
      </div>
    );
  return children;
};

export default SessionProvider;
