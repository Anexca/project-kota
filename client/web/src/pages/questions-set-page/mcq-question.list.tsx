import { useMemo, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { IMCQExam } from "../../interface/question";
import { paths } from "../../routes/route.constant";

import { Button } from "../../componnets/base/button/button";
import Chip from "../../componnets/base/chip";
import Icon from "../../componnets/base/icon";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../../componnets/base/select";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
} from "../../componnets/base/sheet";
import { StyledLink } from "../../componnets/base/styled-link";
import MCQPreviousSubmissions from "../../componnets/shared/mcq-previous-submissions-list";
import NoPremiumBanner from "../../componnets/shared/no-premium-banner";
import useUserProfileStore from "../../store/user-info-store";

export function ViewSubmissionDrawer({
  open,
  setOpen,
  question,
  isOpenExam,
}: {
  open: boolean;
  setOpen: (b: boolean) => void;
  question: IMCQExam | null;
  isOpenExam?: boolean;
}) {
  return (
    <Sheet open={open} onOpenChange={setOpen}>
      <SheetContent
        className="flex flex-col overflow-hidden"
        autoFocus={false}
        onOpenAutoFocus={(e) => e.preventDefault()}
      >
        <SheetHeader>
          <SheetTitle>Submissions</SheetTitle>
        </SheetHeader>
        {question && (
          <MCQPreviousSubmissions question={question} isOpenExam={isOpenExam} />
        )}
        {question && isOpenExam && <>yo</>}
      </SheetContent>
    </Sheet>
  );
}

const selectOptions = {
  label: "All",
  value: "all",
};

type Props = { isOpenMode?: boolean; questions: IMCQExam[] };

const MCQQuestionsList = ({ isOpenMode, questions }: Props) => {
  const [filterType, setFilterType] = useState("all");
  const { profile } = useUserProfileStore();

  const [selectedQuestion, setSelectedQuestions] = useState<IMCQExam | null>(
    null
  );
  const params = useParams();

  const navigate = useNavigate();

  const attempQuestion = (index: number) => {
    const path = isOpenMode
      ? `/${paths.COMMUNITY_EXAMS}/banking/${paths.MCQ}`
      : `/${paths.EXAMS}/banking/${params.categoryId}/${paths.MCQ}/${index}`;
    navigate(path);
  };

  const questionList = useMemo(() => {
    const temp: any = {};
    let key = "";
    questions.forEach((i) => {
      key = `${i.exam_name}$$${i.exam_stage}`;
      if (temp[key]) {
        temp[key].push(i);
      } else {
        temp[key] = [i];
      }
    });
    return Object.keys(temp).map((i) => {
      const [exam_name, exam_stage] = i.split("$$");
      return {
        exam_name: exam_name.toLowerCase().split("_").join(" "),
        exam_stage,
        items: temp[i],
      };
    });
  }, [questions, filterType]);

  const filterOptions = useMemo(() => {
    const set = new Set(questionList.map((i) => i.exam_name));

    return [...set];
  }, [questionList]);

  const filteredQuestions = useMemo(() => {
    if (filterType) {
      return questionList.filter((i) => {
        let condition = true;

        if (filterType !== "all") {
          condition = filterType == i.exam_name;
        }
        return condition;
      });
    }
    return questionList;
  }, [questionList, filterType]);

  return (
    <>
      <div className="py-2">
        <div className="text-sm text-black font-medium mb-2">
          Get started with the questions below.
        </div>
        <ViewSubmissionDrawer
          open={!!selectedQuestion}
          setOpen={() => setSelectedQuestions(null)}
          question={selectedQuestion}
          isOpenExam={isOpenMode}
        />
        <div className="flex items-center gap-2  text-sm">
          <span className="">
            <Select value={filterType} onValueChange={setFilterType}>
              <SelectTrigger className="w-[180px] text-sm h-8 capitalize">
                <SelectValue placeholder="Filter" className="capitalize" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectItem
                    onClick={(e) => e.stopPropagation()}
                    className="!text-sm"
                    value={selectOptions.value}
                  >
                    {selectOptions.label}
                  </SelectItem>
                  {filterOptions.map((i) => (
                    <SelectItem
                      onClick={(e) => e.stopPropagation()}
                      className="!text-sm capitalize"
                      value={i}
                    >
                      {i}
                    </SelectItem>
                  ))}
                </SelectGroup>
              </SelectContent>
            </Select>
          </span>
        </div>
      </div>

      <div className="animate-fadeIn flex flex-col gap-4">
        {filteredQuestions.map(({ items, exam_name, exam_stage }) => {
          return (
            <div className="rounded-md shadow-sm p-2 bg-info/10">
              <div className="mb-2 mx-2">
                Section -{" "}
                <span className="font-medium capitalize  mr-2">
                  {exam_name.toLowerCase().split("_").join(" ")}
                </span>
                <Chip
                  icon="tags"
                  variant={"outline"}
                  className="text-md px-3 capitalize"
                >
                  {exam_stage.toLowerCase()}
                </Chip>
              </div>
              <div className=" flex flex-col gap-2">
                {items.map((item: IMCQExam) => {
                  const attempts = item.max_attempts - item.user_attempts;
                  return (
                    <MCQQuestionCard
                      key={item.exam_id}
                      negativeMarking={item.negative_marking}
                      srNumber={item.exam_id}
                      isAttemped={!!item.user_attempts}
                      handleAttemptClick={() => attempQuestion(item.exam_id)}
                      duration={item.duration_seconds / 60}
                      attempts={attempts}
                      showSubmission={() => setSelectedQuestions(item)}
                      numberOfQuestion={item.number_of_questions}
                    />
                  );
                })}
              </div>
            </div>
          );
        })}
        {isOpenMode && (
          <>
            {!profile?.active_subscriptions?.length ? (
              <NoPremiumBanner />
            ) : (
              <div className="flex p-4 rounded bg-white shadow-sm mt-2 gap-4">
                <div className=" w-2 bg-info rounded-full"></div>
                <div className="flex flex-col flex-1">
                  <div className="text-sm font-semibold">
                    Explore more descriptive question.
                  </div>
                  <div>
                    <StyledLink
                      to={`/${paths.EXAMS}/banking/${paths.DISCRIPTIVE}`}
                      size={"sm"}
                      className="px-3 py-1 h-7 mt-2"
                      variant={"info"}
                    >
                      <Icon icon="send" className="mr-2" /> See More
                    </StyledLink>
                  </div>
                </div>
              </div>
            )}
          </>
        )}
      </div>
    </>
  );
};

export default MCQQuestionsList;

type CardProps = {
  srNumber?: number | null;

  negativeMarking: number;
  isAttemped?: boolean;
  handleAttemptClick?: () => void;
  showSubmission?: () => void;
  duration?: number;
  attempts: number;
  numberOfQuestion: number;
};
const MCQQuestionCard = ({
  srNumber = null,
  numberOfQuestion,
  negativeMarking,
  isAttemped,
  handleAttemptClick,
  showSubmission,
  duration,
  attempts,
}: CardProps) => {
  return (
    <article className="rounded-md shadow-sm bg-white flex flex-col md:flex-row gap-4 p-3 px-4 md:p3 text-sm">
      <div className="flex-1">
        <p className="font-medium text-balance md:text-pretty text-black mb-2">
          Question Paper Set - #{srNumber}
        </p>
        <div className="flex gap-2 flex-wrap">
          <Chip icon="clock" variant={"info"}>
            {duration?.toFixed(0)} min
          </Chip>
          <Chip icon="minus_circle" variant={"danger"}>
            {negativeMarking} Points
          </Chip>
          <Chip icon="rotate_right" variant={"warning"}>
            {attempts} Attempt
          </Chip>
          <Chip icon="target" variant={"success"}>
            {numberOfQuestion} Ques
          </Chip>
        </div>
      </div>
      <div className="flex  items-stretch gap-2 md:justify-center ">
        {attempts >= 1 && (
          <Button
            onClick={handleAttemptClick}
            size={"sm"}
            className="px-3 py-1"
            variant={"info"}
          >
            <Icon icon="play_circle" className="mr-2" /> Attempt
          </Button>
        )}
        {isAttemped && (
          <Button
            onClick={showSubmission}
            size={"sm"}
            className="px-3 py-1"
            variant={"secondary"}
          >
            <Icon icon="send" className="mr-2" /> View Submission
          </Button>
        )}
      </div>
    </article>
  );
};
