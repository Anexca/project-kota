import { MathJaxContext } from "better-react-mathjax";
import { useState } from "react";

import { ScreenSizeQuery } from "../../../constants/shared";
import { useMediaQuery } from "../../../hooks/use-media-query";
import { MCQAnswersModel } from "../../../interface/mcq-exam";
import { IContentGroup, IMCQQuestionSet } from "../../../interface/question";
import { Button } from "../../base/button/button";
import Chip from "../../base/chip";
import { StyledLink } from "../../base/styled-link";
import MarkdownRender from "../markdown-rendere";
import MCQMobileHeader from "../mcq-mobile-header";
import MCQQuestionPallet from "../mcq-question-pallet";
import { ReadMore } from "../read-more-content";
import OptionList from "./options-list";

const MCQSubmissionView = ({
  answers,
  numberOfQuestions,
  sectionList,
  questionSet,
  contentGroup,
  perQuestionMarks,
  negativeMark,
  backLink,
}: {
  isOpenMode?: boolean;
  assessment?: any[] | null;
  answers: Record<
    string,
    (MCQAnswersModel & { timeTaken: number; correctAnswer: number })[]
  >;
  numberOfQuestions: number;
  sectionList: string[];
  contentGroup: Record<string, IContentGroup>;
  questionSet: IMCQQuestionSet | null;
  perQuestionMarks: number;
  negativeMark: number;
  backLink: string;
}) => {
  const isScreenView = useMediaQuery(ScreenSizeQuery.largeScreen, true);
  const isMobileView = useMediaQuery(ScreenSizeQuery.mediumScreen, true);

  const [activeSection, setActiveSection] = useState(sectionList[0]);
  const [activeIndex, setActiveIndex] = useState(0);

  const isLastQuestionOfExam =
    sectionList?.[sectionList?.length - 1] == activeSection &&
    activeIndex == answers[activeSection]?.length - 1;

  const isFirstQuestionOfExam =
    sectionList?.[0] == activeSection && activeIndex == 0;

  const handleQuestionChange = (index: number, section: string) => {
    setActiveIndex(index);
    setActiveSection(section);
  };
  const moveToNextQuestion = () => {
    const isEndOfSection = answers[activeSection].length - 1 == activeIndex;
    const sectionIndex = sectionList.findIndex((e) => e == activeSection);
    const isLastSetion = sectionList.length - 1 == sectionIndex;
    setActiveIndex(isEndOfSection ? 0 : activeIndex + 1);
    isEndOfSection &&
      !isLastSetion &&
      setActiveSection(sectionList[sectionIndex + 1]);
  };
  const moveToPrevQuestion = () => {
    const isStartOfSection = activeIndex == 0;
    const sectionIndex = sectionList.findIndex((e) => e == activeSection);
    const isStartSection = sectionIndex == 0;

    setActiveIndex(
      isStartOfSection && isStartSection
        ? 0
        : isStartOfSection
        ? answers[sectionList[sectionIndex - 1]].length - 1
        : activeIndex - 1
    );
    isStartOfSection &&
      !isStartSection &&
      setActiveSection(sectionList[sectionIndex - 1]);
  };
  const handleNext = () => {
    !isLastQuestionOfExam && moveToNextQuestion();
  };
  const handlePrev = () => {
    !isFirstQuestionOfExam && moveToPrevQuestion();
  };

  const questions = questionSet?.raw_exam_data?.sections?.[activeSection] || [];
  const activeContentReference = questions?.[activeIndex]?.content_reference_id;
  const contentInfo = contentGroup[activeContentReference];
  const attemptedCurrent =
    answers?.[activeSection]?.[activeIndex].state == "ATTEMPTED";
  return (
    <MathJaxContext>
      <div
        className={
          "md:max-h-[calc(100vh_-_53px)] md:h-[calc(100vh_-_57px)] lg:max-h-[calc(100vh_-_54px)] lg:h-[calc(100vh_-_54px)] flex flex-col z-50"
        }
      >
        <div className="flex-1 md:overflow-hidden">
          <div className="flex flex-col md:flex-row items-stretch md:max-h-full md:h-full">
            <div className="flex-1 max-h-full flex flex-col">
              <div className="text-start p-2 pl-4 shadow flex bg-white items-center z-10 sticky md:static top-[48px]">
                <span className="ml-auto flex-1 ">
                  Que No {answers[activeSection]?.[activeIndex].questionNumber}{" "}
                  / {numberOfQuestions || 0}
                </span>
                {!isScreenView && (
                  <MCQMobileHeader
                    activeSection={activeSection}
                    sectionList={sectionList}
                    answers={answers}
                    onQuestionNumberClick={handleQuestionChange}
                    activeQuestionIndex={activeIndex}
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
                  {!!questions.length && (
                    <>
                      <p className="items-start text-start pb-2">
                        <MarkdownRender
                          children={questions?.[activeIndex]?.question || ""}
                        />
                      </p>

                      <div className=" mb-4 flex gap-2">
                        {attemptedCurrent && (
                          <span className="font-semibold text-sm">
                            Time Taken{" "}
                            <Chip icon="clock">
                              {
                                answers?.[activeSection]?.[activeIndex]
                                  .timeTaken
                              }{" "}
                              sec
                            </Chip>
                          </span>
                        )}
                        <span className="font-semibold text-sm">
                          Marks{" "}
                          <Chip icon="target">
                            {attemptedCurrent &&
                            answers?.[activeSection]?.[activeIndex]
                              .correctAnswer ==
                              answers?.[activeSection]?.[activeIndex]
                                .selectedOption
                              ? perQuestionMarks
                              : `-${negativeMark}`}
                            {"/"}
                            {perQuestionMarks}
                          </Chip>
                        </span>
                      </div>

                      <OptionList
                        selected={
                          answers?.[activeSection]?.[activeIndex]
                            ?.selectedOption
                        }
                        correctAnswer={questions[activeIndex].answer[0]}
                        options={questions[activeIndex].options}
                      />
                    </>
                  )}
                </div>
              </div>
              <div className="min-h-12 flex gap-2 p-2 border-orange-300 border-t sticky bottom-0 bg-white">
                <StyledLink size={"sm"} type="button" to={backLink}>
                  Exit Solution
                </StyledLink>

                <Button
                  disabled={isFirstQuestionOfExam}
                  size={"sm"}
                  variant={"warning"}
                  className="justify-self-end"
                  type="button"
                  onClick={() => handlePrev()}
                >
                  Prev
                </Button>
                <Button
                  disabled={isLastQuestionOfExam}
                  size={"sm"}
                  variant={"success"}
                  className="justify-self-end"
                  type="button"
                  onClick={() => handleNext()}
                >
                  Next
                </Button>
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
                />
              </div>
            )}
          </div>
        </div>
      </div>
    </MathJaxContext>
  );
};

export default MCQSubmissionView;
