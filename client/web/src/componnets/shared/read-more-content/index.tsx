import clsx from "clsx";
import { ReactNode, useState } from "react";
import MarkdownRender from "../markdown-rendere";

interface ReadMoreProps {
  text: ReactNode;
}

export const ReadMore = ({ text }: ReadMoreProps) => {
  const [isExpanded, setIsExpanded] = useState(true);

  return (
    <span className={""}>
      <p
        className={clsx(
          !isExpanded && "h-5 overflow-hidden whitespace-nowrap text-ellipsis"
        )}
      >
        <MarkdownRender>{text}</MarkdownRender>
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
