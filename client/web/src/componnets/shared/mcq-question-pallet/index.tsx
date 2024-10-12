import { Button } from "../../base/button/button";
import MCQButtons from "../mcq-button";
import QuestionSelector, {
  QuestionSelectorProps,
} from "../question-selector/question-selector";

type Props = QuestionSelectorProps & { handleSubmit?: () => void };

const MCQQuestionPallet = ({
  activeQuestionIndex,
  answers,
  onQuestionNumberClick,
  handleSubmit,
  sectionList,
  activeSection,
}: Props) => {
  return (
    <div className=" min-w-72 flex flex-col overflow-hidden h-full">
      <div>
        <div className="font-semibold py-2">Legends</div>
        <div className="grid grid-cols-2 gap-2">
          <MCQButtons>Unattempted</MCQButtons>
          <MCQButtons variant={"ATTEMPTED"}>Answered</MCQButtons>
          <MCQButtons variant={"FOR-REVIEW"}>Marked For Review</MCQButtons>
          <MCQButtons variant={"NOT-ANSWERED"}>Not Answered</MCQButtons>
        </div>
        <div className="font-semibold py-2">Questions</div>
      </div>
      <div className="overflow-auto flex-1">
        <QuestionSelector
          sectionList={sectionList}
          answers={answers}
          onQuestionNumberClick={onQuestionNumberClick}
          activeQuestionIndex={activeQuestionIndex}
          activeSection={activeSection}
        />
      </div>
      {handleSubmit && (
        <Button
          size={"sm"}
          variant={"success"}
          onClick={handleSubmit}
          className="w-full my-4"
        >
          Submit Exam
        </Button>
      )}
    </div>
  );
};

export default MCQQuestionPallet;
