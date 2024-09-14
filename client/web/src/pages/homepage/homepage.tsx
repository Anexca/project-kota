import { useEffect } from "react";
import GuestHomePage from "../../componnets/shared/guest-homepage";
import useSessionStore from "../../store/auth-store";
import UserHomePage from "../../componnets/shared/user-homepage";
import useUserProfileStore from "../../store/user-info-store";
import { useToast } from "../../hooks/use-toast";

const HomePage = () => {
  const { session, loadSession, subscribeToAuthChanges } = useSessionStore();
  const { profile, getProfile } = useUserProfileStore();
  const { toast } = useToast();
  const checkSessionToken = async () => {
    try {
      if (session) {
        if (!profile.email) {
          await getProfile();
        }
        return;
      }

      const response = await loadSession();
      if (!profile.email) {
        await getProfile();
      }
      if (!response) {
        return;
      }
    } catch (error) {
      toast({
        title: "Something went wrong",
        description: "We are unable to process your request.",
        variant: "destructive",
      });
    }
  };

  useEffect(() => {
    checkSessionToken();
    return subscribeToAuthChanges();
  }, []);
  return session ? <UserHomePage /> : <GuestHomePage />;
};

export default HomePage;
