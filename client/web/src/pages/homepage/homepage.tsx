import React from "react";
import useSessionStore from "../../store/auth-store";

type Props = {};

const HomePage = (props: Props) => {
  const { logout } = useSessionStore();
  return (
    <div>
      HomePage
      <button onClick={logout}>logout</button>
    </div>
  );
};

export default HomePage;
