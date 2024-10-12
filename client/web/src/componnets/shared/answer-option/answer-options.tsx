import { cn } from "../../../lib/utils";
import { Label } from "../../base/label/label";
import { RadioGroup, RadioGroupItem } from "../../base/radio-group/radio-group";
import MarkdownRender from "../markdown-rendere";

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
      value={selected as any}
    >
      {options.map((item, index) => (
        <div
          className={cn(
            "flex items-center space-x-2 px-4 py-2 rounded shadow bg-white",
            selected == index && "bg-info/15 color-info"
          )}
        >
          <RadioGroupItem value={`${index}`} id={`${index}`} />
          <Label htmlFor={`${index}`}>
            <MarkdownRender>{item}</MarkdownRender>
          </Label>
        </div>
      ))}
    </RadioGroup>
  );
};

export default AnswerOptions;
