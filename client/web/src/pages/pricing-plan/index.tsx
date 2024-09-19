import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import Icon from "../../componnets/base/icon";
import Header from "../../componnets/shared/header/header";
import RazorpayButton from "../../componnets/shared/razorpay-button";
import { paths } from "../../routes/route.constant";
import { getPlans } from "../../services/payment.service";
import useSessionStore from "../../store/auth-store";
import useUserProfileStore from "../../store/user-info-store";
import Chip from "../../componnets/base/chip";
import Loader from "../../componnets/shared/loder";
import { useToast } from "../../hooks/use-toast";

export default function PricingPlan() {
  const { session } = useSessionStore();
  const { toast } = useToast();
  const [plans, setPlans] = useState<any>([]);
  const [loading, setIsloading] = useState(true);
  const { profile } = useUserProfileStore();
  const getPricingPlan = async () => {
    try {
      const res = await getPlans();
      setPlans(res.data);
    } catch (error) {
      toast({
        title: "Something went wrong",
        description:
          "Sorry we are not able to get your profile at this moment.",
        variant: "destructive",
        duration: 5000,
      });
    } finally {
      setIsloading(false);
    }
  };

  useEffect(() => {
    getPricingPlan();
  }, []);
  return (
    <>
      <div className=" bg-neutral-50">
        <Header />

        {loading ? (
          <div className="h-[80vh] flex items-center justify-center">
            <Loader size={"large"} color={"info"} />
          </div>
        ) : (
          <div className="font-sans flex flex-col lg:flex-row justify-center px-5 md:px-4 py-8 lg:py-4 w-full gap-6 items-center lg:items-stretch">
            {/* first portion */}
            <div className="flex flex-col flex-wrap max-w-[360px] md:w-[384px] p-6 bg-info group rounded-2xl relative overflow-hidden">
              <div className="text-start text-white flex flex-col gap-1">
                <span className="font-bold text-xl underline">
                  Select a plan and secure Future
                </span>
                <br />
                <span className="font-bold text-3xl">With Preparation</span>
                <br />
                <div className="text-lg leading-7">
                  Choose a plan and get started in minutes. Begin your journey
                  toward excellence today!
                </div>
                <div className="text-lg leading-7 lg:mt-4">
                  New plans and feature comming soon.
                </div>
                <div className="text-end">
                  <Icon
                    icon="arrow_right"
                    className="text-6xl animate-pulse rotate-90 lg:rotate-0"
                  />
                </div>
              </div>
            </div>
            {/* middle portion  */}
            {staticValue.map((data, index) => (
              <div
                key={index}
                className="flex flex-col max-w-[360px] md:w-[384px] p-6 py-4 shadow bg-white group rounded-2xl border xl:border-none border-[#0B0641] relative"
              >
                <div className="flex flex-row gap-5 items-center">
                  <span className="text-2xl font-bold">{data.passType}</span>
                  {!data.isDisabled &&
                    profile?.active_subscriptions?.[0]
                      ?.provider_subscription_id && (
                      <Chip icon="check_solid" variant={"success"}>
                        Active
                      </Chip>
                    )}
                </div>
                <span className="flex mt-4 text-[#A9A9AA] text-base">
                  What You&apos;ll Get
                </span>
                <div className="pb-3">
                  {data.static.map((myData, index) => (
                    <div
                      key={index}
                      className="flex flex-row gap-3 items-start mt-6 text-left text-sm"
                    >
                      <div className=" shrink-0 ">
                        <Icon icon="check_solid" />
                      </div>
                      <span>{myData}</span>
                    </div>
                  ))}
                </div>
                <div>
                  {!data.isDisabled &&
                    profile?.active_subscriptions?.[0]
                      ?.provider_subscription_id && (
                      <span className="text-sm flex gap-2 font-bold">
                        Valid upto
                        <Chip icon="clock" className="mb-2">
                          {new Date(
                            profile.active_subscriptions[0].end_date
                          ).toLocaleDateString()}
                        </Chip>
                      </span>
                    )}
                </div>
                <div className="mt-auto border border-dashed border-[#A9A9AA] tracking-widest mb-2" />
                <div className="">
                  <div className="flex flex-col gap-4 justify-between ">
                    <div className="flex items-baseline flex-row">
                      {data.original && (
                        <div className="flex items-baseline line-through text-destructive mr-2">
                          <span className="text-2xl font-bold">
                            <Icon icon="rupee" className="text-xl" />
                            {data.original}
                          </span>
                        </div>
                      )}
                      <div className="flex items-baseline">
                        <span className="text-4xl font-bold">
                          <Icon icon="rupee" className="text-3xl" />
                          {data.price}
                        </span>
                        <span>{data.duration}</span>
                      </div>
                    </div>

                    <div className="flex align-bottom">
                      {session ? (
                        <RazorpayButton
                          subscriptionId={plans[0]?.provider_plan_id}
                          id={plans[0]?.id}
                          isDisabled={data.isDisabled}
                        >
                          {data.buttonText}
                        </RazorpayButton>
                      ) : (
                        <Link
                          to={`/${paths.LOGIN}`}
                          className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 px-4 py-2 bg-info text-info-foreground hover:bg-info/90 w-full"
                        >
                          Login to buy
                        </Link>
                      )}
                    </div>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </>
  );
}

const staticValue = [
  {
    passType: "Golden Pass",
    price: "19",
    original: "59",
    duration: "/month",
    static: [
      "AI based descriptive exam assesments.",
      "24/7 support.",
      "Advanced analytics based on your answers.",
      "10 new descriptive question everyday.",
      "Track your progress with assessments history.",
    ],
    buttonText: "Buy Plan",
  },
  {
    passType: "Platinum Pass",
    price: "??",
    duration: "/month",
    static: [
      "Everything in Golden Pass",
      "Advance MCQ",
      "Advance analytics as per attempted exams.",
      "Tracking of attempted exams.",
    ],
    isDisabled: true,
    buttonText: "Comming soon",
  },
];
