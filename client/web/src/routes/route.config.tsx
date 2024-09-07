import { Navigate, RouteObject } from "react-router-dom";
import { ForgotPassword } from "../pages/forgot-password/forgot-password";
import HomePage from "../pages/homepage/homepage";
import { Login } from "../pages/login/login";
import { RegisterPage } from "../pages/register/register";
import ProtectedRoute from "./protected-route";
import { paths } from "./route.constant";

import GeneralLayout from "../layouts/general-layout";
import DescriptiveQuestion from "../pages/descriptive-questions";
import DescriptiveSubmission from "../pages/descriptive-submission";
import DiscriptiveExam from "../pages/discriptive-exam/discriptive-exam";
import UserProfile from "../pages/user-profle";
import SignedInRoute from "./signedin-route";

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
    path: paths.PROFILE,
    element: (
      <ProtectedRoute>
        <UserProfile />
      </ProtectedRoute>
    ),
  },
  {
    path: paths.EXAMS,
    element: (
      <ProtectedRoute>
        <GeneralLayout />
      </ProtectedRoute>
    ),
    children: [
      {
        path: `banking/${paths.DISCRIPTIVE}`,
        element: <DescriptiveQuestion />,
      },
      {
        path: `banking/${paths.DISCRIPTIVE}/:questionId/${paths.SUBMISSION}/:assesmentId`,
        element: <DescriptiveSubmission />,
      },
      {
        path: `${paths.MY_SUMBISSIONS}/:examId`,
        element: (
          <ProtectedRoute>
            <div>Comming soon</div>
          </ProtectedRoute>
        ),
      },
    ],
  },
  {
    path: `${paths.EXAMS}/banking/${paths.DISCRIPTIVE}/:questionId`,
    element: (
      <ProtectedRoute>
        <DiscriptiveExam />
      </ProtectedRoute>
    ),
  },
];

export default routes;
