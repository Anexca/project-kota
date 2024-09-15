import { useEffect, useState } from "react";
import { getPastAttemptedSubmissions } from "../../services/exam.service";

import { Link } from "react-router-dom";
import { paths } from "../../routes/route.constant";
import { supabase } from "../../supabase/client";

import Icon from "../../componnets/base/icon";
import { useToast } from "../../hooks/use-toast";
import { IPastExamAttempt } from "../../interface/past-submission";
import AttemptedSubmissionsCard from "../../componnets/shared/attempted-submissions-card";

const PreviousSubmissionPage = () => {
  const [submisions, setQuestions] = useState<IPastExamAttempt[]>([]);
  const [loading, setLoading] = useState(false);
  const { toast } = useToast();
  const getQuestionsList = async () => {
    setLoading(true);
    try {
      await supabase.auth.getSession();
      const data = await getPastAttemptedSubmissions();
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

  return (
    <div className="pt-2 w-full md:max-w-2xl 2xl:max-w-2xl mx-auto flex flex-col gap-2 p-4">
      <div className="py-2">
        <div className="flex gap-2 items-center">
          <Link to={`/${paths.HOMEPAGE}`} className="p-0">
            <Icon icon="arrow_back" className="text-info text-lg" />
          </Link>
          <span className="text-lg font-semibold">
            All of your previously attempted submissions.
          </span>
        </div>
        <div className="flex items-center gap-2  text-sm"></div>
      </div>
      {loading ? (
        <div className="flex flex-col gap-2 justify-center items-center">
          <span className="rounded-full w-8 h-8 animate-spin border-2 border-info border-t-info/30"></span>
          Getting Your Submissions
        </div>
      ) : (
        <div className="animate-fadeIn flex flex-col gap-2">
          {submisions.map((item) => {
            return (
              <AttemptedSubmissionsCard
                key={item.attempted_exam_id}
                topic={item.topic}
                type={item.type}
                srNumber={item.attempted_exam_id}
                category={item.exam_category}
                submissions={item.attempts}
              />
            );
          })}
        </div>
      )}
    </div>
  );
};

export default PreviousSubmissionPage;
