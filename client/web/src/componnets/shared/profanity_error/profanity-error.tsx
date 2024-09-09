import { EvaluationRejected } from "../../../interface/evaluation";

type Props = {
  data: EvaluationRejected;
};

const ProfanityError = ({ data }: Props) => {
  const isProfanity = data.raw_assesment_data?.profanity_check === "detected";
  return (
    <div className="flex p-4 rounded bg-white shadow-sm mt-2 gap-4">
      <div className=" w-2 bg-destructive rounded-full"></div>
      <div className="flex flex-col flex-1">
        <div className="text-sm font-semibold text-destructive">
          {isProfanity ? "Profanity Check Failed" : "Server Error"}
        </div>
        <div className="text-sm font-medium">
          {isProfanity
            ? "Your submission couldn't be completed due to inappropriate words in your answer."
            : "Your request couldn't be completed due to a server issue. Please contact the administrator."}
        </div>
      </div>
    </div>
  );
};

export default ProfanityError;
