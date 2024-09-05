import { useEffect, useState } from "react";
import { getQuestions } from "../../services/exam.service";

import { useNavigate } from "react-router-dom";
import { Checkbox } from "../../componnets/base/checkbox";
import DescriptiveQuestionCard from "../../componnets/shared/descriptive-question-card";
import Header from "../../componnets/shared/header/header";
import { IQuestion } from "../../interface/question";
import { paths } from "../../routes/route.constant";
import { supabase } from "../../supabase/client";

const QuestionPaper = () => {
  const [questions, setQuestions] = useState<IQuestion[]>([]);
  const navigate = useNavigate();
  const getQuestionsList = async () => {
    try {
      await supabase.auth.getSession();
      const data = await getQuestions();
      console.log(data);
      setQuestions(data.data);
    } catch (error) {
      console.log(error);
    }
  };
  useEffect(() => {
    getQuestionsList();
  }, []);

  const attempQuestion = (index: number) => {
    navigate(`/${paths.EXAMS}/${paths.DISCRIPTIVE}/${index}`, {
      state: {
        question: questions[index],
      },
    });
  };
  return (
    <div className="w-full">
      <Header />

      <div className="pt-2 w-full md:max-w-md lg:max-w-lg xl:max-w-xl 2xl:max-w-2xl mx-auto flex flex-col gap-2 p-4">
        <div className="py-2">
          <div className="text-primary">
            Get started with the questions below. Each question is an
            opportunity to test you're knowledge of topics.
          </div>

          <div className=" flex items-center gap-2  text-sm">
            <span>Show unattempted only</span>
            <Checkbox variant={"info"} />
          </div>
        </div>
        {questions.map((item, index) => {
          const attempts = item.max_attempts - item.user_attempts;
          return (
            <DescriptiveQuestionCard
              isNew={item.user_attempts}
              topic={item.raw_exam_data.topic}
              type={item.raw_exam_data.type}
              srNumber={index + 1}
              isAttemped={!attempts}
              handleAttemptClick={() => attempQuestion(item.id)}
              duration={item.duration_seconds / 60}
              attempts={attempts}
            />
          );
        })}
      </div>
    </div>
  );
};

export default QuestionPaper;
