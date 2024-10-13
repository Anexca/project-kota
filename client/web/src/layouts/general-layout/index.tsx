import { Outlet } from "react-router-dom";
import Header from "../../componnets/shared/header/header";
import { PropsWithChildren } from "react";

const GeneralLayout = ({ children }: PropsWithChildren) => {
  return (
    <div className="w-full bg-neutral-50 min-h-screen ">
      <Header />
      {children || <Outlet />}
    </div>
  );
};

export default GeneralLayout;
