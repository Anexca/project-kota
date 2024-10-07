import { MathJaxContext } from "better-react-mathjax";
import { useEffect, useMemo, useRef, useState } from "react";
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
import MarkdownRender from "../../componnets/shared/markdown-rendere";
import MCQMobileHeader from "../../componnets/shared/mcq-mobile-header";
import MCQQuestionPallet from "../../componnets/shared/mcq-question-pallet";
import { ReadMore } from "../../componnets/shared/read-more-content";
import TestHeader from "../../componnets/shared/test-header/test-header";
import { QUESTION_STATE, ScreenSizeQuery } from "../../constants/shared";
import { useInterval } from "../../hooks/use-interval";
import { useMediaQuery } from "../../hooks/use-media-query";
import { useToast } from "../../hooks/use-toast";
import { MCQFormModal } from "../../interface/mcq-exam";
import { IContentGroup, IMCQQuestionSet } from "../../interface/question";
import { paths } from "../../routes/route.constant";
import { getMCQQuestionById } from "../../services/exam.service";
import { evaluateMcqExam } from "../../services/mcq-exam.services";

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
const AutoSubmitDialog = () => {
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
const IncompleteExamDialog = ({
  open,
  close,
  submitExam,
  stats,
}: {
  open: boolean;
  close: () => void;
  submitExam: () => void;
  stats: { incompleteExam: number; markedForReview: number };
}) => {
  const msg = [];
  if (stats.incompleteExam) {
    msg.push(
      <>
        <span className="text-orange-600">{stats.incompleteExam}</span>{" "}
        unattempted questions
      </>
    );
  }
  if (stats.markedForReview) {
    msg.push(
      <>
        {!!stats.incompleteExam && " and "}{" "}
        <span className="text-orange-600">{stats.markedForReview}</span> marked
        for review
      </>
    );
  }
  return (
    <Dialog open={open}>
      <DialogContent noCloseButton>
        <DialogHeader>
          <DialogTitle>Submit Exam.</DialogTitle>
          <DialogDescription>
            <div className="mb-2 text-neutral-700 font-semibold ">
              You have not completed the exam. You have {msg}. Are you sure you
              want to submit ?
            </div>
            <div className="flex gap-2 justify-center sm:justify-start">
              <Button
                variant={"warning"}
                onClick={() => {
                  submitExam();
                  close();
                }}
                className="p-2 h-8"
              >
                Submit Exam
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

const MCQExamCenter = ({ isOpenMode }: { isOpenMode?: boolean }) => {
  const param = useParams();
  const [startUpModal, setStartUpModal] = useState(true);
  const [exitModelOpen, setExitModelOpen] = useState(false);
  const [incompleteModelOpen, setIncompleteModelOpen] = useState(false);
  const [questionSet, setQuestionSet] = useState<IMCQQuestionSet | null>(null);
  const [examTime, setExamTime] = useState(0);
  const interval = useInterval(() => setExamTime((s) => s - 1), 1000);
  const perQuestionTimeRef = useRef(0);
  const perQuestionTimeInterval = useInterval(
    () => (perQuestionTimeRef.current += 1),
    1000
  );
  const { toast } = useToast();
  const isScreenView = useMediaQuery(ScreenSizeQuery.largeScreen, true);
  const isMobileView = useMediaQuery(ScreenSizeQuery.mediumScreen, true);

  const navigate = useNavigate();

  const { control, setValue, reset, getValues } = useForm<
    MCQFormModal & {
      tempAnswerIndex: number | null;
      sectionList: string[];
      questionTimeMap: Record<string, number>;
    }
  >({
    defaultValues: {
      answers: {},
      activeQuestionIndex: 0,
      tempAnswerIndex: null,
      sectionList: [],
    },
  });
  const activeIndex = useWatch({ control, name: "activeQuestionIndex" });
  const activeSection = useWatch({ control, name: "activeSection" });
  const answers = useWatch({ control, name: "answers" });
  const tempAnswerIndex = useWatch({ control, name: "tempAnswerIndex" });
  const sectionList = useWatch({ control, name: "sectionList" });

  const fetchExam = async () => {
    try {
      const res = await getMCQQuestionById(param.questionId!);
      // setExamTime(5);
      setExamTime(res.data.duration_seconds);
      setQuestionSet(res.data);
      const ques = res.data.raw_exam_data.sections;
      const sets = Object.keys(ques);
      const questionsSet: any = {};
      const questionTimeMap: Record<string, number> = {};
      sets.forEach((item) => {
        const quesArray = ques[item].map((i) => {
          questionTimeMap[i.question_number] = 0;
          return {
            state: QUESTION_STATE.UN_ATTEMPTED,
            selectedOption: null,
            questionNumber: i.question_number,
          };
        });
        questionsSet[item] = quesArray;
      });
      reset({
        activeQuestionIndex: 0,
        answers: questionsSet,
        activeSection: sets[0],
        sectionList: sets,
        questionTimeMap,
      });
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
  const saveTimeForQuestion = () => {
    const currentAnswer = answers?.[activeSection]?.[activeIndex];
    const prevTime = getValues(
      `questionTimeMap.${currentAnswer.questionNumber}`
    );
    perQuestionTimeInterval.stop();
    setValue(
      `questionTimeMap.${currentAnswer.questionNumber}`,
      prevTime + perQuestionTimeRef.current
    );
    perQuestionTimeRef.current = 0;
    perQuestionTimeInterval.start();
    console.log(getValues("questionTimeMap"));
  };
  const handleQuestionChange = (index: number, section: string) => {
    if (index == activeIndex && section == activeSection) return;
    const currentAnswer = answers?.[activeSection]?.[activeIndex];
    if (currentAnswer?.state == "UN-ATTEMPTED") {
      setValue(`answers.${activeSection}.${activeIndex}.state`, "NOT-ANSWERED");
    }
    setValue("activeQuestionIndex", index);
    setValue("activeSection", section);
    saveTimeForQuestion();
  };

  const handleSaveNNext = () => {
    if (tempAnswerIndex) {
      setValue(`answers.${activeSection}.${activeIndex}.state`, "ATTEMPTED");
      setValue(
        `answers.${activeSection}.${activeIndex}.selectedOption`,
        tempAnswerIndex
      );
      setValue("tempAnswerIndex", null);
      saveTimeForQuestion();
    }
  };
  const handleReviewNNext = () => {
    if (tempAnswerIndex) {
      setValue(`answers.${activeSection}.${activeIndex}.state`, "FOR-REVIEW");
      setValue(
        `answers.${activeSection}.${activeIndex}.selectedOption`,
        tempAnswerIndex
      );
      setValue("tempAnswerIndex", null);
      saveTimeForQuestion();
    }
  };

  const questions = questionSet?.raw_exam_data?.sections?.[activeSection] || [];
  const activeContentReference = questions?.[activeIndex]?.content_reference_id;
  const contentInfo = contentGroup[activeContentReference];

  const endExamNavigatePath = isOpenMode
    ? `/${paths.COMMUNITY_EXAMS}/banking/${paths.MCQ}`
    : `/${paths.EXAMS}/banking/${param.categoryId}`;
  const exitExam = () => {
    interval.stop();
    navigate(endExamNavigatePath);
  };

  useEffect(() => {
    if (examTime <= 0 && interval.active) {
      interval.stop();
      handleFormSubmit();
    }
  }, [examTime]);

  const noAttemptLeft = useMemo(() => {
    const hasAttemptLeft =
      (questionSet?.max_attempts || 0) - (questionSet?.user_attempts || 0) ==
        0 && questionSet;
    interval.stop();

    return hasAttemptLeft;
  }, [questionSet]);

  const checkForAllAttempted = () => {
    const allAttempted = sectionList.every((i) =>
      answers[i].every((e) => e.state == QUESTION_STATE.ATTEMPTED)
    );

    if (allAttempted) {
      handleFormSubmit();
      return;
    }
    setIncompleteModelOpen(true);
  };

  const handleFormSubmit = async () => {
    saveTimeForQuestion();
    perQuestionTimeInterval.stop();
    const payload: any = {
      attempted_questions: [],
      completed_seconds: questionSet!.duration_seconds - examTime,
    };
    const timeMap = getValues("questionTimeMap");
    sectionList.forEach((sect) => {
      answers[sect].forEach((i, index) => {
        if (i.selectedOption)
          payload.attempted_questions.push({
            question_number: index + 1,
            user_selected_option_index: [
              Number.parseInt(i.selectedOption as any),
            ],
            time_taken_in_seconds: timeMap[i.questionNumber],
          });
      });
    });

    try {
      const res = await evaluateMcqExam({ id: param.questionId!, payload });
      console.log(res);
    } catch (error) {}
  };
  const { incompleteExam, markedForReview } = useMemo(() => {
    let incompleteExam = 0;
    sectionList.forEach((i) => {
      incompleteExam += answers[i].filter(
        (i) => i.state == "NOT-ANSWERED" || i.state == "UN-ATTEMPTED"
      ).length;
    });
    let markedForReview = 0;
    sectionList.forEach(
      (i) =>
        (markedForReview += answers[i].filter(
          (i) => i.state == "FOR-REVIEW"
        ).length)
    );

    return { incompleteExam, markedForReview };
  }, [incompleteModelOpen]);
  return (
    <MathJaxContext>
      <div className="md:max-h-screen md:h-screen flex flex-col ">
        <TestHeader currentTime={examTime} active={interval.active} />
        <ConformationDialog
          timerStart={() => {
            perQuestionTimeInterval.start();
            interval.start();
          }}
          time={questionSet ? questionSet?.duration_seconds / 60 : 0}
          open={startUpModal && !noAttemptLeft}
          setOpen={setStartUpModal}
        />
        {questionSet && !examTime && <AutoSubmitDialog />}
        <EndExamDialog
          open={exitModelOpen}
          close={() => setExitModelOpen(false)}
          exitExam={exitExam}
        />
        <IncompleteExamDialog
          open={incompleteModelOpen}
          close={() => setIncompleteModelOpen(false)}
          submitExam={handleFormSubmit}
          stats={{ incompleteExam, markedForReview }}
        />

        <div className="flex-1 md:overflow-hidden">
          <div className="flex flex-col md:flex-row items-stretch md:max-h-full md:h-full">
            <div className="flex-1 max-h-full flex flex-col">
              <div className="text-start p-2 pl-4 shadow flex items-center">
                <span className="ml-auto flex-1">
                  Que No {activeIndex + 1} /{" "}
                  {questionSet?.number_of_questions || 0}
                </span>
                {!isScreenView && (
                  <MCQMobileHeader
                    activeSection={activeSection}
                    sectionList={sectionList}
                    answers={answers}
                    onQuestionNumberClick={handleQuestionChange}
                    activeQuestionIndex={activeIndex}
                    handleSubmit={checkForAllAttempted}
                  />
                )}
              </div>
              <div className="flex-1 flex flex-col md:flex-row md:overflow-hidden">
                {contentInfo && (
                  <div className="md:w-1/2 min-w-[50%] overflow-auto p-4 pt-0 md:pt-4  text-pretty font-medium border-r-0 border-b md:border-b-0 md:border-r">
                    {contentInfo?.instructions && (
                      <MarkdownRender className="font-semibold mb-2 pt-4 md:pt-0">
                        {contentInfo.instructions}
                      </MarkdownRender>
                    )}
                    {contentInfo?.content && (
                      <>
                        {isMobileView ? (
                          <MarkdownRender>{contentInfo.content}</MarkdownRender>
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
                <Button
                  onClick={() => {
                    setValue("tempAnswerIndex", null);
                  }}
                  variant={"outline"}
                  size={"sm"}
                  className="mr-auto"
                >
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
                      disabled={tempAnswerIndex == null}
                      onClick={() => {
                        const nextIdx = field.value + 1;
                        handleSaveNNext();
                        field.value + 1 <= questions.length &&
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
                <MCQQuestionPallet
                  activeSection={activeSection}
                  sectionList={sectionList}
                  answers={answers}
                  onQuestionNumberClick={handleQuestionChange}
                  activeQuestionIndex={activeIndex}
                  handleSubmit={checkForAllAttempted}
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
