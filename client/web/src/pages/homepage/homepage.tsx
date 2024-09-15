import { useEffect, useState } from "react";
import GuestHomePage from "../../componnets/shared/guest-homepage";
import useSessionStore from "../../store/auth-store";
import UserHomePage from "../../componnets/shared/user-homepage";
import useUserProfileStore from "../../store/user-info-store";
import { useToast } from "../../hooks/use-toast";
import Loader from "../../componnets/shared/loder";

const HomePage = () => {
  const { session, loadSession, subscribeToAuthChanges } = useSessionStore();
  const { profile, getProfile } = useUserProfileStore();
  const { toast } = useToast();
  const [loading, setLoading] = useState(true);
  const checkSessionToken = async () => {
    setLoading(true);
    try {
      if (session) {
        if (!profile.email) {
          await getProfile();
        }
        setLoading(false);
        return;
      }

      const response = await loadSession();
      if (!profile.email) {
        await getProfile();
      }
      setLoading(false);
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
    setLoading(false);
  };

  useEffect(() => {
    checkSessionToken();
    return subscribeToAuthChanges();
  }, []);
  if (loading)
    return (
      <div className="h-screen flex items-center justify-center">
        <Loader size={"large"} color={"info"} />
      </div>
    );
  return session ? <UserHomePage /> : <GuestHomePage />;
};

export default HomePage;
