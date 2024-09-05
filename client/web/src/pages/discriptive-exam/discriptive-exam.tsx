import { useEffect, useRef, useState } from "react";
import { Button } from "../../componnets/base/button/button";
import { Textarea } from "../../componnets/base/text-area";
import { Controller, useForm, useWatch } from "react-hook-form";
import Container from "../../componnets/shared/container";
import { yupResolver } from "@hookform/resolvers/yup";
import {
  DiscriptiveExamSchema,
  DiscriptiveExamSchemaType,
} from "../../validation-schema/discriptive-exam";
import { IQuestion } from "../../interface/question";
import { useInterval } from "../../hooks/use-interval";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "../../componnets/base/dialog/dialog";
import TestHeader from "../../componnets/shared/test-header/test-header";
import DiffChecker from "../../componnets/shared/diffchecker";
import {
  getAssesmetsResult,
  getQuestionById,
  sendAnswerForAssesment,
} from "../../services/exam.service";
import { useParams } from "react-router-dom";

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
const DiscriptiveExam = () => {
  const param = useParams();
  const [question, setQuestion] = useState<IQuestion | null>(null);
  const [examTime, setExamTime] = useState(0);
  const interval = useInterval(() => setExamTime((s) => s - 1), 1000);
  const fetchResultRef = useRef<any>(null);
  const fetcherTimeOut = useRef<NodeJS.Timeout>();
  const [evaluationResult, setEvaluationResult] = useState(null);
  const [loading, setLoading] = useState(false);
  const { control, handleSubmit, formState } = useForm({
    defaultValues: { answer: "" },
    resolver: yupResolver(DiscriptiveExamSchema),
  });

  const textLength = useWatch({ control, name: "answer" });
  const fetchQuestionById = async () => {
    if (!param?.questionId) return;
    const response = await getQuestionById(param?.questionId);
    console.log(response);
    // setExamTime(response.data.duration_minutes);
    setExamTime(response.data.duration_seconds);
    setQuestion(response.data);
  };
  const getResultByExamId = (examId: number) => {
    return async () => {
      const data = await getAssesmetsResult(examId);
      return data;
    };
  };

  const handleFormSubmit = async (value: DiscriptiveExamSchemaType) => {
    if (!question) return;
    setLoading(true);
    interval.stop();
    const timeTaken = question!.duration_seconds - examTime;

    const response = await sendAnswerForAssesment({
      questionId: question!.id,
      answer: value.answer,
      completedTime: timeTaken,
    });
    fetchResultRef.current = getResultByExamId(response.data.id);
    const recursiveFetcher = async () => {
      const r = await fetchResultRef.current();
      console.log(r);
      clearTimeout(fetcherTimeOut?.current);
      if (!r) {
        fetcherTimeOut.current = setTimeout(recursiveFetcher, 2000);
        return;
      }
      setLoading(false);
      setEvaluationResult(r);
    };
    recursiveFetcher();
  };

  useEffect(() => {
    if (examTime <= 0 && interval.active) {
      interval.stop();
      console.log("stop");
    }
  }, [examTime]);

  useEffect(() => {
    fetchQuestionById();
  }, []);

  return (
    <div className="flex flex-col">
      <ConformationDialog
        timerStart={() => interval.start()}
        time={question ? question?.duration_seconds / 60 : 0}
      />
      <TestHeader currentTime={examTime} active={interval.active} />
      <Container className="p-2">
        <div className="mb-4">
          <div className="text-sm font-medium">
            Que - {question?.raw_exam_data.topic}
          </div>
          <span className="inline-flex items-center justify-center rounded-full bg-emerald-100 px-2.5 py-0.5 text-emerald-700">
            <i className="fa-regular fa-circle-check text-sm mr-2"></i>
            <p className="whitespace-nowrap text-sm capitalize">
              {question?.raw_exam_data.type}
            </p>
          </span>
          <Hints hints={question?.raw_exam_data?.hints || []} />
        </div>
        <form
          onSubmit={handleSubmit(handleFormSubmit, (e, _i) => console.log(e))}
        >
          <div className="w-full mb-4 border border-gray-200 rounded-lg bg-gray-50 dark:bg-gray-700 dark:border-gray-600">
            <div className="px-2 py-2 bg-white rounded-t-lg dark:bg-gray-800">
              <Controller
                name="answer"
                control={control}
                render={({ field }) => {
                  return <Textarea {...field} className="text-sm" rows={10} />;
                }}
              />
            </div>
            {formState?.errors?.answer && (
              <div className="text-destructive text-sm px-4 font-semibold">
                {formState?.errors?.answer?.message}
              </div>
            )}
            <div className="flex items-center justify-between px-3 py-2 border-t dark:border-gray-600">
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
              <div className="text-sm">
                {textLength?.match(/\b\w+(?:[.,!;?])?\b/g)?.length || 0}/250
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

        {evaluationResult ? (
          <DiffChecker modifiedText="" originalText="" />
        ) : null}
      </Container>
    </div>
  );
};

export default DiscriptiveExam;
