import { useEffect, useMemo, useState } from "react";
import { getQuestions } from "../../services/exam.service";

import { useNavigate } from "react-router-dom";
import { Checkbox } from "../../componnets/base/checkbox";
import DescriptiveQuestionCard from "../../componnets/shared/descriptive-question-card";
import { IQuestion } from "../../interface/question";
import { paths } from "../../routes/route.constant";
import { supabase } from "../../supabase/client";

import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
} from "../../componnets/base/sheet";
import PreviousSubmissions from "../../componnets/shared/previous-submissions-list";
import { useToast } from "../../hooks/use-toast";

export function ViewSubmissionDrawer({
  open,
  setOpen,
  question,
}: {
  open: boolean;
  setOpen: (b: boolean) => void;
  question: IQuestion | null;
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
        {question && <PreviousSubmissions question={question} />}
      </SheetContent>
    </Sheet>
  );
}

const DescriptiveQuestion = () => {
  const [questions, setQuestions] = useState<IQuestion[]>([]);
  const [selectedQuestion, setSelectedQuestions] = useState<IQuestion | null>(
    null
  );
  const [showUnattempted, setShowUnattempted] = useState(false);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const { toast } = useToast();
  const getQuestionsList = async () => {
    setLoading(true);
    try {
      await supabase.auth.getSession();
      const data = await getQuestions();
      setQuestions(data.data);
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
    getQuestionsList();
  }, []);

  const attempQuestion = (index: number) => {
    navigate(`/${paths.EXAMS}/banking/${paths.DISCRIPTIVE}/${index}`, {
      state: {
        question: questions[index],
      },
    });
  };

  const unattemptedQuestion = useMemo(() => {
    return questions.filter((i) => !i.user_attempts);
  }, [questions]);
  const questionList = showUnattempted ? unattemptedQuestion : questions;
  return (
    <div className="pt-2 w-full md:max-w-2xl 2xl:max-w-2xl mx-auto flex flex-col gap-2 p-4">
      <div className="py-2">
        <div className="text-sm text-black font-medium mb-2">
          Get started with the questions below.
        </div>
        <ViewSubmissionDrawer
          open={!!selectedQuestion}
          setOpen={() => setSelectedQuestions(null)}
          question={selectedQuestion}
        />
        <div className="flex items-center gap-2  text-sm">
          <span>Show unattempted only</span>
          <Checkbox
            variant={"info"}
            onCheckedChange={(s: boolean) => setShowUnattempted(s)}
          />
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
                topic={item.raw_exam_data.topic}
                type={item.raw_exam_data.type}
                srNumber={item.id}
                isAttemped={!!item.user_attempts}
                handleAttemptClick={() => attempQuestion(item.id)}
                duration={item.duration_seconds / 60}
                attempts={attempts}
                showSubmission={() => setSelectedQuestions(item)}
                // showSubmission={() => showSubmission(item.id)}
              />
            );
          })}
        </div>
      )}
    </div>
  );
};

export default DescriptiveQuestion;
