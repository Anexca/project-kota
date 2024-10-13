import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useToast } from "../../../hooks/use-toast";
import { IMCQExam } from "../../../interface/question";
import { paths } from "../../../routes/route.constant";
import { getPastSubmission } from "../../../services/exam.service";
import { Button } from "../../base/button/button";
import Chip from "../../base/chip";
import Icon from "../../base/icon";
interface ISubmission {
  id: number;
  completed_seconds: number;
  status: "COMPLETED" | "PENDING" | "REJECTED"; // Assuming possible statuses
  created_at: string; // Could also be Date type depending on how you handle dates
  updated_at: string; // Could also be Date type depending on how you handle dates
}
const iconSet = {
  COMPLETED: (
    <Icon
      icon="check_solid"
      className="text-green-700 text-xl mr-4 self-start md:self-auto"
    />
  ),
  PENDING: (
    <Icon
      icon="exclaimation_circle"
      className="text-yellow-700 text-xl mr-4 self-start md:self-auto"
    />
  ),
  REJECTED: (
    <Icon
      icon="xmark_circle"
      className="text-destructive text-xl mr-4 self-start md:self-auto"
    />
  ),
};
const MCQPreviousSubmissions = ({
  question,
  isOpenExam,
}: {
  question: IMCQExam;
  isOpenExam?: boolean;
}) => {
  const [submissionList, setSubmissionList] = useState<ISubmission[]>([]);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const { toast } = useToast();
  const fetchPastSubmissions = async () => {
    setLoading(true);
    try {
      const res = await getPastSubmission(question.exam_id);
      setSubmissionList(res.data);
    } catch (error) {
      toast({
        title: "Oh ho Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in proccessing your request.",
      });
    }
    setLoading(false);
  };
  const viewSubmission = (id: number) => {
    const path = isOpenExam
      ? `/${paths.COMMUNITY_EXAMS}/banking/${paths.MCQ}/${question.exam_id}/${paths.SUBMISSION}/${id}`
      : `/${paths.EXAMS}/${paths.MY_SUMBISSIONS}/${paths.MCQ}/${question.exam_id}/${paths.SUBMISSION}/${id}`;
    navigate(path);
  };
  useEffect(() => {
    fetchPastSubmissions();
  }, []);

  return (
    <div className="flex flex-col flex-1 overflow-hidden">
      <div className=" flex flex-col gap-2 p-1 items-start">
        <div className="text-sm font-medium capitalize">
          Que - {question?.exam_name.toLowerCase()}{" "}
          {question?.exam_stage.toLowerCase()}
        </div>
        <div className="text-sm font-medium capitalize">
          Question Paper Set - #{question.exam_id}
        </div>

        <Chip variant={"success"} icon={"file"}>
          {question?.exam_type || "--"}
        </Chip>
      </div>
      {loading ? (
        <div className="flex flex-col gap-2 justify-center items-center">
          <span className="rounded-full w-8 h-8 animate-spin border-2 border-info border-t-info/30"></span>
          Getting Your Submission
        </div>
      ) : (
        <div className="animate-fadeIn flex-1 overflow-hidden">
          <div className="overflow-y-auto flex flex-col p-1 gap-2 max-h-full">
            {submissionList.map((item, index) => {
              const seconds = Math.floor(item.completed_seconds % 60);
              const minutes = Math.floor(item.completed_seconds / 60);
              return (
                <article className="rounded-md shadow bg-white flex flex-col gap-4 p-3 px-4 md:p3 text-sm">
                  <div className="flex flex-1 items-center text-sm text-black">
                    <div>
                      <div className="mb-2 font-medium flex items-center">
                        {iconSet[item.status]}
                        Submission: {index + 1}
                      </div>
                      <div className="flex gap-2 flex-wrap">
                        <Chip icon="clock" variant={"info"}>
                          {minutes}:{seconds} min
                        </Chip>
                        <Chip icon="calender_solid" variant={"info"}>
                          {new Date(item.created_at).toLocaleString()}
                        </Chip>
                      </div>
                    </div>
                  </div>

                  <Button
                    size="sm"
                    className="text-sm py-1"
                    variant={"secondary"}
                    onClick={() => viewSubmission(item.id)}
                  >
                    <Icon icon="send" className="mr-2" /> View Assessent
                  </Button>
                </article>
              );
            })}
          </div>
        </div>
      )}
    </div>
  );
};

export default MCQPreviousSubmissions;
