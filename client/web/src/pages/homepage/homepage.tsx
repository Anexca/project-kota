import GuestHomePage from "../../componnets/shared/guest-homepage";
import UserHomePage from "../../componnets/shared/user-homepage";
import useSessionStore from "../../store/auth-store";

const HomePage = () => {
  const { session } = useSessionStore();

  return session ? <UserHomePage /> : <GuestHomePage />;
};

export default HomePage;
