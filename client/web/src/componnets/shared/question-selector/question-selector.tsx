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
          className="p-2 shadow"
          variant={"secondary"}
        >
          {index + 1}
        </Button>
      ))}
    </div>
  );
};

export default QuestionSelector;
