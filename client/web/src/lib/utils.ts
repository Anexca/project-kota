import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const delay = async (interval: number) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve("");
    }, interval);
  });
};
