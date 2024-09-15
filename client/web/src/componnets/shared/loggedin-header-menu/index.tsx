import * as React from "react";

import { cn } from "../../../lib/utils";

import { Link, useNavigate } from "react-router-dom";
import { paths } from "../../../routes/route.constant";
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
} from "../../base/navigation-menu";
import { Button } from "../../base/button/button";
import useSessionStore from "../../../store/auth-store";
import { useMediaQuery } from "../../../hooks/use-media-query";
import { ScreenSizeQuery } from "../../../constants/shared";

const ListItem = React.forwardRef<
  React.ElementRef<"a">,
  React.ComponentPropsWithoutRef<"a">
>(({ className, title, children, ...props }, ref) => {
  return (
    <li>
      <NavigationMenuLink asChild>
        <Link
          to={`${props.href}`}
          ref={ref}
          className={cn(
            "block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground",
            className
          )}
          {...props}
        >
          <div className="text-sm font-medium leading-none">{title}</div>
          <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">
            {children}
          </p>
        </Link>
      </NavigationMenuLink>
    </li>
  );
});
ListItem.displayName = "ListItem";

const LogoutButton = () => {
  const navigation = useNavigate();
  const { logout } = useSessionStore();
  const logoutHandler = async () => {
    await logout();
    navigation(`/${paths.HOMEPAGE}`);
  };
  return (
    <Button
      onClick={logoutHandler}
      className="relative flex h-9 items-center justify-center px-4 rounded-full"
    >
      <span className="capitalize relative text-sm font-semibold text-white">
        Logout
      </span>
    </Button>
  );
};
const components: JSX.Element[] = [
  <ListItem
    title="My Profile"
    href={`/${paths.PROFILE}`}
    children={"Manage and update your profile"}
  />,
  <ListItem
    title="My Plan"
    href={`/${paths.PRICING_PLAN}`}
    children={"View your plans."}
  />,
  <LogoutButton />,
];

const mobileComponents: JSX.Element[] = [
  <Link
    to={`/${paths.PROFILE}`}
    children={"My Profile"}
    className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 px-4 py-2 border border-input bg-background hover:bg-accent hover:text-accent-foreground w-full"
  />,
  <Link
    to={`/${paths.PRICING_PLAN}`}
    children={"My Plan"}
    className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 px-4 py-2 border border-input bg-background hover:bg-accent hover:text-accent-foreground w-full"
  />,
  <LogoutButton />,
];

export function NavigationHeaderMenu({ initial }: { initial: string }) {
  const isMobileView = useMediaQuery(ScreenSizeQuery.largeScreen, true);

  return isMobileView ? (
    <NavigationMenu>
      <NavigationMenuList>
        <NavigationMenuItem>
          <NavigationMenuTrigger
            noChevron
            className="relative flex h-9 w-full items-center justify-center px-4 before:absolute before:inset-0 before:rounded-full before:bg-primary before:transition before:duration-300 hover:before:scale-105 active:duration-75 active:before:scale-95 sm:w-max gap-2 !text-white"
          >
            <span className="relative w-3 text-white text-sm">{initial}</span>
          </NavigationMenuTrigger>

          <NavigationMenuContent>
            <ul className="grid w-[300px] gap-3 p-4 ">{components}</ul>
          </NavigationMenuContent>
        </NavigationMenuItem>
      </NavigationMenuList>
    </NavigationMenu>
  ) : (
    <div className="flex flex-col w-full gap-2">{mobileComponents}</div>
  );
}
