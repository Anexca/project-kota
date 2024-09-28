import { useEffect, useMemo, useState } from "react";
import { getMCQQuestions } from "../../services/exam.service";

import { Link, useNavigate, useParams } from "react-router-dom";
import { Checkbox } from "../../componnets/base/checkbox";
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
import NoPremiumBanner from "../../componnets/shared/no-premium-banner";
import { questionType } from "../../constants/shared";
import { useToast } from "../../hooks/use-toast";
import useUserProfileStore from "../../store/user-info-store";

export function ViewSubmissionDrawer({
  open,
  setOpen,
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
        {/* {question && (
          <PreviousSubmissions question={question} isOpenExam={isOpenExam} />
        )} */}
      </SheetContent>
    </Sheet>
  );
}

const selectOptions = [
  {
    label: "All",
    value: "all",
  },
  {
    label: "Essay",
    value: "essay",
  },
  {
    label: "Formal Letter",
    value: "formal_letter",
  },
];

const MCQQuestions = ({ isOpenMode }: { isOpenMode?: boolean }) => {
  const [filterType, setFilterType] = useState("all");
  const { profile } = useUserProfileStore();

  const [questions, setQuestions] = useState<IMCQExam[]>([]);
  const [selectedQuestion, setSelectedQuestions] = useState<IMCQExam | null>(
    null
  );
  const params = useParams();
  const [showUnattempted, setShowUnattempted] = useState(false);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const { toast } = useToast();
  const getQuestionsList = async () => {
    setLoading(true);
    try {
      const data = await getMCQQuestions({
        categoryId: params.categoryId! || "open",
      });
      setQuestions(data.data);
    } catch (error) {
      toast({
        title: "Oh ho Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in processing your request.",
      });
    }
    setLoading(false);
  };
  useEffect(() => {
    getQuestionsList();
  }, []);

  const attemptQuestion = (index: number) => {
    const path = isOpenMode
      ? `/${paths.COMMUNITY_EXAMS}/banking/${paths.MCQ}`
      : `/${paths.EXAMS}/banking/${paths.MCQ}/${params.categoryId}/${index}`;
    navigate(path);
  };

  const questionList = useMemo(() => {
    if (filterType || showUnattempted) {
      return questions.filter((i) => {
        let condition = true;
        if (showUnattempted) {
          condition = !i.user_attempts;
        }
        if (filterType !== "all") {
          condition = condition && filterType == i.exam_type;
        }
        return condition;
      });
    }
    return questions;
  }, [questions, filterType, showUnattempted]);
  const backPath = isOpenMode
    ? `/${paths.HOMEPAGE}`
    : `/${paths.EXAMS}/banking/${paths.DISCRIPTIVE}`;
  return (
    <div className="pt-2 w-full md:max-w-2xl 2xl:max-w-2xl mx-auto flex flex-col gap-2 p-4">
      <div className="py-2">
        <div className="flex gap-2 items-center">
          <Link to={backPath} className="p-0">
            <Icon icon="arrow_back" className="text-info text-lg" />
          </Link>
          <span className="text-lg font-semibold">Banking MCQ Question</span>
        </div>
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
          <span>Show unattempted only</span>
          <Checkbox
            variant={"info"}
            onCheckedChange={(s: boolean) => setShowUnattempted(s)}
          />
          <span className="ml-auto">
            <Select value={filterType} onValueChange={setFilterType}>
              <SelectTrigger className="w-[180px] text-sm h-8">
                <SelectValue placeholder="Filter" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  {selectOptions.map((i) => (
                    <SelectItem className="!text-sm" value={i.value}>
                      {i.label}
                    </SelectItem>
                  ))}
                </SelectGroup>
              </SelectContent>
            </Select>
          </span>
        </div>
      </div>
      {loading ? (
        <div className="flex flex-col gap-2 justify-center items-center">
          <span className="rounded-full w-8 h-8 animate-spin border-2 border-info border-t-info/30"></span>
          Getting Exciting Questions
        </div>
      ) : (
        <div className="animate-fadeIn flex flex-col gap-2">
          {questionList.map((item) => {
            const attempts = item.max_attempts - item.user_attempts;
            return (
              <DescriptiveQuestionCard
                key={item.id}
                topic={item.exam_name}
                type={item.exam_type}
                srNumber={item.id}
                isAttemped={!!item.user_attempts}
                handleAttemptClick={() => attemptQuestion(item.id)}
                duration={item.duration_seconds / 60}
                attempts={attempts}
                showSubmission={() => setSelectedQuestions(item)}
              />
            );
          })}
          {isOpenMode && !profile?.active_subscriptions?.length ? (
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
        </div>
      )}
    </div>
  );
};

export default MCQQuestions;

type Props = {
  srNumber?: number | null;
  topic: string;
  type: string;
  isAttemped?: boolean;
  handleAttemptClick?: () => void;
  showSubmission?: () => void;
  duration?: number;
  attempts: number;
};

const DescriptiveQuestionCard = ({
  srNumber = null,
  topic,
  type,
  isAttemped,
  handleAttemptClick,
  showSubmission,
  duration,
  attempts,
}: Props) => {
  return (
    <article className="rounded-md shadow-sm bg-white flex flex-col md:flex-row gap-4 p-3 px-4 md:p3 text-sm">
      <div className="flex-1">
        <p className="font-medium text-balance md:text-pretty text-black mb-2">
          #{srNumber} - {topic}
        </p>
        <div className="flex gap-2 flex-wrap">
          <Chip icon="clock" variant={"danger"}>
            {duration} min
          </Chip>
          <Chip icon="file" variant={"success"}>
            {questionType[type] || "--"}
          </Chip>
          <Chip icon="rotate_right" variant={"warning"}>
            {attempts} Attempt
          </Chip>
        </div>
      </div>
      <div className="flex flex-col items-stretch gap-2 md:justify-center md:w-32">
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
