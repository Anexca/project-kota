import { useMemo } from "react";
import { paths } from "../../../routes/route.constant";
import useUserProfileStore from "../../../store/user-info-store";
import { Button } from "../../base/button/button";
import Chip from "../../base/chip";
import Icon from "../../base/icon";
import { StyledLink } from "../../base/styled-link";
import Container from "../container";
import Footer from "../footer";
import Header from "../header/header";
import { AnimatedBeamHeading } from "./beam-heading";

const exams = [
  {
    title: "Community Banking Descriptive Exams",
    desc: "2 free Questions every week. AI-Assisted Descriptive Assessment with Insights. Get understanding about your weak and strong points.",
    link: `/${paths.COMMUNITY_EXAMS}/banking/${paths.DISCRIPTIVE}`,
    type: "Descriptive",
    isActive: true,
  },
  {
    title: "Banking Descriptive Exams",
    desc: "AI-Assisted Descriptive Assessment with Insights. Get understanding about your weak and strong points.",
    link: `/${paths.EXAMS}/banking/${paths.DISCRIPTIVE}`,
    type: "Descriptive",
    isActive: true,
    isPremium: true,
  },
  {
    title: "Banking MCQ Exams",
    desc: "Practice Banking MCQs to boost your knowledge and exam readiness with detailed solutions and explanations.",
    link: `/${paths.EXAMS}/banking/${paths.MCQ}`,
    type: "MCQ",
    isActive: false,
  },
  {
    title: "Ranked MCQ Exams",
    desc: "Compete in Ranked MCQ Exams and see how you rank among others while improving your skills.",
    link: `/${paths.EXAMS}/ranked/${paths.MCQ}`,
    type: "MCQ",
    isActive: false,
  },
];

const UserHomePage = () => {
  return (
    <div>
      <Header />
      <div className="relative bg-neutral-100/50" id="home">
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
              <div className="text-lg mb-2">Your Go-To Place for </div>
              <div className="text-2xl text-info font-semibold mb-4">
                Complete Exam Preparation
              </div>
              <div className="flex items-center justify-center mt-[10%]">
                <AnimatedBeamHeading />
              </div>

              <div className="mt-20 font-semibold">Take a test now.</div>
              <div>
                <Icon icon="arrow_right" className="text-info rotate-90" />
              </div>
            </div>
            <div className="w-full py-28">
              <Container>
                <h4 className="border-b py-2 font-bold text-center">
                  Available Exams
                </h4>
                <div className="p-2 px-0 my-4 space-y-4">
                  {exams.map((item, index) => (
                    <ExamCard {...item} srNumber={index + 1} />
                  ))}
                </div>
              </Container>
            </div>
          </div>
        </div>
      </div>

      <Footer />
    </div>
  );
};

export default UserHomePage;

const ExamCard = ({
  desc,
  link,
  title,
  srNumber,
  type,
  isActive,
  isPremium,
}: {
  title: string;
  desc: string;
  link: string;
  srNumber: number;
  type: string;
  isActive: boolean;
  isPremium?: boolean;
}) => {
  const { profile } = useUserProfileStore();
  const showAttempt = useMemo(() => {
    if (isPremium) {
      return !!profile.active_subscriptions;
    }
    return true;
  }, [profile]);
  return (
    <article className="rounded-md shadow-sm bg-white flex flex-col md:flex-row gap-4 p-3 px-4 md:p3 text-sm">
      <div className="flex-1">
        <p className="font-medium text-balance md:text-pretty text-black mb-2">
          #{srNumber} - {title}
        </p>
        <p className="text-balance md:text-pretty text-sm text-black mb-2">
          {desc}
        </p>
        <div className="flex gap-2 flex-wrap">
          <Chip icon="tags" variant={"info"}>
            {type}
          </Chip>
        </div>
      </div>
      <div className="flex flex-col items-stretch gap-2 md:justify-center md:w-32">
        {isActive ? (
          showAttempt ? (
            <StyledLink
              to={link}
              size={"sm"}
              className="px-3 py-1"
              variant={"info"}
            >
              <Icon icon="play_circle" className="mr-2" /> Attempt
            </StyledLink>
          ) : (
            <StyledLink
              to={`/${paths.PRICING_PLAN}`}
              size={"sm"}
              className="px-3 py-1"
              variant={"warning"}
            >
              <Icon icon="send" className="mr-2" /> Buy Plan
            </StyledLink>
          )
        ) : (
          <Button
            disabled
            size={"sm"}
            className="px-3 py-1"
            variant={"secondary"}
          >
            <Icon icon="clock" className="mr-2" /> Comming soon
          </Button>
        )}
      </div>
    </article>
  );
};
