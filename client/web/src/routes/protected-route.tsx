import { PropsWithChildren, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useToast } from "../hooks/use-toast";
import useSessionStore from "../store/auth-store";
import useUserProfileStore from "../store/user-info-store";
import { paths } from "./route.constant";

const ProtectedRoute = ({ children }: PropsWithChildren) => {
  const navigate = useNavigate();
  const { toast } = useToast();
  const { session } = useSessionStore();
  const { profile, notificationShow, disableNotification } =
    useUserProfileStore();

  const showUserProfileNotification = async () => {
    if (!profile.phone_number && !profile.first_name && !profile.last_name) {
      try {
        if (notificationShow) {
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

  useEffect(() => {
    if (!session) {
      navigate(`/${paths.LOGIN}`);
    } else {
      showUserProfileNotification();
    }
  }, [session]);

  return session ? children : null;
};

export default ProtectedRoute;
