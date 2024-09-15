"use client";

import React, { forwardRef, useRef } from "react";

import { cn } from "../../../lib/utils";
import { AnimatedBeam } from "../animated-beam/idex";

const Circle = forwardRef<
  HTMLDivElement,
  { className?: string; children?: React.ReactNode }
>(({ className, children }, ref) => {
  return (
    <div
      ref={ref}
      className={cn(
        "z-10 flex rounded items-center justify-center border-2 bg-white px-3 py-1 font-semibold text-green-600 shadow-[0_0_20px_-12px_rgba(0,0,0,0.8)]",
        className
      )}
    >
      {children}
    </div>
  );
});

export function AnimatedBeamHeading() {
  const containerRef = useRef<HTMLDivElement>(null);
  const div1Ref = useRef<HTMLDivElement>(null);
  const div2Ref = useRef<HTMLDivElement>(null);
  const div3Ref = useRef<HTMLDivElement>(null);
  const div4Ref = useRef<HTMLDivElement>(null);

  return (
    <div
      className="relative flex w-full max-w-[500px] items-center justify-center overflow-hidden py-10 md:p-10  "
      ref={containerRef}
    >
      <div className="flex h-full w-full flex-col items-stretch justify-between gap-10">
        <div className="flex flex-row justify-between">
          <Circle ref={div1Ref}>Study</Circle>

          <Circle ref={div2Ref}>Practice</Circle>

          <Circle ref={div3Ref}>Grow</Circle>

          <Circle ref={div4Ref}>Success</Circle>
        </div>
      </div>

      <AnimatedBeam
        containerRef={containerRef}
        fromRef={div1Ref}
        toRef={div2Ref}
        startYOffset={10}
        endYOffset={10}
        curvature={100}
        delay={0.5}
      />
      <AnimatedBeam
        delay={0.5}
        containerRef={containerRef}
        fromRef={div2Ref}
        toRef={div3Ref}
        startYOffset={10}
        endYOffset={10}
        curvature={-40}
      />
      <AnimatedBeam
        delay={0.5}
        containerRef={containerRef}
        fromRef={div3Ref}
        toRef={div4Ref}
        startYOffset={10}
        endYOffset={10}
        curvature={100}
      />
    </div>
  );
}
