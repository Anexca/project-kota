import { Button } from "../../base/button/button";
import Chip from "../../base/chip";
import Icon from "../../base/icon";

type Props = {
  srNumber?: number | null;
  topic: string;
  type: string;
  isAttemped?: boolean;
  handleAttemptClick?: () => void;
  showSubmission?: () => void;
  duration?: number;
  attempts: number;
};

const DescriptiveQuestionCard = ({
  srNumber = null,
  topic,
  type,
  isAttemped,
  handleAttemptClick,
  showSubmission,
  duration,
  attempts,
}: Props) => {
  return (
    <article className="rounded-md shadow-sm bg-white flex flex-col md:flex-row gap-4 p-3 px-4 md:p3 text-sm">
      <div className="flex-1">
        <p className="font-medium text-balance md:text-pretty text-black mb-2">
          #{srNumber} - {topic}
        </p>
        <div className="flex gap-2 flex-wrap">
          <Chip icon="clock" variant={"danger"}>
            {duration} min
          </Chip>
          <Chip icon="clock" variant={"success"}>
            {type}
          </Chip>
          <Chip icon="clock" variant={"warning"}>
            {attempts} Attempt
          </Chip>
        </div>
      </div>
      <div className="flex flex-col items-stretch gap-2 md:justify-center md:w-32">
        {attempts >= 1 && (
          <Button
            onClick={handleAttemptClick}
            size={"sm"}
            className="px-3 py-1"
            variant={"info"}
          >
            <Icon icon="play_circle" className="mr-2" /> Attempt
          </Button>
        )}
        {isAttemped && (
          <Button
            onClick={showSubmission}
            size={"sm"}
            className="px-3 py-1"
            variant={"secondary"}
          >
            <Icon icon="send" className="mr-2" /> View Submission
          </Button>
        )}
      </div>
    </article>
  );
};

export default DescriptiveQuestionCard;
