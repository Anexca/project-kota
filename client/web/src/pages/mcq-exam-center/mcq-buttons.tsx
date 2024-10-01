import { cva, VariantProps } from "class-variance-authority";
import React from "react";
import { Button } from "../../componnets/base/button/button";
import { cn } from "../../lib/utils";
import { QUESTION_STATE } from "../../constants/shared";

const buttonVariants = cva("p-2 text-center w-full rounded h-full", {
  variants: {
    variant: {
      [`${QUESTION_STATE.UN_ATTEMPTED}`]: "bg-white",
      [`${QUESTION_STATE.ATTEMPTED}`]: "bg-success text-white",
      [`${QUESTION_STATE.FOR_REVIEW}`]: "bg-warning text-white",
      [`${QUESTION_STATE.NOT_ANSWERED}`]: "bg-primary text-white",
    },
  },
  defaultVariants: {
    variant: `${QUESTION_STATE.UN_ATTEMPTED}`,
  },
});

interface MCQButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonVariants> {}

const MCQButtons = React.forwardRef<HTMLButtonElement, MCQButtonProps>(
  ({ className, variant, children, ...props }, ref) => {
    return (
      <Button {...props} className="shadow !p-0" variant={"ghost"} ref={ref}>
        <span className={cn(buttonVariants({ variant, className }))}>
          {children}
        </span>
      </Button>
    );
  }
);

export default MCQButtons;
