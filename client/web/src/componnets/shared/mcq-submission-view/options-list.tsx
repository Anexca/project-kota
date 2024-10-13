import { cn } from "../../../lib/utils";
import { Label } from "../../base/label/label";
import { RadioGroup, RadioGroupItem } from "../../base/radio-group/radio-group";
import MarkdownRender from "../markdown-rendere";

type Props = {
  options: string[];
  selected: number | null;
  correctAnswer: number;
};

const OptionList = ({ options, selected, correctAnswer }: Props) => {
  const isCorrectAnswer = selected == correctAnswer;
  return (
    <RadioGroup value={selected as any}>
      {options.map((item, index) => (
        <div
          className={cn(
            "flex items-center space-x-2 px-4 py-2 rounded shadow bg-white",
            isCorrectAnswer && "bg-success/15 color-success",
            !isCorrectAnswer &&
              correctAnswer == index &&
              "bg-success/15 color-success",
            !isCorrectAnswer &&
              selected == index &&
              "bg-red-400/15 color-red-500"
          )}
        >
          <RadioGroupItem
            checked={selected == index}
            value={`${index}`}
            id={`${index}`}
          />
          <Label htmlFor={`${index}`}>
            <MarkdownRender>{item}</MarkdownRender>
          </Label>
        </div>
      ))}
    </RadioGroup>
  );
};

export default OptionList;
