import { MathJax, MathJaxContext } from "better-react-mathjax";
import { useEffect, useMemo, useState } from "react";
import { Controller, useForm, useWatch } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import { Button } from "../../componnets/base/button/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "../../componnets/base/dialog/dialog";
import AnswerOptions from "../../componnets/shared/answer-option/answer-options";
import { ReadMore } from "../../componnets/shared/read-more-content";
import TestHeader from "../../componnets/shared/test-header/test-header";
import { QUESTION_STATE, ScreenSizeQuery } from "../../constants/shared";
import { useInterval } from "../../hooks/use-interval";
import { useMediaQuery } from "../../hooks/use-media-query";
import { useToast } from "../../hooks/use-toast";
import { MCQFormModal } from "../../interface/mcq-exam";
import { IContentGroup, IMCQQuestionSet } from "../../interface/question";
import { getMCQQuestionById } from "../../services/exam.service";
import m from "./m.module.scss";
const ConformationDialog = ({ timerStart, time, open, setOpen }: any) => {
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

const MCQExamCenter = ({ isOpenMode }: { isOpenMode?: boolean }) => {
  const param = useParams();
  const [startUpModal, setStartUpModal] = useState(true);
  const [exitModelOpen, setExitModelOpen] = useState(false);
  const [questionSet, setQuestionSet] = useState<IMCQQuestionSet | null>(null);
  const [examTime, setExamTime] = useState(0);
  const interval = useInterval(() => setExamTime((s) => s - 1), 1000);
  const { toast } = useToast();
  const isScreenView = useMediaQuery(ScreenSizeQuery.largeScreen, true);
  const [assessmentTimeoutModelOpen, setAssessmentTimeoutModelOpen] =
    useState(false);
  const navigate = useNavigate();

  const { control, setValue, reset } = useForm<
    MCQFormModal & { tempAnswerIndex: number | null }
  >({
    defaultValues: {
      answers: [],
      activeQuestionIndex: 0,
      tempAnswerIndex: null,
    },
  });
  const activeIndex = useWatch({ control, name: "activeQuestionIndex" });
  const answers = useWatch({ control, name: "answers" });
  const tempAnswerIndex = useWatch({ control, name: "tempAnswerIndex" });

  const fetchExam = async () => {
    try {
      const res = await getMCQQuestionById(param.questionId!);
      setExamTime(res.data.duration_seconds);
      setQuestionSet(res.data);
      const ques = res.data.raw_exam_data.questions.map((_) => ({
        state: QUESTION_STATE.UN_ATTEMPTED,
        selectedOption: null,
      }));
      reset({ activeQuestionIndex: 0, answers: ques });
    } catch (error) {
      toast({
        title: "Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in processing your request.",
      });
    }
  };

  useEffect(() => {
    fetchExam();
  }, []);

  const contentGroup = useMemo(() => {
    const temp: Record<string, IContentGroup> = {};
    if (questionSet) {
      questionSet?.raw_exam_data?.content_groups?.forEach((e) => {
        temp[e.content_id] = e;
      });
    }
    return temp;
  }, [questionSet]);
  const handleQuestionChange = (index: number) => {
    if (index == activeIndex) return;
    if (answers?.[activeIndex]?.state == "UN-ATTEMPTED") {
      setValue(`answers.${activeIndex}.state`, "NOT-ANSWERED");
    }
    setValue("activeQuestionIndex", index);
  };

  const handleSaveNNext = () => {
    if (tempAnswerIndex) {
      setValue(`answers.${activeIndex}.state`, "ATTEMPTED");
      setValue(`answers.${activeIndex}.selectedOption`, tempAnswerIndex);
      setValue("tempAnswerIndex", null);
    }
  };
  const handleReviewNNext = () => {
    if (tempAnswerIndex) {
      setValue(`answers.${activeIndex}.state`, "FOR-REVIEW");
      setValue(`answers.${activeIndex}.selectedOption`, tempAnswerIndex);
      setValue("tempAnswerIndex", null);
    }
  };

  const questions = questionSet?.raw_exam_data?.questions || [];
  const activeContentReference = questions?.[activeIndex]?.content_reference_id;
  const contentInfo = contentGroup[activeContentReference];

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
      // handleFormSubmit({ answer });
      exitExamAfterTimeout();
    }
  }, [examTime]);
  const exitExamAfterTimeout = () => {
    setAssessmentTimeoutModelOpen(true);
  };
  const noAttemptLeft = useMemo(() => {
    const hasAttemptLeft =
      (questionSet?.max_attempts || 0) - (questionSet?.user_attempts || 0) ==
        0 && questionSet;
    interval.stop();

    return hasAttemptLeft;
  }, [questionSet]);
  return (
    <MathJaxContext>
      <div className="md:max-h-screen md:h-screen flex flex-col ">
        <TestHeader currentTime={examTime} active={interval.active} />
        <ConformationDialog
          timerStart={() => interval.start()}
          time={questionSet ? questionSet?.duration_seconds / 60 : 0}
          open={startUpModal && !noAttemptLeft}
          setOpen={setStartUpModal}
        />
        {questionSet && !examTime && <AutoSubmittDialog />}
        <EndExamDialog
          open={exitModelOpen}
          close={() => setExitModelOpen(false)}
          exitExam={exitExam}
        />
        <AssessmentResponseDialog
          open={assessmentTimeoutModelOpen}
          categoryId={param.categoryId!}
        />
        <div className="flex-1 md:overflow-hidden">
          <div className="flex flex-col md:flex-row items-stretch md:max-h-full md:h-full">
            <div className="flex-1 max-h-full flex flex-col">
              <div className="text-start p-2 pl-4 shadow">
                Que No {activeIndex + 1} /{" "}
                {questionSet?.number_of_questions || 0}
              </div>
              <div className="flex-1 flex flex-col md:flex-row md:overflow-hidden">
                {contentInfo && (
                  <div className="md:w-1/2 min-w-[50%] overflow-auto p-4 pt-0 md:pt-4  text-pretty font-medium border-r-0 border-b md:border-b-0 md:border-r">
                    {contentInfo?.instructions && (
                      <p className="font-semibold mb-2 pt-4 md:pt-0">
                        {contentInfo.instructions}
                      </p>
                    )}
                    {contentInfo?.content && (
                      <>
                        {isScreenView ? (
                          <p>{contentInfo.content}</p>
                        ) : (
                          <ReadMore text={contentInfo.content} />
                        )}
                      </>
                    )}
                  </div>
                )}
                <div className="md:max-h-[70vh] h-full flex flex-col p-4 flex-1">
                  {questions.length && (
                    <>
                      <p className="items-start text-start pb-2">
                        <MarkdownRender
                          children={questions?.[activeIndex]?.question || ""}
                        />
                      </p>

                      <Controller
                        name={"tempAnswerIndex"}
                        control={control}
                        render={({ field }) => (
                          <AnswerOptions
                            name={field.name}
                            onChange={field.onChange}
                            options={questions[activeIndex].options}
                            selected={field.value}
                          />
                        )}
                      />
                    </>
                  )}
                </div>
              </div>
              <div className="min-h-12 flex gap-2 p-2 border-orange-300 border-t">
                <Controller
                  name={`activeQuestionIndex`}
                  control={control}
                  render={({ field }) => (
                    <Button
                      size={"sm"}
                      type="button"
                      disabled={
                        field.value + 1 >= questions.length ||
                        tempAnswerIndex == null
                      }
                      onClick={() => {
                        const nextIdx = field.value + 1;
                        handleReviewNNext();
                        field.onChange(nextIdx);
                      }}
                    >
                      Mark for Review & Next
                    </Button>
                  )}
                />
                <Button variant={"outline"} size={"sm"} className="mr-auto">
                  Clear
                </Button>
                <Controller
                  name={`activeQuestionIndex`}
                  control={control}
                  render={({ field }) => (
                    <Button
                      size={"sm"}
                      variant={"success"}
                      className="justify-self-end"
                      type="button"
                      disabled={
                        field.value + 1 >= questions.length ||
                        tempAnswerIndex == null
                      }
                      onClick={() => {
                        const nextIdx = field.value + 1;
                        handleSaveNNext();
                        field.onChange(nextIdx);
                      }}
                    >
                      SAVE & NEXT
                    </Button>
                  )}
                />
              </div>
            </div>

            {isScreenView && (
              <div className=" bg-neutral-100/75 min-w-72 md:w-auto w-full h-full px-2 ">
                <QuestionSelectorSection
                  answers={answers}
                  onQuestionNumberClick={handleQuestionChange}
                  activeQuestionIndex={activeIndex}
                />
              </div>
            )}
          </div>
        </div>
      </div>
    </MathJaxContext>
  );
};

export default MCQExamCenter;

import ReactMarkdown, { Components } from "react-markdown";

import remarkGfm from "remark-gfm";
import RemarkMathPlugin from "remark-math";
import { StyledLink } from "../../componnets/base/styled-link";
import { paths } from "../../routes/route.constant";
import QuestionSelectorSection from "./question-selector-section";

function MarkdownRender({ ...props }) {
  const cp: Components = {
    code: (props) => {
      return (
        <MathJax style={{ display: "inline" }}>\({props.children}\)</MathJax>
      );
    },
    table: (props) => {
      return <span className={m.markdown}>{props.children}</span>;
    },
  };
  const newProps = {
    ...props,
    remarkPlugins: [RemarkMathPlugin, remarkGfm],
    components: cp,
  };
  return (
    <MathJaxContext>
      <ReactMarkdown {...newProps} />
    </MathJaxContext>
  );
}
/**
<div className="md:max-h-screen h-screen flex flex-col ">
  <TestHeader currentTime={examTime} active={interval.active} />
  <div className="flex-1 overflow-hidden">
    <div className="  flex flex-col md:flex-row items-start max-h-full">
      <div className="pt-2 flex-1 overflow-hidden  max-h-full flex flex-col justify-start bg-neutral-50/50 border-r-0 md:border-r ">
        <div className="text-start p-2 pl-4 ">
          Que No {activeIndex + 1} / {questionSet?.number_of_questions || 0}
        </div>
        <div className="flex-1  overflow-hidden flex md:flex-row flex-col">
          {activeContentReference && (
            <div className="md:w-1/2 min-w-[50%] overflow-auto p-4 pt-0 md:pt-4  text-pretty font-medium border-r-0 border-b md:border-b-0 md:border-r">
              <p className="font-semibold mb-2">
                {contentGroup[activeContentReference].instructions}
              </p>
              <p>{contentGroup[activeContentReference].content}</p>
            </div>
          )}
          <div className="md:max-h-[70vh] h-full flex flex-col p-4">
            {questions.length && (
              <>
                <p className="items-start text-start pb-2">
                  <MarkdownRender
                    children={questions?.[activeIndex]?.question || ""}
                  />
                </p>

                <Controller
                  name={`answers.${activeIndex}.selectedOption`}
                  control={control}
                  render={({ field }) => (
                    <AnswerOptions
                      name={field.name}
                      onChange={field.onChange}
                      options={questions[activeIndex].options}
                      selected={field.value}
                    />
                  )}
                />
              </>
            )}
          </div>
        </div>
        <div className="min-h-12 flex gap-2 p-2 border-orange-300 border-t">
          <Button size={"sm"}>Mark for Review & Next</Button>
          <Button variant={"outline"} size={"sm"} className="mr-auto">
            Clear
          </Button>
          <Controller
            name={`activeQuestionIndex`}
            control={control}
            render={({ field }) => (
              <Button
                size={"sm"}
                variant={"success"}
                className="justify-self-end"
                type="button"
                disabled={field.value + 1 >= questions.length}
                onClick={() => field.onChange(field.value + 1)}
              >
                SAVE & NEXT
              </Button>
            )}
          />
        </div>
      </div>
      <div className="py-2 bg-neutral-100/75 min-w-72 md:w-auto w-full h-full px-2 ">
        <div className=" min-w-72">
          <QuestionSelector
            questions={answers}
            onQuestionNumberClick={(index) =>
              setValue("activeQuestionIndex", index)
            }
          />
        </div>
      </div>
    </div>
  </div>
</div>;
 */
