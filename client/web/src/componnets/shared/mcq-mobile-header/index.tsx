import clsx from "clsx";
import { useState } from "react";

import MCQQuestionPallet from "../mcq-question-pallet";
import { QuestionSelectorProps } from "../question-selector/question-selector";

const MCQMobileHeader = ({
  answers,
  activeQuestionIndex,
  onQuestionNumberClick,
  handleSubmit,
  sectionList,
  activeSection,
}: QuestionSelectorProps & {
  handleSubmit?: () => void;
  sectionList: string[];
}) => {
  const [mobileViewHeader, setMobileViewHeader] = useState(false);

  return (
    <div>
      <div className="relative flex w-full justify-between md:px-0 lg:w-max">
        <div className="relative flex max-h-10 items-center lg:hidden">
          <button
            onClick={() => setMobileViewHeader(!mobileViewHeader)}
            aria-label="humburger"
            id="hamburger"
            className="relative p-2 md:p-6"
          >
            <div
              aria-hidden="true"
              id="line"
              className={clsx(
                "m-auto h-0.5 w-5 rounded bg-sky-900 transition duration-300 dark:bg-gray-300",
                mobileViewHeader && "translate-y-1 rotate-45"
              )}
            ></div>
            <div
              aria-hidden="true"
              id="line2"
              className={clsx(
                "m-auto mt-2 h-0.5 w-5 rounded bg-sky-900 transition duration-300 dark:bg-gray-300",
                mobileViewHeader && "-translate-y-1.5 -rotate-45"
              )}
            ></div>
          </button>
        </div>
      </div>

      <div className="z-20">
        <div
          id="navLayer"
          aria-hidden="true"
          className={clsx(
            "fixed inset-0 z-10 h-screen  origin-bottom scale-y-0 bg-white/70 backdrop-blur-2xl transition duration-500 dark:bg-gray-900/70 lg:hidden",
            mobileViewHeader && "origin-top scale-y-100"
          )}
        ></div>

        <div
          id="navLayer"
          aria-hidden="true"
          className={clsx(
            "fixed invisible inset-0 z-10  flex overflow-hidden h-screen flex-col  origin-bottom scale-90 opacity-0 bg-white/70 backdrop-blur-2xl transition duration-500 dark:bg-gray-900/70 lg:hidden",
            mobileViewHeader && "origin-top  !visible !scale-100 !opacity-100"
          )}
        >
          <div className=" p-6 py-4">
            <button
              onClick={() => setMobileViewHeader(!mobileViewHeader)}
              aria-label="humburger"
              id="hamburger"
              className="relative"
            >
              <div
                aria-hidden="true"
                id="line"
                className={clsx(
                  "m-auto h-0.5 w-5 rounded bg-sky-900 transition duration-300 dark:bg-gray-300",
                  mobileViewHeader && "translate-y-1 rotate-45"
                )}
              ></div>
              <div
                aria-hidden="true"
                id="line2"
                className={clsx(
                  "m-auto mt-2 h-0.5 w-5 rounded bg-sky-900 transition duration-300 dark:bg-gray-300",
                  mobileViewHeader && "-translate-y-1.5 -rotate-45"
                )}
              ></div>
            </button>
          </div>
          <div className=" flex justify-center   gap-2 w-full flex-1 overflow-hidden">
            <MCQQuestionPallet
              activeSection={activeSection}
              answers={answers}
              onQuestionNumberClick={onQuestionNumberClick}
              activeQuestionIndex={activeQuestionIndex}
              handleSubmit={handleSubmit}
              sectionList={sectionList}
            />
          </div>
        </div>
      </div>
    </div>
  );
};
export default MCQMobileHeader;
