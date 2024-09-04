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
      <div></div>
      <div className="pt-20 w-full md:max-w-md lg:max-w-lg xl:max-w-xl 2xl:max-w-2xl mx-auto flex flex-col gap-2 p-4">
        <div className="py-2">
          <div className="text-primary">
            Get started with the questions below. Each question is an
            opportunity to test you're knowledge of topics.
          </div>
          <div className="text-sm font-semibold mt-1">
            {" "}
            Each Question Duration :
            <span className="text-orange-400 ml-2">
              <i className="fa-regular fa-clock "></i> 15min
            </span>
          </div>
          <div className=" flex items-center gap-2  text-sm">
            <span>Show unattempted only</span>
            <Checkbox variant={"info"} />
          </div>
        </div>
        {questions.map((item, index) => {
          const isAttemped = item.max_attempts == item.user_attempts;
          return (
            <DescriptiveQuestionCard
              topic={item.raw_exam_data.topic}
              type={item.raw_exam_data.type}
              srNumber={index + 1}
              isAttemped={isAttemped}
              handleAttemptClick={() => attempQuestion(item.id)}
              duration={item.duration_minutes / 60}
            />
          );
        })}
      </div>
    </div>
  );
};

export default QuestionPaper;
