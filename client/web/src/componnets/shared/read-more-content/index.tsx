import clsx from "clsx";
import { useState } from "react";

interface ReadMoreProps {
  text: string;
}

export const ReadMore = ({ text }: ReadMoreProps) => {
  const [isExpanded, setIsExpanded] = useState(true);

  return (
    <span className={""}>
      <p
        className={clsx(
          !isExpanded && "overflow-hidden whitespace-nowrap text-ellipsis"
        )}
      >
        {text}
      </p>

      <>
        <span
          className="text-info"
          role="button"
          tabIndex={0}
          aria-expanded={isExpanded}
          onClick={() => setIsExpanded(!isExpanded)}
        >
          {isExpanded ? "Show less" : "Show more"}
        </span>
      </>
    </span>
  );
};
