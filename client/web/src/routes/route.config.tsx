import { Navigate, RouteObject } from "react-router-dom";
import { paths } from "./route.constant";
import { Login } from "../pages/login/login";
import { RegisterPage } from "../pages/register/register";
import HomePage from "../pages/homepage/homepage";
import ProtectedRoute from "./protected-route";
import { ForgotPassword } from "../pages/forgot-password/forgot-password";
import QuestionPaper from "../pages/question-paper/question-paper";
import DiscriptiveExam from "../pages/discriptive-exam/discriptive-exam";
import SignedInRoute from "./signedin-route";
import PreviousSolutions from "../pages/previous-solutions/previous-solutions";
import UserProfile from "../pages/user-profle";

const routes: RouteObject[] = [
  {
    path: "/",
    element: <Navigate to={paths.HOMEPAGE} />,
  },
  {
    path: paths.HOMEPAGE,
    element: (
      <SignedInRoute>
        <HomePage />
      </SignedInRoute>
    ),
  },
  {
    path: paths.REGISTER,
    element: (
      <SignedInRoute>
        <RegisterPage />
      </SignedInRoute>
    ),
  },
  {
    path: paths.LOGIN,
    element: (
      <SignedInRoute>
        <Login />
      </SignedInRoute>
    ),
  },
  {
    path: paths.FORGOT_PASSWORD,
    element: <ForgotPassword />,
  },
  {
    path: paths.QUESTION_PAPER,
    element: (
      <ProtectedRoute>
        <QuestionPaper />
      </ProtectedRoute>
    ),
  },
  {
    path: paths.PROFILE,
    element: (
      <ProtectedRoute>
        <UserProfile />
      </ProtectedRoute>
    ),
  },
  {
    path: paths.EXAMS,
    children: [
      {
        path: `${paths.DISCRIPTIVE}/:questionId`,
        element: (
          <ProtectedRoute>
            <DiscriptiveExam />
          </ProtectedRoute>
        ),
      },
      {
        path: `${paths.PREVIOUS_SOLUTIONS}/:questionId`,
        element: (
          <ProtectedRoute>
            <PreviousSolutions />
          </ProtectedRoute>
        ),
      },
    ],
  },
  {
    path: paths.MY_SUMBISSIONS,
    element: (
      <ProtectedRoute>
        <div>Feature comming soon</div>
      </ProtectedRoute>
    ),
  },
];

export default routes;
