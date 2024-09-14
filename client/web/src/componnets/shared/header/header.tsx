import clsx from "clsx";
import { useState } from "react";
import { Link, NavLink } from "react-router-dom";
import { paths } from "../../../routes/route.constant";
import useSessionStore from "../../../store/auth-store";
import { NavigationHeaderMenu } from "../loggedin-header-menu";
import { cn } from "../../../lib/utils";
import useUserProfileStore from "../../../store/user-info-store";

const loggedOutLinks = [
  {
    to: `/${paths.PRICING_PLAN}`,
    label: "Pricing",
  },
];
const loggedInLinks = [
  {
    to: `/${paths.EXAMS}/${paths.MY_SUMBISSIONS}`,
    label: "My Submissions",
  },
];
const Header = () => {
  const [mobileViewHeader, setMobileViewHeader] = useState(false);
  const { session } = useSessionStore();
  const { profile } = useUserProfileStore();
  const links = session ? loggedInLinks : loggedOutLinks;

  return (
    <header className="sticky top-0 z-10 ">
      <nav className="z-10 w-full border-b border-black/5 dark:border-white/5  bg-white/30 backdrop-blur-md">
        <div className="max-w-7xl mx-auto px-4 md:px-4 xl:px-4">
          <div className="relative flex flex-wrap items-center justify-between gap-6 py-2 md:gap-0 ">
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
                  PseudoTest
                </span>
              </Link>

              <div className="relative flex max-h-10 items-center lg:hidden">
                <button
                  onClick={() => setMobileViewHeader(!mobileViewHeader)}
                  aria-label="humburger"
                  id="hamburger"
                  className="relative  p-6"
                >
                  <div
                    aria-hidden="true"
                    id="line"
                    className={clsx(
                      "m-auto h-0.5 w-5 rounded bg-sky-900 transition duration-300 dark:bg-gray-300",
                      mobileViewHeader && "translate-y-1 rotate-45"
                    )}
                  ></div>
                  <div
                    aria-hidden="true"
                    id="line2"
                    className={clsx(
                      "m-auto mt-2 h-0.5 w-5 rounded bg-sky-900 transition duration-300 dark:bg-gray-300",
                      mobileViewHeader && "-translate-y-1.5 -rotate-45"
                    )}
                  ></div>
                </button>
              </div>
            </div>
            <div
              id="navLayer"
              aria-hidden="true"
              className={clsx(
                "fixed inset-0 z-10 h-screen  origin-bottom scale-y-0 bg-white/70 backdrop-blur-2xl transition duration-500 dark:bg-gray-900/70 lg:hidden",
                mobileViewHeader && "origin-top scale-y-100"
              )}
            ></div>
            <div
              id="navlinks"
              className={clsx(
                "invisible absolute top-full left-0 z-20 w-full origin-top-right translate-y-1 scale-90 flex-col flex-wrap justify-end gap-6 rounded-3xl border border-gray-100 bg-white p-8 opacity-0 shadow-2xl shadow-gray-600/10 transition-all duration-300 dark:border-gray-700 dark:bg-gray-800 dark:shadow-none lg:visible lg:relative lg:flex lg:w-7/12 lg:translate-y-0 lg:scale-100 lg:flex-row lg:items-center lg:gap-0 lg:border-none lg:bg-transparent lg:p-0 lg:opacity-100 lg:shadow-none",
                mobileViewHeader &&
                  "!visible !scale-100 !opacity-100 !lg:translate-y-0"
              )}
            >
              <div className="w-full text-gray-600 dark:text-gray-200 lg:w-auto lg:pr-4 lg:pt-0">
                <ul className="flex flex-col gap-6 tracking-wide lg:flex-row lg:gap-0 lg:text-sm">
                  {links.map((link) => (
                    <li>
                      <NavLink
                        to={link.to}
                        className={({ isActive }) =>
                          cn(
                            "hover:text-primary block transition dark:hover:text-white md:px-4 text-center",
                            isActive && "bg-neutral-200/50 p-2 rounded-full"
                          )
                        }
                      >
                        <span>{link.label}</span>
                      </NavLink>
                    </li>
                  ))}
                </ul>
              </div>

              <div className="mt-12 lg:mt-0 flex flex-col sm:flex-row gap-2">
                {session ? (
                  <NavigationHeaderMenu
                    initial={profile.email?.[0]?.toUpperCase()}
                  />
                ) : (
                  <Link
                    to={`/${paths.REGISTER}`}
                    className="relative flex h-9 w-full items-center justify-center px-4 before:absolute before:inset-0 before:rounded-full before:bg-primary before:transition before:duration-300 hover:before:scale-105 active:duration-75 active:before:scale-95 sm:w-max"
                  >
                    <span className="relative text-sm font-semibold text-white">
                      Sign Up
                    </span>
                  </Link>
                )}
              </div>
            </div>
          </div>
        </div>
      </nav>
    </header>
  );
};

export default Header;
