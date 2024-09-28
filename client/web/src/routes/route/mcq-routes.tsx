import { RouteObject } from "react-router-dom";
import { ExamDomain, paths } from "../route.constant";
import MCQQuestionCategories from "../../pages/mcq-question-categories";
import MCQQuestions from "../../pages/mcq-questions";

export const mcqRoutes: RouteObject[] = [
  {
    path: `${ExamDomain.banking}/${paths.MCQ}`,
    element: <MCQQuestionCategories />,
  },
  {
    path: `${ExamDomain.banking}/${paths.MCQ}/:categoryId`,
    element: <MCQQuestions />,
  },
];
