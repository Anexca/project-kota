import { Controller, useForm, useWatch } from "react-hook-form";
import AnswerOptions from "../../componnets/shared/answer-option/answer-options";
import { useState } from "react";
import QuestionSelector from "../../componnets/shared/question-selector/question-selector";
import { QUESTION_STATE } from "../../constants/shared";
import { MathJaxContext, MathJax } from "better-react-mathjax";
import { questionsList } from "./questionList";
import { Button } from "../../componnets/base/button/button";

type Props = {};

const ExamCenter = (props: Props) => {
  const [questions, setQuestions] = useState(questionsList);
  const {
    control,
    handleSubmit,
    register,
    setValue,
    formState: { errors },
  } = useForm({
    defaultValues: {
      answers: questions.map((_) => ({
        state: QUESTION_STATE.NOT_ATTEMPED,
        selectedOption: null,
      })),
      activeQuestionIndex: 0,
    },
  });
  const activeIndex = useWatch({ control, name: "activeQuestionIndex" });
  const answers = useWatch({ control, name: "answers" });

  const renderQuestion = (question: string) => {
    const parts = question.split(/(\\\(.*?\\\))/g);
    return parts.map((part, index) => {
      if (part.startsWith("\\(") && part.endsWith("\\)")) {
        // It's a LaTeX part, render it with KaTeX
        return <MathJax style={{ display: "inline" }}>{part}</MathJax>;
      } else {
        // It's plain text
        return part;
      }
    });
  };
  return (
    <MathJaxContext>
      <div className="max-h-screen h-screen flex flex-col">
        <div className="bg-gray-800 h-12 min-h-12"></div>
        <div className=" flex-1">
          <div className="h-full flex items-start">
            <div className="flex-1 h-full max-h-full flex flex-col justify-start bg-neutral-100 border-r shadow">
              <div className="bg-slate-100 text-start p-2">
                Que No {activeIndex + 1} / {questions.length}
              </div>
              <div className="flex-1 overflow-auto ">
                <div className="max-h-[70vh] h-full flex flex-col p-4">
                  <p className="items-start text-start pb-2">
                    {renderQuestion(questions[activeIndex].question)}
                    {/* <MathJax>{questions[activeIndex].question}</MathJax> */}
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
            <div className="bg-neutral-200 min-w-72 h-full px-2">
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
    </MathJaxContext>
  );
};

export default ExamCenter;
