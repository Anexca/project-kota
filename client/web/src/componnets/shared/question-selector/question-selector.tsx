import { Button } from "../../base/button/button";

type Props = {
  questions: any[];
  onQuestionNumberClick: (data: number) => void;
};

const QuestionSelector = ({ questions, onQuestionNumberClick }: Props) => {
  return (
    <div className="grid grid-cols-5 gap-2">
      {questions.map((_, index) => (
        <Button
          onClick={() => onQuestionNumberClick(index)}
          className="shadow !p-0"
          variant={"ghost"}
        >
          <span className="p-2 bg-white w-full rounded">{index + 1}</span>
        </Button>
      ))}
    </div>
  );
};

export default QuestionSelector;
