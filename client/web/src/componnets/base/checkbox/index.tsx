import * as React from "react";
import * as CheckboxPrimitive from "@radix-ui/react-checkbox";

import { cn } from "../../../lib/utils";
import { cva, VariantProps } from "class-variance-authority";

const checkboxVariants = cva(
  "peer h-4 w-4 shrink-0 rounded-sm border border-primary ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 text-xs",
  {
    variants: {
      variant: {
        default:
          "data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground",

        success:
          "data-[state=checked]:bg-success data-[state=checked]:text-success-foreground",

        info: "data-[state=checked]:bg-info data-[state=checked]:text-info-foreground",

        warning:
          "data-[state=checked]:bg-warning data-[state=checked]:text-warning-foreground",

        danger:
          "data-[state=checked]:bg-danger data-[state=checked]:text-danger-foreground",

        outline:
          "data-[state=checked]:border-primary data-[state=checked]:text-primary",

        secondary:
          "data-[state=checked]:bg-secondary data-[state=checked]:text-secondary-foreground",

        ghost:
          "data-[state=checked]:bg-accent data-[state=checked]:text-accent-foreground",

        destructive:
          "data-[state=checked]:bg-destructive data-[state=checked]:text-destructive-foreground",
      },
      //   size: {
      //     default: "h-10 px-4 py-2",
      //     sm: "h-9 rounded-md px-3 text-xs",
      //     lg: "h-11 rounded-md px-8",
      //     icon: "h-10 w-10",
      //   },
    },
    defaultVariants: {
      variant: "default",
      //   size: "default",
    },
  }
);
export interface CheckboxProps
  extends React.ComponentPropsWithoutRef<typeof CheckboxPrimitive.Root>,
    VariantProps<typeof checkboxVariants> {}

const Checkbox = React.forwardRef<
  React.ElementRef<typeof CheckboxPrimitive.Root>,
  CheckboxProps
>(({ className, variant, ...props }, ref) => (
  <CheckboxPrimitive.Root
    ref={ref}
    className={cn(checkboxVariants({ variant }), className)}
    {...props}
  >
    <CheckboxPrimitive.Indicator
      className={cn("flex items-center justify-center text-current")}
    >
      <i className="fa-solid fa-check  text-xs"></i>
    </CheckboxPrimitive.Indicator>
  </CheckboxPrimitive.Root>
));
Checkbox.displayName = CheckboxPrimitive.Root.displayName;

export { Checkbox };
