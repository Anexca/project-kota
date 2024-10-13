import clsx from "clsx";
import { memo } from "react";
import { MCQFormModal } from "../../../interface/mcq-exam";
import MCQButtons from "../mcq-button";

export type QuestionSelectorProps = {
  onQuestionNumberClick: (data: number, section: string) => void;
  sectionList: string[];
} & MCQFormModal;

const QuestionSelector = memo(
  ({
    answers,
    onQuestionNumberClick,
    activeQuestionIndex,
    sectionList,
    activeSection,
  }: QuestionSelectorProps) => {
    return sectionList.map((sections) => (
      <div className="mt-1" key={`${sections}-${activeSection}`}>
        <div className="py-2">
          <span
            className={clsx(
              "capitalize px-2 py-1 text-md font-semibold rounded",
              sections == activeSection && "bg-info/25 text-info"
            )}
          >
            {sections.toLowerCase().split("_").join(" ")}
          </span>
        </div>

        <div className="grid grid-cols-5 gap-2">
          {answers?.[sections].map((item, index) => {
            return (
              <MCQButtons
                key={`${sections}-${activeSection}-${index}`}
                onClick={() => onQuestionNumberClick(index, sections)}
                variant={item.state}
                className={clsx(
                  activeQuestionIndex == index &&
                    sections == activeSection &&
                    "border-black border-dashed border-4 rounded-3xl"
                )}
              >
                {index + 1}
              </MCQButtons>
            );
          })}
        </div>
      </div>
    ));
  }
);

export default QuestionSelector;
