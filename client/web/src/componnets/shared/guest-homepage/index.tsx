import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { ScreenSizeQuery } from "../../../constants/shared";
import { useMediaQuery } from "../../../hooks/use-media-query";
import { useToast } from "../../../hooks/use-toast";
import { paths } from "../../../routes/route.constant";

import { IExamCategories } from "../../../interface/exam-categories";
import { getQuestionsCategories } from "../../../services/exam-categories.services";
import Chip from "../../base/chip";
import { ToastAction } from "../../base/toast";
import Footer from "../footer";
import Header from "../header/header";
import StrippedLoader from "../stripped-loader";

import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "../accordian";

export const Card = ({
  title,
  description,
  icon,
  chip,
}: {
  title: string;
  description: string;
  icon: string;
  chip: JSX.Element;
}) => (
  <article className="relative block overflow-hidden h-full rounded-lg border border-gray-100 p-4 sm:p-6 lg:p-8">
    <span className="absolute inset-x-0 bottom-0 h-2 bg-gradient-to-r from-green-300 via-blue-500 to-purple-600"></span>

    <div className="sm:flex sm:justify-between sm:gap-4">
      <div>
        <h3 className="text-lg font-bold text-gray-900 sm:text-xl">{title}</h3>
      </div>

      <div className="hidden sm:block sm:shrink-0">
        <i className={icon}></i>
      </div>
    </div>

    <div className="mt-4">
      <p className="text-pretty text-sm text-gray-500">{description}</p>
    </div>

    <dl className="mt-6 flex gap-4 sm:gap-6">
      <div className="flex flex-col-reverse">
        <dt className="text-sm font-medium text-gray-600">{chip}</dt>
      </div>
    </dl>
  </article>
);

const featureWeOffer = [
  {
    title: "Descriptive Answer Test with AI Insights",
    description:
      "Submit descriptive answers and receive detailed insights, feedback, and assessments using our AI-powered analysis system.",
    icon: "fas fa-robot",
  },
  {
    title: "AI-Powered Test Analysis",
    description:
      "Get personalized insights and recommendations based on your performance, highlighting areas of improvement.",
    icon: "fas fa-chart-pie",
  },

  {
    title: "Daily New Questions",
    description:
      "Access fresh questions every day to keep your preparation up-to-date and challenging.",
    icon: "fas fa-calendar-alt",
  },
  {
    title: "View Previous Submissions",
    description:
      "Review your past test submissions and track your progress over time with ease.",
    icon: "fas fa-history",
  },
];

const upcomingFeatures = [
  {
    title: "Live Mock Test Challenges",
    description:
      "Compete with peers in real-time mock tests and see where you stand with instant rankings.",
    icon: "fas fa-trophy",
  },
  {
    title: "Detailed Performance Reports",
    description:
      "Receive in-depth performance reports with topic-wise breakdowns to track your progress and focus on areas of improvement.",
    icon: "fas fa-chart-line",
  },

  {
    title: "Mobile App",
    description:
      "Take your preparation on the go with our upcoming mobile app—practice anytime, anywhere!",
    icon: "fas fa-mobile-alt",
  },
  {
    title: "Roadmap for Exam Preparation",
    description:
      "Follow structured roadmaps that guide you through each step of your exam preparation journey.",
    icon: "fas fa-map",
  },
];

const GuestHomePage = () => {
  const [categories, setCategories] = useState<IExamCategories[]>([]);
  const isDesktopMode = useMediaQuery(ScreenSizeQuery.largeScreen, true);
  const [loading, setLoading] = useState(false);
  const { toast } = useToast();
  const navigate = useNavigate();
  const getQuestionsList = async () => {
    setLoading(true);
    try {
      const data = await getQuestionsCategories();
      setCategories(data.data);
    } catch (error) {
      toast({
        title: "Oh ho Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in processing your request.",
      });
    }
    setLoading(false);
  };
  const showNotification = () => {
    toast({
      title: "Hold On!",
      variant: "info",
      description:
        "You'll need to sign in before taking the exam. Let's get you set up!",
      action: (
        <ToastAction
          altText="Login"
          onClick={() => navigate(`/${paths.LOGIN}`)}
        >
          Login
        </ToastAction>
      ),
    });
  };
  useEffect(() => {
    getQuestionsList();
  }, []);
  const categoryList = isDesktopMode ? categories : categories.slice(0, 6);
  return (
    <>
      <Header />
      <div className="relative" id="home">
        <div
          aria-hidden="true"
          className="absolute inset-0 grid grid-cols-2 -space-x-52 opacity-40 dark:opacity-20"
        >
          <div className="blur-[106px] h-56 bg-gradient-to-br from-primary to-purple-400 dark:from-blue-700"></div>
          <div className="blur-[106px] h-32 bg-gradient-to-r from-cyan-400 to-sky-300 dark:to-indigo-600"></div>
        </div>
        <div className="max-w-7xl mx-auto px-6 md:px-12 xl:px-6">
          <div className="relative ml-auto pt-20">
            <div className="lg:w-2/3 text-center mx-auto ">
              <h1 className="text-balance text-gray-900 dark:text-white font-bold text-5xl md:text-6xl xl:text-7xl">
                Ace Your Exams with{" "}
                <span className="text-purple-500 dark:text-white">
                  AI-Powered Mock
                </span>{" "}
                Tests
              </h1>
              <p className="mt-8 text-gray-700 dark:text-gray-300">
                Experience the next generation of mock tests designed to give
                you an edge in your exam preparation. Our platform offers
                in-depth practice for various goverment and entrance exam with
                an exclusive feature of AI-evaluated descriptive answers.
              </p>
              <div className="mt-16 flex flex-wrap justify-center gap-y-4 gap-x-6">
                <Link
                  to={`/${paths.REGISTER}`}
                  className="relative flex h-11 w-full items-center justify-center px-6 before:absolute before:inset-0 before:rounded-full before:bg-primary before:transition before:duration-300 hover:before:scale-105 active:duration-75 active:before:scale-95 sm:w-max"
                >
                  <span className="relative text-base font-semibold text-white">
                    Register to Get started
                  </span>
                </Link>
                <Link
                  to={`/${paths.LOGIN}`}
                  className="relative flex h-11 w-full items-center justify-center px-6 before:absolute before:inset-0 before:rounded-full before:border before:border-transparent before:bg-primary/10 before:bg-gradient-to-b before:transition before:duration-300 hover:before:scale-105 active:duration-75 active:before:scale-95 dark:before:border-gray-700 dark:before:bg-gray-800 sm:w-max"
                >
                  <span className="relative text-base font-semibold text-primary dark:text-white">
                    Log In
                  </span>
                </Link>
              </div>
              <div className="hidden py-8 mt-16 border-y border-gray-100 dark:border-gray-800 sm:flex justify-between  gap-2">
                <div className="text-left flex-1">
                  <h6 className="text-lg font-semibold text-gray-700 dark:text-white">
                    Lowest price
                  </h6>
                  <p className="mt-2 text-gray-500">
                    Get top-tier mock tests without breaking the bank.
                  </p>
                </div>
                <div className="text-left flex-1">
                  <h6 className="text-lg font-semibold text-gray-700 dark:text-white">
                    Fastest on the market
                  </h6>
                  <p className="mt-2 text-gray-500">
                    Experience lightning-fast results and feedback.
                  </p>
                </div>
                <div className="text-left flex-1">
                  <h6 className="text-lg font-semibold text-gray-700 dark:text-white">
                    Most loved
                  </h6>
                  <p className="mt-2 text-gray-500">
                    Our platform delivers AI-powered evaluations that students
                    rave about.
                  </p>
                </div>
              </div>
            </div>

            <>
              <div className="text-center my-4 font-bold text-2xl underline mt-14">
                Exams We Offer
              </div>
              <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4">
                {loading
                  ? "*"
                      .repeat(10)
                      .split("")
                      .map(() => (
                        <div className="p-4">
                          <StrippedLoader className="w-full h-20" />
                        </div>
                      ))
                  : categoryList.map((item) => (
                      <div className="p-4 ">
                        <article
                          onClick={showNotification}
                          role="button"
                          className=" cursor-pointer overflow-hidden rounded-lg border border-gray-100 p-4 flex gap-3 items-center"
                        >
                          <div className="h-10 aspect-square">
                            <img
                              src={item.logo_url}
                              className="object-contain "
                            />
                          </div>
                          <h3 className="text-lg font-bold text-gray-900 sm:text-xl">
                            {item.exam_group_name.split("_").join(" ")}
                          </h3>
                        </article>
                      </div>
                    ))}
              </div>
              {!isDesktopMode && (
                <div className="w-full font-medium text-center animate-pulse pb-8">
                  And Many More Exams
                </div>
              )}
            </>

            <div className="text-center my-4 font-bold text-2xl underline mt-14">
              Feature We Provide
            </div>
            <div className=" max-w-3xl mx-auto">
              <Accordion type="single" collapsible className="w-full">
                {featureWeOffer.map((item, index) => (
                  <AccordionItem value={`item-${index}`}>
                    <AccordionTrigger>
                      <span>
                        <i className={item.icon}></i>
                        <span className="mx-4">{item.title}</span>
                        <Chip variant={"success"} icon="play_circle">
                          Live
                        </Chip>
                      </span>
                    </AccordionTrigger>
                    <AccordionContent>{item.description}</AccordionContent>
                  </AccordionItem>
                ))}
                {upcomingFeatures.map((item, index) => (
                  <AccordionItem value={`coming-soon-${index}`}>
                    <AccordionTrigger>
                      <span>
                        <i className={item.icon}></i>
                        <span className="mx-4">{item.title}</span>
                        <Chip icon="clock">Coming Soon</Chip>
                      </span>
                    </AccordionTrigger>
                    <AccordionContent>{item.description}</AccordionContent>
                  </AccordionItem>
                ))}
              </Accordion>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default GuestHomePage;
