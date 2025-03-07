import { useEffect, useState } from "react";

import { Link, useNavigate } from "react-router-dom";
import { ExamDomain, paths } from "../../routes/route.constant";

import Icon from "../../componnets/base/icon";
import NoPremiumBanner from "../../componnets/shared/no-premium-banner";
import { useToast } from "../../hooks/use-toast";

import { Button } from "../../componnets/base/button/button";
import { StyledLink } from "../../componnets/base/styled-link";
import { IExamCategories } from "../../interface/exam-categories";
import { getQuestionsCategories } from "../../services/exam-categories.services";
import useUserProfileStore from "../../store/user-info-store";

const QuestionCategories = ({ isOpenMode }: { isOpenMode?: boolean }) => {
  const [categories, setCategories] = useState<IExamCategories[]>([]);
  const { profile } = useUserProfileStore();
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const { toast } = useToast();
  const getQuestionsList = async () => {
    setLoading(true);
    try {
      const data = await getQuestionsCategories(isOpenMode);
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
  useEffect(() => {
    getQuestionsList();
  }, []);

  const getQuestions = (index: number) => {
    const path = isOpenMode
      ? `/${paths.COMMUNITY_EXAMS}/${ExamDomain.banking}/${index}`
      : `/${paths.EXAMS}/${ExamDomain.banking}/${index}`;
    navigate(path);
  };

  return (
    <div className="pt-2 w-full md:max-w-2xl 2xl:max-w-2xl mx-auto flex flex-col gap-2 p-4">
      <div className="py-2">
        <div className="flex gap-2 items-center">
          <Link to={`/${paths.HOMEPAGE}`} className="p-0">
            <Icon icon="arrow_back" className="text-info text-lg" />
          </Link>
          <span className="text-lg font-semibold">Banking Exams</span>
        </div>
        <div className="text-sm text-black font-medium mb-2">
          Get started with the exam below.
        </div>
      </div>
      {loading ? (
        <div className="flex flex-col gap-2 justify-center items-center">
          <span className="rounded-full w-8 h-8 animate-spin border-2 border-info border-t-info/30"></span>
          Getting Questions Categories
        </div>
      ) : (
        <div className="animate-fadeIn flex flex-col gap-2">
          {categories.map((item) => {
            return (
              <article className="rounded-md shadow-sm bg-white flex flex-col md:flex-row gap-4 p-3 px-4 md:p3 text-sm">
                <div className="flex flex-1">
                  <div className="mr-2">
                    <img
                      className="h-10 aspect-square object-contain"
                      src={item.logo_url}
                    />
                  </div>
                  <div className="flex-1">
                    <p className="font-semibold text-balance md:text-pretty text-black mb-1">
                      {item.exam_group_name.split("_").join(" ")}
                    </p>
                    <p className="font-medium text-balance md:text-pretty text-black mb-2">
                      {item.description}
                    </p>
                  </div>
                </div>
                <div className="flex flex-col items-stretch gap-2 md:justify-center md:w-32">
                  {profile?.active_subscriptions?.length ? (
                    <Button
                      onClick={() => getQuestions(item.exam_group_id)}
                      disabled={!item.is_active}
                      size={"sm"}
                      className="px-3 py-1"
                      variant={item.is_active ? "info" : "outline"}
                    >
                      {item.is_active ? (
                        <>
                          <Icon icon="play_circle" className="mr-2" /> Attempt
                        </>
                      ) : (
                        <>
                          <Icon icon="clock" className="mr-2" /> Coming soon
                        </>
                      )}
                    </Button>
                  ) : (
                    <StyledLink
                      to={`/${paths.PRICING_PLAN}`}
                      size={"sm"}
                      className="px-3 py-1"
                      variant={"warning"}
                    >
                      <Icon icon="send" className="mr-2" /> Buy Plan
                    </StyledLink>
                  )}
                </div>
              </article>
            );
          })}
          {isOpenMode ? <NoPremiumBanner /> : null}
        </div>
      )}
    </div>
  );
};

export default QuestionCategories;
