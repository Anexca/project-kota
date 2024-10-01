import { Button } from "../../componnets/base/button/button";
import QuestionSelector, {
  QuestionSelectorProps,
} from "../../componnets/shared/question-selector/question-selector";
import MCQButtons from "./mcq-buttons";

type Props = QuestionSelectorProps;

const QuestionSelectorSection = ({
  activeQuestionIndex,
  answers,
  onQuestionNumberClick,
}: Props) => {
  return (
    <div className=" min-w-72">
      <div className="font-semibold py-2">Legends</div>
      <div className="grid grid-cols-2 gap-2">
        <MCQButtons>Unattempted</MCQButtons>
        <MCQButtons variant={"ATTEMPTED"}>Answered</MCQButtons>
        <MCQButtons variant={"FOR-REVIEW"}>Marked For Review</MCQButtons>
        <MCQButtons variant={"NOT-ANSWERED"}>Not Answered</MCQButtons>
      </div>
      <div className="font-semibold py-2">Questions</div>

      <QuestionSelector
        answers={answers}
        onQuestionNumberClick={onQuestionNumberClick}
        activeQuestionIndex={activeQuestionIndex}
      />
      <div className="py-2 flex items-center justify-end">
        <Button variant={"destructive"}>Exit Exam</Button>
      </div>
    </div>
  );
};

export default QuestionSelectorSection;
