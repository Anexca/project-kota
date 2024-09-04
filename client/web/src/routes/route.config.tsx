import { RouteObject } from "react-router-dom";
import { paths } from "./route.constant";
import { Login } from "../pages/login/login";
import { RegisterPage } from "../pages/register/register";
import HomePage from "../pages/homepage/homepage";
import ProtectedRoute from "./protected-route";
import { ForgotPassword } from "../pages/forgot-password/forgot-password";
import SupabaseAuth from "../pages/supabase-auth/supabase-auth";
import QuestionPaper from "../pages/question-paper/question-paper";
import DiscriptiveExam from "../pages/discriptive-exam/discriptive-exam";

const routes: RouteObject[] = [
  {
    path: paths.HOMEPAGE,
    element: (
      <ProtectedRoute>
        <HomePage />
      </ProtectedRoute>
    ),
  },
  {
    path: paths.REGISTER,
    element: <RegisterPage />,
  },
  {
    path: paths.LOGIN,
    element: <Login />,
  },
  {
    path: paths.FORGOT_PASSWORD,
    element: <ForgotPassword />,
  },
  {
    path: paths.SUPABASE_AUTH,
    element: <SupabaseAuth />,
  },
  {
    path: paths.QUESTION_PAPER,
    element: <QuestionPaper />,
  },
  {
    path: paths.EXAMS,
    children: [
      {
        path: `${paths.DISCRIPTIVE}/:questionId`,
        element: <DiscriptiveExam />,
      },
    ],
  },
];

export default routes;
