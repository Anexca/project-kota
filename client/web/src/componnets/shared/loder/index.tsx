import { cva, VariantProps } from "class-variance-authority";

const loaderStyles = cva(
  "border-t-transparent border-solid animate-spin rounded-full", // common styles for the loader
  {
    variants: {
      size: {
        small: "w-4 h-4 border-2",
        medium: "w-8 h-8 border-4",
        large: "w-12 h-12 border-8",
      },
      color: {
        default: "border-primary",
        destructive: "border-destructive",
        outline: "border-input",
        secondary: "border-secondary",
        ghost: "border-accent",
        link: "border-primary",
        info: "border-info",
        danger: "border-danger",
        warning: "border-warning",
        success: "border-success",
      },
    },
    defaultVariants: {
      size: "small",
      color: "default",
    },
  }
);

const Loader = ({
  size = "small",
  color = "default",
}: VariantProps<typeof loaderStyles>) => {
  return <div className={loaderStyles({ size, color })}></div>;
};

export default Loader;
