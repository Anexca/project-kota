import { MathJax } from "better-react-mathjax";
import { cn } from "../../../lib/utils";
import { Label } from "../../base/label/label";
import { RadioGroup, RadioGroupItem } from "../../base/radio-group/radio-group";

type Props = {
  options: string[];
  selected: number | null;
  onChange: (data: string) => void;
  name: string;
};

const AnswerOptions = ({ options, selected, onChange, name }: Props) => {
  return (
    <RadioGroup
      name={name}
      onValueChange={(a) => {
        onChange(a);
      }}
    >
      {options.map((item, index) => (
        <div
          className={cn(
            "flex items-center space-x-2 px-4 py-2 rounded shadow bg-white",
            selected == index && "bg-info/15 color-info"
          )}
        >
          <RadioGroupItem value={item} id={item} />
          <Label htmlFor={item}>
            <MathJax>{item}</MathJax>
          </Label>
        </div>
      ))}
    </RadioGroup>
  );
};

export default AnswerOptions;
