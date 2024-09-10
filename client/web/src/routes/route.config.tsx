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
import NotFound from "../pages/not-found";
import PricingPlan from "../pages/pricing-plan";
import TermsOfService from "../pages/terms-of-service";
import PrivacyPolicy from "../pages/privacy-policy";
import ContactUs from "../pages/contact-us";

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
        path: `${paths.MY_SUMBISSIONS}`,
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
  { path: "/pricing-plan", element: <PricingPlan /> },
  { path: "/terms-of-service", element: <TermsOfService /> },
  { path: "/privacy-policy", element: <PrivacyPolicy /> },
  { path: "/contact-us", element: <ContactUs /> },
  { path: "*", element: <NotFound /> },
];

export default routes;
