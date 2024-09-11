import { Button } from "../../componnets/base/button/button";
import Icon from "../../componnets/base/icon";
import Header from "../../componnets/shared/header/header";

export default function PricingPlan() {
  return (
    <>
      <div className=" bg-neutral-50">
        <Header />
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
              <div className="text-end">
                <Icon icon="arrow_right" className="text-6xl animate-pulse" />
              </div>
              <div className="text-lg leading-7">
                New plans and feature comming soon.
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
              <div className="mt-auto border border-dashed border-[#A9A9AA] tracking-widest mb-2" />
              <div className="h-28 ">
                <div className="flex flex-col gap-4 justify-between absolute left-6 right-6 bottom-6">
                  <div className="flex items-baseline">
                    <span className="text-4xl font-bold">
                      <Icon icon="rupee" className="text-3xl" />
                      {data.price}
                    </span>
                    <span>{data.duration}</span>
                  </div>
                  <div className="flex align-bottom">
                    <Button disabled={data.isDisables} className="w-full">
                      {data.buttonText}
                    </Button>
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </>
  );
}

const staticValue = [
  {
    passType: "Golden Pass",
    price: "19",
    duration: "/month",
    static: [
      "AI based descriptive exam assesments.",
      "24/7 support.",
      "Advanced analytics based on your answers.",
      "10 new descriptive question everyday.",
      "Track your progress with assessments history.",
    ],
    buttonText: "Choose",
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
    isDisables: true,
    buttonText: "Comming soon",
  },
];
