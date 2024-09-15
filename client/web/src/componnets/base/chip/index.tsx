import { cva, VariantProps } from "class-variance-authority";
import { PropsWithChildren } from "react";
import { ICONS, IconType } from "../../../constants/icons";
import { cn } from "../../../lib/utils";

const chipVariants = cva(
  "text-xs inline-flex items-center justify-center rounded-full px-2 py-0.5",
  {
    variants: {
      variant: {
        default: "bg-blue-100 text-blue-700 hover:bg-blue-200",
        destructive: "bg-red-100 text-red-700 hover:bg-red-200",
        outline:
          "border border-gray-300 bg-white text-gray-700 hover:bg-gray-100",
        secondary: "bg-gray-100 text-gray-700 hover:bg-gray-200",
        ghost: "bg-transparent text-gray-700 hover:bg-gray-100",
        link: "text-blue-700 underline-offset-4 hover:underline",
        info: "bg-sky-100 text-sky-700 hover:bg-sky-200",
        danger: "bg-red-100 text-red-700 hover:bg-red-200",
        warning: "bg-yellow-100 text-yellow-700 hover:bg-yellow-200",
        success: "bg-green-100 text-green-700 hover:bg-green-200",
      },
    },
    defaultVariants: {
      variant: "default",
    },
  }
);

const Chip = ({
  children,
  className,
  variant,
  icon,
}: PropsWithChildren &
  VariantProps<typeof chipVariants> & {
    icon: IconType;
    className?: string;
  }) => {
  return (
    <span className={cn(chipVariants({ variant }), className)}>
      <i className={`${ICONS[icon]} mr-2`}></i>
      <p className="whitespace-nowrap ">{children}</p>
    </span>
  );
};

export default Chip;
