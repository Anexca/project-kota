import clsx from "clsx";
import { MCQFormModal } from "../../../interface/mcq-exam";
import MCQButtons from "../../../pages/mcq-exam-center/mcq-buttons";

export type QuestionSelectorProps = {
  onQuestionNumberClick: (data: number) => void;
} & MCQFormModal;

const QuestionSelector = ({
  answers,
  onQuestionNumberClick,
  activeQuestionIndex,
}: QuestionSelectorProps) => {
  return (
    <div className="grid grid-cols-5 gap-2">
      {answers.map((item, index) => (
        <MCQButtons
          onClick={() => onQuestionNumberClick(index)}
          variant={item.state}
          className={clsx(
            activeQuestionIndex == index && "border-info border-4"
          )}
        >
          {index + 1}
        </MCQButtons>
      ))}
    </div>
  );
};

export default QuestionSelector;
