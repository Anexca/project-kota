import { yupResolver } from "@hookform/resolvers/yup";
import { useEffect, useMemo, useRef, useState } from "react";
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

import Icon from "../../componnets/base/icon";
import { useToast } from "../../hooks/use-toast";
import Chip from "../../componnets/base/chip";
import ProfanityError from "../../componnets/shared/profanity_error/profanity-error";
import { Evalution } from "../../interface/evaluation";
import { questionType } from "../../constants/shared";
import { StyledLink } from "../../componnets/base/styled-link";

const ConformationDialog = ({ timerStart, time, open, setOpen }: any) => {
  // const [open, setOpen] = useState(true);
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
const AssessmentResponseDialog = ({
  open,
  categoryId,
}: {
  open: boolean;
  categoryId: string;
}) => {
  return (
    <Dialog open={open}>
      <DialogContent noCloseButton>
        <DialogHeader>
          <DialogTitle className="text-center mb-2">
            Taking longer than expected to complete.
          </DialogTitle>
          <DialogDescription>
            <div className="mb-2 text-neutral-700 font-semibold text-center text-balance">
              Our AI is currently processing a high volume of requests. You can
              view the assessment on the 'My Submissions' page or by checking
              'View Assessments' after 2-4 minutes.
            </div>
            <div className="flex gap-2 justify-center ">
              <StyledLink
                variant={"info"}
                to={`/${paths.EXAMS}/banking/${paths.DISCRIPTIVE}/${categoryId}`}
                className="p-2 h-8"
              >
                Back To Questions
              </StyledLink>
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
const EXAM_TIMEOUT = 30;
const DiscriptiveExam = ({ isOpenMode }: { isOpenMode?: boolean }) => {
  const param = useParams();
  const navigate = useNavigate();
  const [startUpModal, setStartUpModal] = useState(true);
  const [exitModelOpen, setExitModelOpen] = useState(false);
  const [assessmentTimeoutModelOpen, setAssessmentTimeoutModelOpen] =
    useState(false);
  const [question, setQuestion] = useState<IQuestion | null>(null);
  const [examTime, setExamTime] = useState(0);
  const interval = useInterval(() => setExamTime((s) => s - 1), 1000);
  const fetchResultRef = useRef<any>(null);
  const fetcherTimeOut = useRef<NodeJS.Timeout>();
  const timeoutRef = useRef<number>(0);
  const [evaluationResult, setEvaluationResult] = useState<Evalution | null>(
    null
  );
  const [loading, setLoading] = useState(false);
  const { control, handleSubmit, formState, setValue } = useForm({
    defaultValues: { answer: "" },
    resolver: yupResolver(DiscriptiveExamSchema),
  });
  const { toast } = useToast();
  const answer = useWatch({ control, name: "answer" });
  const fetchQuestionById = async () => {
    if (!param?.questionId) return;
    try {
      const response = await getQuestionById(param?.questionId, isOpenMode);
      setValue("maxWords", response.data.raw_exam_data.max_number_of_words);
      setExamTime(response.data.duration_seconds);
      setQuestion(response.data);
    } catch (error) {
      toast({
        title: "Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in processing your request.",
      });
    }
  };
  const getResultByExamId = (examId: number) => {
    return async () => {
      const data = await getAssesmetsResult(examId);
      return data as Evalution;
      // return { data: { status: "PENDING" } };
    };
  };

  const handleFormSubmit = async (value: DiscriptiveExamSchemaType) => {
    if (!question) return;
    setLoading(true);
    interval.stop();
    const timeTaken = question!.duration_seconds - examTime;
    try {
      const response = await sendAnswerForAssesment({
        questionId: question!.exam_id,
        answer: value.answer,
        completedTime: timeTaken,
        isOpen: isOpenMode,
      });
      fetchResultRef.current = getResultByExamId(response.data.id);

      toast({
        title: "Successfully submitted the exam.",
        variant: "success",
        description: "Please wait while we evaluate your answer.",
      });
    } catch (error) {
      toast({
        title: "Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in processing your request.",
      });
      return;
    }
    timeoutRef.current = new Date().getTime();
    const recursiveFetcher = async () => {
      try {
        const r = await fetchResultRef.current();
        const data: Evalution = r.data;

        clearTimeout(fetcherTimeOut?.current);
        if (data.status === "PENDING") {
          if (
            Math.floor((new Date().getTime() - timeoutRef.current) / 1000) >=
            EXAM_TIMEOUT
          ) {
            clearTimeout(fetcherTimeOut?.current);
            exitExamAfterTimeout();
            return;
          }
          fetcherTimeOut.current = setTimeout(recursiveFetcher, 2000);
          return;
        }
        if (data.status === "COMPLETED") {
          setEvaluationResult(data);
        }
        if (data.status === "REJECTED") {
          setEvaluationResult(data);
        }
      } catch (error) {
        clearTimeout(fetcherTimeOut?.current);
        toast({
          title: "Something went wrong.",
          variant: "destructive",
          description:
            "Sorry there is some problem in processing your request.",
        });
      }
      setLoading(false);
    };
    recursiveFetcher();
  };

  const endExamNavigatePath = isOpenMode
    ? `/${paths.COMMUNITY_EXAMS}/banking/${paths.DISCRIPTIVE}`
    : `/${paths.EXAMS}/banking/${param.categoryId}`;
  const exitExam = () => {
    interval.stop();
    navigate(endExamNavigatePath);
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
  const exitExamAfterTimeout = () => {
    setAssessmentTimeoutModelOpen(true);
  };

  const noAttemptLeft = useMemo(() => {
    const hasAttemptLeft =
      (question?.max_attempts || 0) - (question?.user_attempts || 0) == 0 &&
      question;
    interval.stop();

    return hasAttemptLeft;
  }, [question]);
  return (
    <div className="flex flex-col pb-4">
      <ConformationDialog
        timerStart={() => interval.start()}
        time={question ? question?.duration_seconds / 60 : 0}
        open={startUpModal && !noAttemptLeft}
        setOpen={setStartUpModal}
      />
      {question && !examTime && <AutoSubmittDialog />}
      <EndExamDialog
        open={exitModelOpen}
        close={() => setExitModelOpen(false)}
        exitExam={exitExam}
      />
      <AssessmentResponseDialog
        open={assessmentTimeoutModelOpen}
        categoryId={param.categoryId!}
      />
      <TestHeader currentTime={examTime} active={interval.active} />
      <Container className="p-2">
        <div className="mb-4">
          {evaluationResult && (
            <Link
              to={endExamNavigatePath}
              className="text-info pb-2 inline-block"
            >
              <Icon icon="arrow_back" className="text-sm mr-1" /> Exit Exam
            </Link>
          )}
          <div className="text-sm font-medium">
            Que - {question?.raw_exam_data.topic}
          </div>
          <div>
            <Chip icon="clock" variant={"success"}>
              {question?.raw_exam_data.type
                ? questionType[question?.raw_exam_data.type] || "--"
                : "--"}
            </Chip>
            <Chip icon="target" className="ml-2">
              Marks {question?.raw_exam_data.total_marks}
            </Chip>
          </div>
          {!evaluationResult &&
            question?.raw_exam_data?.hints &&
            !!question?.raw_exam_data?.hints.length && (
              <Hints hints={question?.raw_exam_data?.hints || []} />
            )}
        </div>
        {!evaluationResult && !noAttemptLeft && (
          <>
            <form onSubmit={handleSubmit(handleFormSubmit)}>
              <div className="w-full mb-4 border border-gray-200 rounded-lg bg-gray-50 dark:bg-gray-700 dark:border-gray-600">
                <div className="px-2 py-2 bg-white rounded-t-lg dark:bg-gray-800">
                  <Controller
                    name="answer"
                    control={control}
                    render={({ field }) => {
                      return (
                        <Textarea
                          {...field}
                          className="text-sm min-h-[50vh]"
                          rows={10}
                          spellCheck={"false"}
                          autoComplete="off"
                          autoCorrect="off"
                          autoCapitalize="off"
                        />
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
                      {answer?.match(/\b\w+(?:[.,!;?])?\b/g)?.length || 0}/
                      {question?.raw_exam_data.max_number_of_words}
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
        {noAttemptLeft && (
          <div className="flex p-4 rounded bg-white shadow-sm mt-2 gap-4">
            <div className=" w-2 bg-info rounded-full"></div>
            <div className="flex flex-col flex-1">
              <div className="text-sm font-semibold text-info">
                Max Attempts Exceeded
              </div>
              <div className="text-sm font-medium">
                You've reached the maximum number of attempts for this question.
              </div>
              <div className="text-sm font-medium pt-2">
                <StyledLink
                  size={"sm"}
                  className="h-7 px-3"
                  to={`/${paths.EXAMS}/banking/${paths.DISCRIPTIVE}`}
                >
                  Back To Questions
                </StyledLink>
              </div>
            </div>
          </div>
        )}
        {evaluationResult?.status === "COMPLETED" ? (
          <DiffChecker
            rating={evaluationResult.raw_assesment_data.rating}
            weaknesses={evaluationResult.raw_assesment_data.weaknesses}
            strength={evaluationResult.raw_assesment_data.strengths}
            modifiedText={evaluationResult.raw_assesment_data.corrected_version}
            originalText={evaluationResult.raw_user_submission.content}
          />
        ) : null}
        {evaluationResult?.status === "REJECTED" ? (
          <ProfanityError data={evaluationResult} />
        ) : null}
      </Container>
    </div>
  );
};

export default DiscriptiveExam;
