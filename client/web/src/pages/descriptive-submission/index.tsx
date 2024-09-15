import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import Chip from "../../componnets/base/chip";
import Icon from "../../componnets/base/icon";
import Container from "../../componnets/shared/container";
import DiffChecker from "../../componnets/shared/diffchecker";
import { useToast } from "../../hooks/use-toast";
import { Evalution } from "../../interface/evaluation";
import { IQuestion } from "../../interface/question";
import { paths } from "../../routes/route.constant";
import {
  getAssesmetsResult,
  getQuestionById,
} from "../../services/exam.service";
import ProfanityError from "../../componnets/shared/profanity_error/profanity-error";
import { questionType } from "../../constants/shared";

const DescriptiveSubmission = ({ backLink }: { backLink?: string }) => {
  const [question, setQuestion] = useState<IQuestion | null>(null);

  const [assessment, setAssessment] = useState<Evalution | null>(null);
  const [loading, setLoading] = useState(false);

  const params = useParams();
  const { toast } = useToast();
  const fetchQuestionById = async () => {
    if (!params?.questionId) return;
    const response = await getQuestionById(params?.questionId);
    setQuestion(response.data);
  };
  const getResultByExamId = async () => {
    setLoading(true);
    try {
      const data = await getAssesmetsResult(Number(params.assesmentId!));
      setAssessment(data.data);
    } catch (error) {
      toast({
        title: "Oh ho Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in proccessing your request.",
      });
    }
    setLoading(false);
  };
  useEffect(() => {
    fetchQuestionById();
    getResultByExamId();
  }, []);
  return (
    <Container className="p-2 mb-4">
      <div>
        <div className="text-sm font-medium">
          <Link
            className="text-info mr-2 text-sm"
            to={backLink || `/${paths.EXAMS}/banking/${paths.DISCRIPTIVE}`}
          >
            <Icon icon="arrow_back" /> Back
          </Link>{" "}
          {question?.raw_exam_data.topic}
        </div>

        <Chip variant={"success"} className="capitalize" icon="file">
          {question?.raw_exam_data.type
            ? questionType[question?.raw_exam_data.type] || "--"
            : "--"}
        </Chip>
      </div>
      {loading ? (
        <div className="flex flex-col gap-2 justify-center items-center">
          <span className="rounded-full w-8 h-8 animate-spin border-2 border-info border-t-info/30"></span>
          Getting Your Submission
        </div>
      ) : (
        <>
          {assessment?.status === "COMPLETED" ? (
            <DiffChecker
              modifiedText={assessment?.raw_assesment_data.corrected_version}
              originalText={assessment?.raw_user_submission.content}
              rating={assessment?.raw_assesment_data.rating}
              weaknesses={assessment?.raw_assesment_data.weaknesses}
              strength={assessment?.raw_assesment_data.strengths}
            />
          ) : null}
          {assessment?.status === "REJECTED" ? (
            <ProfanityError data={assessment} />
          ) : null}
        </>
      )}
    </Container>
  );
};

export default DescriptiveSubmission;
