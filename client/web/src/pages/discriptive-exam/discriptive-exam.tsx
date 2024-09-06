import { yupResolver } from "@hookform/resolvers/yup";
import { useEffect, useRef, useState } from "react";
import { Controller, useForm, useWatch } from "react-hook-form";
import { Link, useNavigate, useParams } from "react-router-dom";
import { Button } from "../../componnets/base/button/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "../../componnets/base/dialog/dialog";
import { Textarea } from "../../componnets/base/text-area";
import Container from "../../componnets/shared/container";
import DiffChecker from "../../componnets/shared/diffchecker";
import TestHeader from "../../componnets/shared/test-header/test-header";
import { useInterval } from "../../hooks/use-interval";
import { IQuestion } from "../../interface/question";
import { paths } from "../../routes/route.constant";
import {
  getAssesmetsResult,
  getQuestionById,
  sendAnswerForAssesment,
} from "../../services/exam.service";
import {
  DiscriptiveExamSchema,
  DiscriptiveExamSchemaType,
} from "../../validation-schema/discriptive-exam";

const ConformationDialog = ({ timerStart, time }: any) => {
  const [open, setOpen] = useState(true);
  return (
    <Dialog open={open} defaultOpen>
      <DialogContent noCloseButton>
        <DialogHeader>
          <DialogTitle>Exam Instruction.</DialogTitle>
          <DialogDescription>
            <div className="mb-2 text-neutral-700 font-semibold">
              The exam will automatically close when the time is up, so please
              manage your time carefully. You have{" "}
              <span className="text-orange-600">{time}</span> minutes to
              complete the exam.
            </div>
            <Button
              variant={"info"}
              onClick={() => {
                timerStart();
                setOpen(false);
              }}
              className="p-2 h-8"
            >
              Start Exam
            </Button>
          </DialogDescription>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
};
const AutoSubmittDialog = () => {
  const [open, setOpen] = useState(true);
  return (
    <Dialog open={open} defaultOpen>
      <DialogContent noCloseButton>
        <DialogHeader>
          <DialogTitle>Exam End.</DialogTitle>
          <DialogDescription>
            <div className="mb-2 text-neutral-700 font-semibold">
              Time is up. your answer is automatically submitted.
            </div>
            <Button
              variant={"info"}
              onClick={() => {
                setOpen(false);
              }}
              className="p-2 h-8"
            >
              Ok
            </Button>
          </DialogDescription>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
};
const EndExamDialog = ({
  open,
  close,
  exitExam,
}: {
  open: boolean;
  close: () => void;
  exitExam: () => void;
}) => {
  return (
    <Dialog open={open}>
      <DialogContent noCloseButton>
        <DialogHeader>
          <DialogTitle>Exam End.</DialogTitle>
          <DialogDescription>
            <div className="mb-2 text-neutral-700 font-semibold ">
              Are you sure.
            </div>
            <div className="flex gap-2 justify-center sm:justify-start">
              <Button
                variant={"warning"}
                onClick={() => {
                  exitExam();
                  close();
                }}
                className="p-2 h-8"
              >
                Exit
              </Button>
              <Button variant={"secondary"} onClick={close} className="p-2 h-8">
                Cancel
              </Button>
            </div>
          </DialogDescription>
        </DialogHeader>
        <DialogFooter></DialogFooter>
      </DialogContent>
    </Dialog>
  );
};
const Hints = ({ hints }: { hints: string[] }) => {
  return (
    <details className="mt-2 group overflow-hidden rounded [&_summary::-webkit-details-marker]:hidden">
      <summary className="flex cursor-pointer items-center justify-between gap-2 bg-white px-2 text-info transition">
        <span className="text-sm font-medium">Hints</span>

        <span className="transition group-open:-rotate-180">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            strokeWidth="1.5"
            stroke="currentColor"
            className="size-4"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M19.5 8.25l-7.5 7.5-7.5-7.5"
            />
          </svg>
        </span>
      </summary>

      <div className="border border-gray-200 rounded-lg bg-white p-4">
        <ul className="space-y-1 border-gray-200 ml-4">
          {hints.map((hint) => (
            <li className=" text-sm list-disc">{hint}</li>
          ))}
        </ul>
      </div>
    </details>
  );
};

type EvaluationPending = {
  id: number;
  completed_seconds: number;
  status: "PENDING";
  created_at: string;
  updated_at: string;
};
type EvaluationCompleted = {
  id: number;
  completed_seconds: number;
  raw_assesment_data: {
    corrected_version: string;
    rating: string;
    strengths: string[];
    weakness: string[];
  };
  raw_user_submission: {
    content: string;
  };
  status: "COMPLETED";
  created_at: string;
  updated_at: string;
};

type Evalution = EvaluationPending | EvaluationCompleted;

const DiscriptiveExam = () => {
  const param = useParams();
  const navigate = useNavigate();
  const [exitModelOpen, setExitModelOpen] = useState(false);
  const [question, setQuestion] = useState<IQuestion | null>(null);
  const [examTime, setExamTime] = useState(0);
  const interval = useInterval(() => setExamTime((s) => s - 1), 1000);
  const fetchResultRef = useRef<any>(null);
  const fetcherTimeOut = useRef<NodeJS.Timeout>();
  const [evaluationResult, setEvaluationResult] =
    useState<EvaluationCompleted | null>(null);
  const [loading, setLoading] = useState(false);
  const { control, handleSubmit, formState } = useForm({
    defaultValues: { answer: "" },
    resolver: yupResolver(DiscriptiveExamSchema),
  });

  const answer = useWatch({ control, name: "answer" });
  const fetchQuestionById = async () => {
    if (!param?.questionId) return;
    const response = await getQuestionById(param?.questionId);

    setExamTime(response.data.duration_seconds);
    setQuestion(response.data);
  };
  const getResultByExamId = (examId: number) => {
    return async () => {
      const data = await getAssesmetsResult(examId);
      return data as Evalution;
    };
  };

  const handleFormSubmit = async (value: DiscriptiveExamSchemaType) => {
    if (!question) return;
    setLoading(true);
    interval.stop();
    const timeTaken = question!.duration_seconds - examTime;
    try {
      const response = await sendAnswerForAssesment({
        questionId: question!.id,
        answer: value.answer,
        completedTime: timeTaken,
      });
      fetchResultRef.current = getResultByExamId(response.data.id);
    } catch (error) {
      return;
    }

    const recursiveFetcher = async () => {
      try {
        const r = await fetchResultRef.current();
        const data: Evalution = r.data;

        clearTimeout(fetcherTimeOut?.current);
        if (data.status === "PENDING") {
          fetcherTimeOut.current = setTimeout(recursiveFetcher, 2000);
          return;
        }
        if (data.status === "COMPLETED") {
          setEvaluationResult(data);
        }
      } catch (error) {
        clearTimeout(fetcherTimeOut?.current);
      }
      setLoading(false);
    };
    recursiveFetcher();
  };

  const exitExam = () => {
    interval.stop();
    navigate(`/${paths.QUESTION_PAPER}`);
  };

  useEffect(() => {
    if (examTime <= 0 && interval.active) {
      interval.stop();
      handleFormSubmit({ answer });
    }
  }, [examTime]);

  useEffect(() => {
    fetchQuestionById();
  }, []);

  return (
    <div className="flex flex-col pb-4">
      <ConformationDialog
        timerStart={() => interval.start()}
        time={question ? question?.duration_seconds / 60 : 0}
      />
      {question && !examTime && <AutoSubmittDialog />}
      <EndExamDialog
        open={exitModelOpen}
        close={() => setExitModelOpen(false)}
        exitExam={exitExam}
      />
      <TestHeader currentTime={examTime} active={interval.active} />
      <Container className="p-2">
        <div className="mb-4">
          {evaluationResult && (
            <Link
              to={`/${paths.QUESTION_PAPER}`}
              className="text-info pb-2 inline-block"
            >
              <i className="fa-solid fa-arrow-left text-sm mr-1"></i> Exit Exam
            </Link>
          )}
          <div className="text-sm font-medium">
            Que - {question?.raw_exam_data.topic}
          </div>
          <span className="inline-flex items-center justify-center rounded-full bg-emerald-100 px-2.5 py-0.5 text-emerald-700">
            <i className="fa-regular fa-circle-check text-sm mr-2"></i>
            <p className="whitespace-nowrap text-sm capitalize">
              {question?.raw_exam_data.type}
            </p>
          </span>
          {!evaluationResult && question?.raw_exam_data?.hints && (
            <Hints hints={question?.raw_exam_data?.hints || []} />
          )}
        </div>
        {!evaluationResult && (
          <>
            <form
              onSubmit={handleSubmit(handleFormSubmit, (e, _i) =>
                console.log(e)
              )}
            >
              <div className="w-full mb-4 border border-gray-200 rounded-lg bg-gray-50 dark:bg-gray-700 dark:border-gray-600">
                <div className="px-2 py-2 bg-white rounded-t-lg dark:bg-gray-800">
                  <Controller
                    name="answer"
                    control={control}
                    render={({ field }) => {
                      return (
                        <Textarea {...field} className="text-sm" rows={10} />
                      );
                    }}
                  />
                </div>
                {formState?.errors?.answer && (
                  <div className="text-destructive text-sm px-4 font-semibold">
                    {formState?.errors?.answer?.message}
                  </div>
                )}
                <div className="flex flex-col px-3 py-2 border-t dark:border-gray-600">
                  <div className="flex justify-end mb-2">
                    {" "}
                    <div className="text-xs">
                      {answer?.match(/\b\w+(?:[.,!;?])?\b/g)?.length || 0}/250
                      words
                    </div>
                  </div>
                  <div className="flex justify-between">
                    <Button
                      disabled={loading || !!evaluationResult}
                      variant={"info"}
                      type="submit"
                      className="p-2 h-8 transition-all"
                    >
                      {loading ? (
                        <span className="rounded-full w-4 h-4 animate-spin border border-t-gray-100/30"></span>
                      ) : (
                        "Submit Answer"
                      )}
                    </Button>
                    <Button
                      type="button"
                      disabled={loading || !!evaluationResult}
                      variant={"warning"}
                      onClick={() => setExitModelOpen(true)}
                      className="p-2 h-8 transition-all"
                    >
                      Exit Test
                    </Button>
                  </div>
                </div>
              </div>
            </form>

            <p className="ms-auto text-xs font-semibold text-gray-500 dark:text-gray-400">
              Remember, you have only{" "}
              <span className="text-fuchsia-700">
                {(question?.max_attempts || 0) - (question?.user_attempts || 0)}
              </span>{" "}
              attempts left for this exam.
            </p>
          </>
        )}

        {evaluationResult ? (
          <DiffChecker
            rating={evaluationResult.raw_assesment_data.rating}
            weaknesses={evaluationResult.raw_assesment_data.weakness}
            strength={evaluationResult.raw_assesment_data.strengths}
            modifiedText={evaluationResult.raw_assesment_data.corrected_version}
            originalText={evaluationResult.raw_user_submission.content}
          />
        ) : null}
      </Container>
    </div>
  );
};

export default DiscriptiveExam;
