// import { MathJaxContext } from "better-react-mathjax";
// import { useEffect, useMemo, useState } from "react";
// import { Controller, useForm, useWatch } from "react-hook-form";
// import { useNavigate, useParams } from "react-router-dom";
// import { Button } from "../../componnets/base/button/button";
// import AnswerOptions from "../../componnets/shared/answer-option/answer-options";
// import { ReadMore } from "../../componnets/shared/read-more-content";
// import TestHeader from "../../componnets/shared/test-header/test-header";
// import { QUESTION_STATE, ScreenSizeQuery } from "../../constants/shared";
// import { useInterval } from "../../hooks/use-interval";
// import { useMediaQuery } from "../../hooks/use-media-query";
// import { useToast } from "../../hooks/use-toast";
// import { MCQFormModal } from "../../interface/mcq-exam";
// import { IContentGroup, IMCQQuestionSet } from "../../interface/question";
// import { paths } from "../../routes/route.constant";
// import { getMCQQuestionById } from "../../services/exam.service";
// import { evaluateMcqExam } from "../../services/mcq-exam.services";

// import MarkdownRender from "../../componnets/shared/markdown-rendere";
// import MCQMobileHeader from "../../componnets/shared/mcq-mobile-header";
// import MCQQuestionPallet from "../../componnets/shared/mcq-question-pallet";

// const MCQSubmission = ({ isOpenMode }: { isOpenMode?: boolean }) => {
//   const param = useParams();
//   const [questionSet, setQuestionSet] = useState<IMCQQuestionSet | null>(null);
//   const [examTime, setExamTime] = useState(0);
//   const interval = useInterval(() => setExamTime((s) => s - 1), 1000);
//   const { toast } = useToast();
//   const isScreenView = useMediaQuery(ScreenSizeQuery.largeScreen, true);
//   const isMobileView = useMediaQuery(ScreenSizeQuery.mediumScreen, true);

//   const navigate = useNavigate();

//   const { control, setValue, reset } = useForm<
//     MCQFormModal & { tempAnswerIndex: number | null }
//   >({
//     defaultValues: {
//       answers: [],
//       activeQuestionIndex: 0,
//       tempAnswerIndex: null,
//     },
//   });
//   const activeIndex = useWatch({ control, name: "activeQuestionIndex" });
//   const answers = useWatch({ control, name: "answers" });
//   const tempAnswerIndex = useWatch({ control, name: "tempAnswerIndex" });

//   const fetchExam = async () => {
//     try {
//       const res = await getMCQQuestionById(param.questionId!);
//       // setExamTime(5);
//       setExamTime(res.data.duration_seconds);
//       setQuestionSet(res.data);
//       const ques = res.data.raw_exam_data.questions.map((_) => ({
//         state: QUESTION_STATE.UN_ATTEMPTED,
//         selectedOption: null,
//       }));
//       reset({ activeQuestionIndex: 0, answers: ques });
//     } catch (error) {
//       toast({
//         title: "Something went wrong.",
//         variant: "destructive",
//         description: "Sorry there is some problem in processing your request.",
//       });
//     }
//   };

//   useEffect(() => {
//     fetchExam();
//   }, []);

//   const contentGroup = useMemo(() => {
//     const temp: Record<string, IContentGroup> = {};
//     if (questionSet) {
//       questionSet?.raw_exam_data?.content_groups?.forEach((e) => {
//         temp[e.content_id] = e;
//       });
//     }
//     return temp;
//   }, [questionSet]);
//   const handleQuestionChange = (index: number) => {
//     if (index == activeIndex) return;
//     if (answers?.[activeIndex]?.state == "UN-ATTEMPTED") {
//       setValue(`answers.${activeIndex}.state`, "NOT-ANSWERED");
//     }
//     setValue("activeQuestionIndex", index);
//   };

//   const handleSaveNNext = () => {
//     if (tempAnswerIndex) {
//       setValue(`answers.${activeIndex}.state`, "ATTEMPTED");
//       setValue(`answers.${activeIndex}.selectedOption`, tempAnswerIndex);
//       setValue("tempAnswerIndex", null);
//     }
//   };

//   const questions = questionSet?.raw_exam_data?.questions || [];
//   const activeContentReference = questions?.[activeIndex]?.content_reference_id;
//   const contentInfo = contentGroup[activeContentReference];

//   const endExamNavigatePath = isOpenMode
//     ? `/${paths.COMMUNITY_EXAMS}/banking/${paths.MCQ}`
//     : `/${paths.EXAMS}/banking/${param.categoryId}`;
//   const exitExam = () => {
//     navigate(endExamNavigatePath);
//   };

//   return (
//     <MathJaxContext>
//       <div className="md:max-h-[calc(100vh_-_54px)] md:h-[calc(100vh_-_54px)] flex flex-col ">
//         <div className="flex-1 md:overflow-hidden">
//           <div className="flex flex-col md:flex-row items-stretch md:max-h-full md:h-full">
//             <div className="flex-1 max-h-full flex flex-col">
//               <div className="text-start p-2 pl-4 shadow flex items-center">
//                 <span className="ml-auto flex-1">
//                   Que No {activeIndex + 1} /{" "}
//                   {questionSet?.number_of_questions || 0}
//                 </span>
//                 {!isScreenView && (
//                   <MCQMobileHeader
//                     answers={answers}
//                     onQuestionNumberClick={handleQuestionChange}
//                     activeQuestionIndex={activeIndex}
//                   />
//                 )}
//               </div>
//               <div className="flex-1 flex flex-col md:flex-row md:overflow-hidden">
//                 {contentInfo && (
//                   <div className="md:w-1/2 min-w-[50%] overflow-auto p-4 pt-0 md:pt-4  text-pretty font-medium border-r-0 border-b md:border-b-0 md:border-r">
//                     {contentInfo?.instructions && (
//                       <p className="font-semibold mb-2 pt-4 md:pt-0">
//                         {contentInfo.instructions}
//                       </p>
//                     )}
//                     {contentInfo?.content && (
//                       <>
//                         {isMobileView ? (
//                           <p>{contentInfo.content}</p>
//                         ) : (
//                           <ReadMore text={contentInfo.content} />
//                         )}
//                       </>
//                     )}
//                   </div>
//                 )}
//                 <div className="md:max-h-[70vh] h-full flex flex-col p-4 flex-1">
//                   {questions.length && (
//                     <>
//                       <p className="items-start text-start pb-2">
//                         <MarkdownRender
//                           children={questions?.[activeIndex]?.question || ""}
//                         />
//                       </p>
//                       <Controller
//                         name={"tempAnswerIndex"}
//                         control={control}
//                         render={({ field }) => (
//                           <AnswerOptions
//                             name={field.name}
//                             onChange={field.onChange}
//                             options={questions[activeIndex].options}
//                             selected={field.value}
//                           />
//                         )}
//                       />
//                     </>
//                   )}
//                 </div>
//               </div>
//               <div className="min-h-12 flex gap-2 p-2 border-orange-300 border-t">
//                 <Controller
//                   name={`activeQuestionIndex`}
//                   control={control}
//                   render={({ field }) => (
//                     <Button
//                       size={"sm"}
//                       variant={"success"}
//                       className="justify-self-end ml-auto"
//                       type="button"
//                       disabled={!field.value}
//                       onClick={() => {
//                         const nextIdx = field.value - 1;
//                         field.value + 1 <= questions.length &&
//                           field.onChange(nextIdx);
//                       }}
//                     >
//                       PREV
//                     </Button>
//                   )}
//                 />
//                 <Controller
//                   name={`activeQuestionIndex`}
//                   control={control}
//                   render={({ field }) => (
//                     <Button
//                       size={"sm"}
//                       variant={"success"}
//                       className="justify-self-end ml-2"
//                       type="button"
//                       disabled={field.value == questions.length - 1}
//                       onClick={() => {
//                         const nextIdx = field.value + 1;

//                         field.value + 1 <= questions.length &&
//                           field.onChange(nextIdx);
//                       }}
//                     >
//                       NEXT
//                     </Button>
//                   )}
//                 />
//               </div>
//             </div>

//             {isScreenView && (
//               <div className=" bg-neutral-100/75 min-w-72 md:w-auto w-full h-full px-2 ">
//                 <MCQQuestionPallet
//                   answers={answers}
//                   onQuestionNumberClick={handleQuestionChange}
//                   activeQuestionIndex={activeIndex}
//                 />
//               </div>
//             )}
//           </div>
//         </div>
//       </div>
//     </MathJaxContext>
//   );
// };

// export default MCQSubmission;
