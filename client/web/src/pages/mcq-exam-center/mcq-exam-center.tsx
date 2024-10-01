import { MathJax, MathJaxContext } from "better-react-mathjax";
import { useEffect, useState } from "react";
import { Controller, useForm, useWatch } from "react-hook-form";
import { useParams } from "react-router-dom";
import { Button } from "../../componnets/base/button/button";
import AnswerOptions from "../../componnets/shared/answer-option/answer-options";
import QuestionSelector from "../../componnets/shared/question-selector/question-selector";
import TestHeader from "../../componnets/shared/test-header/test-header";
import { QUESTION_STATE } from "../../constants/shared";
import { useInterval } from "../../hooks/use-interval";
import { useToast } from "../../hooks/use-toast";
import { IMCQQuestionSet } from "../../interface/question";
import { getMCQQuestionById } from "../../services/exam.service";

const MCQExamCenter = () => {
  const param = useParams();

  const [questionSet, setQuestionSet] = useState<IMCQQuestionSet | null>(null);
  const [examTime, setExamTime] = useState(0);
  const interval = useInterval(() => setExamTime((s) => s - 1), 1000);
  const { toast } = useToast();

  const { control, setValue, reset } = useForm({
    defaultValues: {
      answers: [] as any,
      activeQuestionIndex: 0,
    },
  });
  const activeIndex = useWatch({ control, name: "activeQuestionIndex" });
  const answers = useWatch({ control, name: "answers" });

  const fetchExam = async () => {
    try {
      const res = await getMCQQuestionById(param.questionId!);
      setExamTime(res.data.duration_seconds);
      setQuestionSet(res.data);
      const ques = res.data.raw_exam_data.exam_content.map((_) => ({
        state: QUESTION_STATE.NOT_ATTEMPED,
        selectedOption: null,
      }));
      reset({ activeQuestionIndex: 0, answers: ques });
    } catch (error) {
      toast({
        title: "Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in processing your request.",
      });
    }
  };

  useEffect(() => {
    fetchExam();
  }, []);

  const questions = questionSet?.raw_exam_data?.exam_content || [];

  return (
    <MathJaxContext>
      <div className="max-h-screen h-screen flex flex-col">
        <TestHeader currentTime={examTime} active={interval.active} />
        <div className="flex-1">
          <div className="h-full flex flex-col md:flex-row items-start">
            <div className="pt-2 flex-1 h-full max-h-full flex flex-col justify-start bg-neutral-50/50 border-r-0 md:border-r ">
              <div className="text-start p-2 pl-4 ">
                Que No {activeIndex + 1} /{" "}
                {questionSet?.number_of_questions || 0}
              </div>
              <div className="flex-1 overflow-auto flex md:flex-row flex-col">
                {questions?.[activeIndex]?.content && (
                  <div className="md:w-1/2 min-w-[50%] p-4 pt-0 md:pt-4  text-pretty font-medium border-r-0 border-b md:border-b-0 md:border-r">
                    <p className="font-semibold mb-2">
                      {questions[activeIndex].note}
                    </p>
                    <p>{questions[activeIndex].content}</p>
                  </div>
                )}
                <div className="md:max-h-[70vh] h-full flex flex-col p-4">
                  {questions.length && (
                    <>
                      <p className="items-start text-start pb-2">
                        <MarkdownRender
                          children={questions?.[activeIndex]?.question || ""}
                        />
                      </p>

                      <Controller
                        name={`answers.${activeIndex}.selectedOption`}
                        control={control}
                        render={({ field }) => (
                          <AnswerOptions
                            name={field.name}
                            onChange={field.onChange}
                            options={questions[activeIndex].options}
                            selected={field.value}
                          />
                        )}
                      />
                    </>
                  )}
                </div>
              </div>
              <div className="min-h-12 flex gap-2 p-2 border-orange-300 border-t">
                <Button size={"sm"}>Mark for Review & Next</Button>
                <Button variant={"outline"} size={"sm"} className="mr-auto">
                  Clear
                </Button>
                <Controller
                  name={`activeQuestionIndex`}
                  control={control}
                  render={({ field }) => (
                    <Button
                      size={"sm"}
                      variant={"success"}
                      className="justify-self-end"
                      type="button"
                      disabled={field.value + 1 >= questions.length}
                      onClick={() => field.onChange(field.value + 1)}
                    >
                      SAVE & NEXT
                    </Button>
                  )}
                />
              </div>
            </div>
            <div className="py-2 bg-neutral-100/75 min-w-72 md:w-auto w-full h-full px-2 ">
              <div className=" min-w-72">
                <QuestionSelector
                  questions={answers}
                  onQuestionNumberClick={(index) =>
                    setValue("activeQuestionIndex", index)
                  }
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </MathJaxContext>
  );
};

export default MCQExamCenter;

import ReactMarkdown, { Components } from "react-markdown";

import RemarkMathPlugin from "remark-math";
import remarkGfm from "remark-gfm";
import m from "./m.module.scss";
function MarkdownRender({ ...props }) {
  const cp: Components = {
    code: (props) => {
      return (
        <MathJax style={{ display: "inline" }}>\({props.children}\)</MathJax>
      );
    },
    table: (props) => {
      return <span className={m.markdown}>{props.children}</span>;
    },
  };
  const newProps = {
    ...props,
    remarkPlugins: [RemarkMathPlugin, remarkGfm],
    components: cp,
  };
  return (
    <MathJaxContext>
      <ReactMarkdown {...newProps} />
    </MathJaxContext>
  );
}
