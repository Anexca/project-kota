import clsx from "clsx";
import { useMemo } from "react";
import { Link } from "react-router-dom";

type Props = { currentTime: number; active: boolean };

const TestHeader = ({ currentTime, active }: Props) => {
  const time = useMemo(() => {
    const seconds = Math.floor(currentTime % 60);
    const minutes = Math.floor(currentTime / 60);
    return `${minutes}:${seconds}`;
  }, [currentTime]);
  return (
    <header className="sticky top-0 z-10 ">
      <nav className="z-10 w-full border-b border-black/5 dark:border-white/5  bg-white/30 backdrop-blur-md">
        <div className="max-w-7xl mx-auto px-6 md:px-12 xl:px-6">
          <div className="relative flex flex-wrap items-center justify-between gap-6 py-3 md:gap-0 md:py-4">
            <div className="relative z-20 flex w-full justify-between md:px-0 lg:w-max">
              <Link
                to="/home"
                aria-label="logo"
                className="flex items-center space-x-2"
              >
                <div aria-hidden="true" className="flex space-x-1">
                  <div className="h-4 w-4 rounded-full bg-gray-900 dark:bg-white"></div>
                  <div className="h-6 w-2 bg-primary"></div>
                </div>
                <span className="text-2xl font-bold text-gray-900 dark:text-white">
                  Revaluate
                </span>
              </Link>

              <div
                className={clsx(
                  "relative flex max-h-10 items-center lg:hidden",
                  active && "text-orange-500"
                )}
              >
                {time} Left
              </div>
            </div>
            <div
              id="navLayer"
              aria-hidden="true"
              className={clsx(
                "fixed inset-0 z-10 h-screen w-screen origin-bottom scale-y-0 bg-white/70 backdrop-blur-2xl transition duration-500 dark:bg-gray-900/70 lg:hidden"
              )}
            ></div>
            <div
              id="navlinks"
              className={clsx(
                "invisible absolute top-full left-0 z-20 w-full origin-top-right translate-y-1 scale-90 flex-col flex-wrap justify-end gap-6 rounded-3xl border border-gray-100 bg-white p-8 opacity-0 shadow-2xl shadow-gray-600/10 transition-all duration-300 dark:border-gray-700 dark:bg-gray-800 dark:shadow-none lg:visible lg:relative lg:flex lg:w-7/12 lg:translate-y-0 lg:scale-100 lg:flex-row lg:items-center lg:gap-0 lg:border-none lg:bg-transparent lg:p-0 lg:opacity-100 lg:shadow-none",
                active && "text-orange-500"
              )}
            >
              {time} Left
            </div>
          </div>
        </div>
      </nav>
    </header>
  );
};

export default TestHeader;
