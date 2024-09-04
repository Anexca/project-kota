import clsx from "clsx";
import { PropsWithChildren } from "react";
type Props = {
  className?: string;
};

const Container = ({ className, children }: PropsWithChildren<Props>) => {
  return (
    <div
      className={clsx(
        "w-full md:max-w-md lg:max-w-lg xl:max-w-xl 2xl:max-w-2xl mx-auto",
        className
      )}
    >
      {children}
    </div>
  );
};

export default Container;
