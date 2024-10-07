import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { IMCQExam, IQuestion } from "../../interface/question";
import { paths } from "../../routes/route.constant";

import Icon from "../../componnets/base/icon";
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from "../../componnets/base/tabs";
import { useToast } from "../../hooks/use-toast";
import { IExamCategories } from "../../interface/exam-categories";
import {
  getCategoryById,
  getQuestions,
} from "../../services/exam-categories.services";
import DescriptiveQuestionsList from "./descriptive-question.list";
import MCQQuestionsList from "./mcq-question.list";

const QuestionSetPage = ({ isOpenMode }: { isOpenMode?: boolean }) => {
  const [questions, setQuestions] = useState<{
    DESCRIPTIVE: IQuestion[];
    MCQ: IMCQExam[];
  }>({ DESCRIPTIVE: [], MCQ: [] });
  const [category, setCategory] = useState<IExamCategories | null>(null);
  const params = useParams();

  const [loading, setLoading] = useState(false);

  const { toast } = useToast();
  const getQuestionsList = async () => {
    setLoading(true);
    try {
      const [data, catdata] = await Promise.all([
        getQuestions({
          categoryId: params.categoryId! || "open",
        }),
        getCategoryById({ examGroupId: params.categoryId! }),
      ]);

      const ques = {
        DESCRIPTIVE: data.data["DESCRIPTIVE"] || [],
        MCQ: data.data["MCQ"] || [],
      };
      setQuestions(ques);
      setCategory(catdata.data);
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

  const backPath = isOpenMode
    ? `/${paths.HOMEPAGE}`
    : `/${paths.EXAMS}/banking`;
  return (
    <div className="pt-2 w-full md:max-w-2xl 2xl:max-w-2xl mx-auto flex flex-col gap-2 p-4">
      <div className="py-2">
        <div className="flex gap-2 items-center">
          <Link to={backPath} className="p-0">
            <Icon icon="arrow_back" className="text-info text-lg" />
          </Link>
          <span className="px-2">
            <img
              className="h-10 aspect-square object-contain"
              src={category?.logo_url}
            />
          </span>{" "}
          <span className="text-lg font-semibold">
            {category?.exam_group_name?.split("_").join(" ")}
          </span>
        </div>
      </div>

      {loading ? (
        <div className="flex flex-col gap-2 justify-center items-center">
          <span className="rounded-full w-8 h-8 animate-spin border-2 border-info border-t-info/30"></span>
          Getting Exciting Questions
        </div>
      ) : (
        <>
          <Tabs defaultValue="Descriptive" className="w-full">
            <TabsList>
              <TabsTrigger value="Descriptive">Descriptive</TabsTrigger>
              <TabsTrigger value="MCQ">MCQ</TabsTrigger>
            </TabsList>
            <TabsContent value="Descriptive">
              {!!questions.DESCRIPTIVE?.length && (
                <DescriptiveQuestionsList
                  questions={questions.DESCRIPTIVE}
                  isOpenMode={isOpenMode}
                />
              )}
            </TabsContent>
            <TabsContent value="MCQ">
              {!!questions.MCQ.length && (
                <MCQQuestionsList
                  // questions={exams}
                  questions={questions.MCQ}
                  isOpenMode={isOpenMode}
                />
              )}
            </TabsContent>
          </Tabs>
        </>
      )}
    </div>
  );
};

export default QuestionSetPage;
