import { RouteObject } from "react-router-dom";
import { ExamDomain, paths } from "../route.constant";
import MCQQuestionCategories from "../../pages/mcq-question-categories";
import MCQQuestions from "../../pages/mcq-questions";
import MCQSubmission from "../../pages/mcq-submission";

export const mcqRoutes: RouteObject[] = [
  {
    path: `${ExamDomain.banking}/${paths.MCQ}`,
    element: <MCQQuestionCategories />,
  },
  {
    path: `${ExamDomain.banking}/${paths.MCQ}/:categoryId`,
    element: <MCQQuestions />,
  },
  // {
  //   path: `${ExamDomain.banking}/${paths.MCQ}/:questionId/${paths.SUBMISSION}/:assesmentId`,
  //   element: <MCQSubmission />,
  // },
  {
    path: `${paths.MY_SUMBISSIONS}/${paths.MCQ}/:questionId/${paths.SUBMISSION}/:assesmentId`,
    element: <MCQSubmission />,
  },
];
