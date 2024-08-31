import { MathJax } from "better-react-mathjax";
import { cn } from "../../../lib/utils";
import { Label } from "../../base/label/label";
import { RadioGroup, RadioGroupItem } from "../../base/radio-group/radio-group";

type Props = {
  options: { label: string; value: string }[];
  selected: string | null;
  onChange: (data: string) => void;
  name: string;
};

const AnswerOptions = ({ options, selected, onChange, name }: Props) => {
  return (
    <RadioGroup
      name={name}
      onValueChange={(a) => {
        console.log(a);
        onChange(a);
      }}
    >
      {options.map((item) => (
        <div
          className={cn(
            "flex items-center space-x-2 px-4 py-2 rounded border shadow bg-white",
            selected == item.value && "bg-slate-200"
          )}
        >
          <RadioGroupItem value={item.value} id={item.value} />
          <Label htmlFor={item.value}>
            <MathJax>{item.label}</MathJax>
          </Label>
        </div>
      ))}
    </RadioGroup>
  );
};

export default AnswerOptions;
