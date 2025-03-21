import { useEffect, useMemo, useState } from "react";
import { getQuestions } from "../../services/exam.service";

import { Link, useNavigate, useParams } from "react-router-dom";
import { Checkbox } from "../../componnets/base/checkbox";
import DescriptiveQuestionCard from "../../componnets/shared/descriptive-question-card";
import { IQuestion } from "../../interface/question";
import { paths } from "../../routes/route.constant";

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
import NoPremiumBanner from "../../componnets/shared/no-premium-banner";
import PreviousSubmissions from "../../componnets/shared/previous-submissions-list";
import { useToast } from "../../hooks/use-toast";
import useUserProfileStore from "../../store/user-info-store";
import { StyledLink } from "../../componnets/base/styled-link";

export function ViewSubmissionDrawer({
  open,
  setOpen,
  question,
  isOpenExam,
}: {
  open: boolean;
  setOpen: (b: boolean) => void;
  question: IQuestion | null;
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
          <PreviousSubmissions question={question} isOpenExam={isOpenExam} />
        )}
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

const DescriptiveQuestion = ({ isOpenMode }: { isOpenMode?: boolean }) => {
  const [filterType, setFilterType] = useState("all");
  const { profile } = useUserProfileStore();

  const [questions, setQuestions] = useState<IQuestion[]>([]);
  const [selectedQuestion, setSelectedQuestions] = useState<IQuestion | null>(
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
      const data = await getQuestions({
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

  const attempQuestion = (index: number) => {
    const path = isOpenMode
      ? `/${paths.COMMUNITY_EXAMS}/banking/${paths.DISCRIPTIVE}/${index}`
      : `/${paths.EXAMS}/banking/${paths.DISCRIPTIVE}/${params.categoryId}/${index}`;
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
          condition = condition && filterType == i.raw_exam_data.type;
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
          <span className="text-lg font-semibold">
            Banking Descriptive Question
          </span>
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
                key={item.exam_id}
                topic={item.raw_exam_data.topic}
                type={item.raw_exam_data.type}
                srNumber={item.exam_id}
                isAttemped={!!item.user_attempts}
                handleAttemptClick={() => attempQuestion(item.exam_id)}
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
                    to={`/${paths.EXAMS}/banking`}
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

export default DescriptiveQuestion;
