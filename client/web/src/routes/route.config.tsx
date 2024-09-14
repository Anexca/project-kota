import { Navigate, Outlet, RouteObject } from "react-router-dom";
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
    element: <HomePage />,
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
    path: "",
    element: (
      <ProtectedRoute>
        <Outlet />
      </ProtectedRoute>
    ),
    children: [
      {
        path: paths.PROFILE,
        element: (
          <GeneralLayout>
            <UserProfile />
          </GeneralLayout>
        ),
      },
      {
        path: paths.EXAMS,
        element: <GeneralLayout />,
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
            element: <div>Comming soon</div>,
          },
        ],
      },
      {
        path: `${paths.EXAMS}/banking/${paths.DISCRIPTIVE}/:questionId`,
        element: <DiscriptiveExam />,
      },
    ],
  },
  { path: paths.PRICING_PLAN, element: <PricingPlan /> },
  { path: "/terms-of-service", element: <TermsOfService /> },
  { path: "/privacy-policy", element: <PrivacyPolicy /> },
  { path: "/contact-us", element: <ContactUs /> },
  { path: "*", element: <NotFound /> },
];

export default routes;
