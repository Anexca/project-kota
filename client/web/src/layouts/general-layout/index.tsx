import { Outlet } from "react-router-dom";
import Header from "../../componnets/shared/header/header";

const GeneralLayout = () => {
  return (
    <div className="w-full bg-neutral-50 min-h-screen">
      <Header />
      <Outlet />
    </div>
  );
};

export default GeneralLayout;
