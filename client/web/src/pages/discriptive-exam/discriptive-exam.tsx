import { useEffect, useState } from "react";
import { diffWords } from "diff"; // Importing the diff functions
import { modified, original } from "./testdata";
import { Button } from "../../componnets/base/button/button";
import { Textarea } from "../../componnets/base/text-area";
import { Controller, useForm, useWatch } from "react-hook-form";
import Container from "../../componnets/shared/container";
import { yupResolver } from "@hookform/resolvers/yup";
import { DiscriptiveExamSchema } from "../../validation-schema/discriptive-exam";
import { useLocation } from "react-router-dom";
import { IQuestion } from "../../interface/question";
import { useInterval } from "../../hooks/use-interval";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "../../componnets/base/dialog/dialog";
import TestHeader from "../../componnets/shared/test-header/test-header";
type Props = {};

// src/DiffChecker.js

const DiffChecker = () => {
  const [text1, setText1] = useState(original);
  const [text2, setText2] = useState(modified);
  const [diffResult, setDiffResult] = useState([]);

  const handleTextChange = (e) => {
    const { name, value } = e.target;
    if (name === "text1") {
      setText1(value);
    } else if (name === "text2") {
      setText2(value);
    }
  };

  const handleDiff = () => {
    // You can choose to use diffChars or diffLines based on your requirement
    const diff = diffWords(text1, text2);
    console.log(diff);

    setDiffResult(diff);
  };

  useEffect(() => {
    handleDiff();
  }, []);

  return (
    <div className="mt-4">
      <h2 className="mb-2">AI Assesed answer :</h2>
      <p
        style={{ whiteSpace: "pre-wrap" }}
        className="text-sm rounded p-2 border"

      >
        {diffResult.map((part, index) =>
          part.removed || part.added ? (
            <span
              key={index}
              style={{
                color: part.added ? "green" : part.removed ? "red" : "black",
                textDecoration: part.added
                  ? "underline"
                  : part.removed
                  ? "line-through"
                  : "none",
              }}
            >
              {part.value}
            </span>
          ) : (
            part.value
          )
        )}
      </p>
    </div>
  );
};
type LocationType = {
  state: {
    question: IQuestion;
  };
};

const ConformationDialog = ({ timerStart, time }: any) => {
  const [open, setOpen] = useState(true);
  return (
    <Dialog open={open} defaultOpen>
      <DialogContent noCloseButton>
        <DialogHeader>
          <DialogTitle>Exam Instruction.</DialogTitle>
          <DialogDescription>
            <div className="mb-2 text-neutral-700 font-semibold">
              The exam will automatically close when the time is up, so please
              manage your time carefully. You have{" "}
              <span className="text-orange-600">{time}</span> minutes to
              complete the exam.
            </div>
            <Button
              variant={"info"}
              onClick={() => {
                timerStart();
                setOpen(false);
              }}
              className="p-2 h-8"
            >
              Start Exam
            </Button>
          </DialogDescription>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
};

const DiscriptiveExam = (props: Props) => {
  const location: LocationType = useLocation();
  const [examTime, setExamTime] = useState(
    location.state.question.duration_minutes * 60
  );
  const interval = useInterval(() => setExamTime((s) => s - 1), 1000);
  const { control } = useForm({
    defaultValues: { answer: "" },
    resolver: yupResolver(DiscriptiveExamSchema),
  });

  const textLength = useWatch({ control, name: "answer" });

  const question = location.state.question;

  return (
    <div className="flex flex-col">
      <ConformationDialog
        timerStart={() => interval.start()}
        time={location.state.question.duration_minutes}
      />
      <TestHeader currentTime={examTime} active={interval.active} />
      <Container className="p-2">
        <div className="mb-4">
          <div className="text-sm font-medium">
            Que - {question.raw_exam_data.topic}
          </div>
          <span className="inline-flex items-center justify-center rounded-full bg-emerald-100 px-2.5 py-0.5 text-emerald-700">
            <i className="fa-regular fa-circle-check text-sm mr-2"></i>
            <p className="whitespace-nowrap text-sm capitalize">
              {question.raw_exam_data.type}
            </p>
          </span>
          <details className="mt-2 group overflow-hidden rounded [&_summary::-webkit-details-marker]:hidden">
            <summary className="flex cursor-pointer items-center justify-between gap-2 bg-white px-2 text-info transition">
              <span className="text-sm font-medium">Hints</span>

              <span className="transition group-open:-rotate-180">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  strokeWidth="1.5"
                  stroke="currentColor"
                  className="size-4"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M19.5 8.25l-7.5 7.5-7.5-7.5"
                  />
                </svg>
              </span>
            </summary>

            <div className="border border-gray-200 rounded-lg bg-white p-4">
              <ul className="space-y-1 border-gray-200 ml-4">
                {question.raw_exam_data.hints.map((hint) => (
                  <li className=" text-sm list-disc">{hint}</li>
                ))}
              </ul>
            </div>
          </details>
        </div>
        <form>
          <div className="w-full mb-4 border border-gray-200 rounded-lg bg-gray-50 dark:bg-gray-700 dark:border-gray-600">
            <div className="px-2 py-2 bg-white rounded-t-lg dark:bg-gray-800">
              <label htmlFor="comment" className="sr-only">
                Your comment
              </label>
              <Controller
                name="answer"
                control={control}
                render={({ field }) => {
                  return <Textarea {...field} className="text-sm" rows={10} />;
                }}
              />
            </div>
            <div className="flex items-center justify-between px-3 py-2 border-t dark:border-gray-600">
              <Button variant={"info"} className="p-2 h-8">
                Submit Answer
              </Button>
              <div className="text-sm">{textLength.length}/250</div>
            </div>
          </div>
        </form>

        <p className="ms-auto text-xs font-semibold text-gray-500 dark:text-gray-400">
          Remember, you have only{" "}
          {location.state.question.max_attempts -
            location.state.question.user_attempts}{" "}
          attempts left for this exam.
        </p>

        <DiffChecker />
      </Container>
    </div>
  );
};

export default DiscriptiveExam;
