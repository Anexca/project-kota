import { Button } from "../../base/button/button";

type Props = {
  srNumber?: number | null;
  topic: string;
  type: string;
  isAttemped?: boolean;
  handleAttemptClick?: () => void;
  duration?: number;
  attempts: number;
  isNew: number;
};

const DescriptiveQuestionCard = ({
  srNumber = null,
  topic,
  type,
  isAttemped,
  handleAttemptClick,
  duration,
  attempts,
  isNew,
}: Props) => {
  return (
    <article className="rounded-xl border-2 border-gray-100 bg-white">
      <div className="flex items-start gap-2 !pb-1 p-2 sm:p-4 lg:p-6">
        {srNumber !== null && <div>{srNumber}.</div>}

        <div className="space-y-2">
          <h3 className="font-medium text-sm sm:text-md">{topic}</h3>

          <div className="flex items-center gap-2 flex-wrap">
            <span className="inline-flex items-center justify-center rounded-full bg-orange-100 px-2.5 py-0.5 text-orange-700">
              <i className="fa-regular fa-clock text-sm mr-2"></i>
              <p className="whitespace-nowrap text-sm">{duration} min</p>
            </span>
            <span className="inline-flex items-center justify-center rounded-full bg-emerald-100 px-2.5 py-0.5 text-emerald-700">
              <i className="fa-regular fa-circle-check text-sm mr-2"></i>
              <p className="whitespace-nowrap text-sm capitalize">{type}</p>
            </span>
            <span className="inline-flex items-center justify-center rounded-full bg-yellow-100 px-2.5 py-0.5 text-yellow-700">
              <i className="fa-solid fa-rotate-right text-sm mr-2"></i>
              <p className="whitespace-nowrap text-sm capitalize">
                {attempts} attempt left
              </p>
            </span>
          </div>
        </div>
      </div>
      <div className="flex justify-end">
        {!!isNew && (
          <Button
            variant={"secondary"}
            className="px-3 py-1 mr-auto  h-auto text-sm rounded-none rounded-es-xl rounded-se-xl"
          >
            <i className="fa-regular fa-paper-plane text-sm mr-2"></i>
            Attempted
          </Button>
        )}
        {isAttemped ? (
          <strong className="-mb-[2px] -me-[2px] inline-flex items-center gap-1  rounded-ee-xl rounded-ss-xl bg-green-600 px-3 py-1.5 text-white">
            <span className="text-[10px] font-medium sm:text-xs">
              🎉 Solved!
            </span>
          </strong>
        ) : (
          <>
            {handleAttemptClick ? (
              <Button
                onClick={handleAttemptClick}
                variant={"info"}
                className="px-2 py-1 rounded-none h-auto rounded-ee-xl rounded-ss-xl"
              >
                <i className="fa-regular fa-play-circle text-sm mr-2"></i>
                Take Test
              </Button>
            ) : null}
          </>
        )}
      </div>
    </article>
  );
};

export default DescriptiveQuestionCard;
