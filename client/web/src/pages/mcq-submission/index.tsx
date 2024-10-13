import { useEffect, useMemo, useState } from "react";
import { Link, useLocation, useParams } from "react-router-dom";
import { Button } from "../../componnets/base/button/button";
import Chip from "../../componnets/base/chip";
import Icon from "../../componnets/base/icon";
import Container from "../../componnets/shared/container";
import MCQSubmissionView from "../../componnets/shared/mcq-submission-view";
import { QUESTION_STATE } from "../../constants/shared";
import { useToast } from "../../hooks/use-toast";
import { IMCQSubmission } from "../../interface/mcq-submission";
import { IContentGroup, IMCQQuestionSet } from "../../interface/question";
import { paths } from "../../routes/route.constant";
import {
  getAssesmetsResult,
  getMCQQuestionById,
} from "../../services/exam.service";
import useUserProfileStore from "../../store/user-info-store";

const MCQSubmission = ({ backLink }: { backLink?: string }) => {
  const location = useLocation();
  const param = useParams();
  const [assessment, setAssessment] = useState<IMCQSubmission | null>(null);
  const [questionSet, setQuestionSet] = useState<IMCQQuestionSet | null>(null);
  const [summaryView, setSummaryView] = useState(true);
  const profile = useUserProfileStore();
  const { toast } = useToast();
  const assesmentId = Math.abs(Number.parseInt(param.assesmentId || ""));
  const questionId = Math.abs(Number.parseInt(param.questionId || ""));

  const categoryId = location?.state?.categoryId as any;

  const fetchSubmission = async () => {
    try {
      const res = await getAssesmetsResult(assesmentId);
      setAssessment(res.data);
    } catch (error) {
      toast({
        title: "Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in processing your request.",
      });
    }
  };

  const fetchExam = async () => {
    try {
      const res = await getMCQQuestionById(`${questionId}`);
      setQuestionSet(res.data);
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
    fetchSubmission();
  }, []);

  const clearedExam = useMemo(() => {
    if (!assessment) return false;
    return assessment.obtained_marks >= assessment.cutoff_marks;
  }, [assessment]);

  const examMetaData = useMemo(() => {
    if (!assessment || !questionSet) {
      return {
        activeQuestionIndex: 0,
        answers: {},
        activeSection: "",
        sectionList: [],
        contentGroup: {},
        perQuestionMarks: 0,
        negativeMark: 0,
      };
    }
    const assessmentList =
      assessment?.raw_user_submission.attempted_questions || [];
    const ques = questionSet?.raw_exam_data?.sections || {};
    const sets = Object.keys(ques);
    const questionsSet: any = {};
    const questionTimeMap: Record<string, number> = {};
    const answerMap: any = {};
    assessmentList?.forEach((i) => {
      answerMap[i.question_number] = {
        timeTaken: i.time_taken_in_seconds,
        optionTaken: i.user_selected_option_index[0],
      };
    });
    sets.forEach((item) => {
      const quesArray = ques[item].map((i) => {
        questionTimeMap[i.question_number] = 0;
        return {
          state: answerMap[i.question_number]
            ? QUESTION_STATE.ATTEMPTED
            : QUESTION_STATE.NOT_ANSWERED,
          selectedOption: answerMap[i.question_number]?.optionTaken,
          questionNumber: i.question_number,
          timeTaken: answerMap[i.question_number]?.timeTaken,
          correctAnswer: i.answer[0],
        };
      });
      questionsSet[item] = quesArray;
    });
    const temp: Record<string, IContentGroup> = {};
    if (questionSet) {
      questionSet?.raw_exam_data?.content_groups?.forEach((e) => {
        temp[e.content_id] = e;
      });
    }

    return {
      activeQuestionIndex: 0,
      answers: questionsSet,
      activeSection: sets[0],
      sectionList: sets,
      contentGroup: temp,
      perQuestionMarks:
        assessment!.total_marks / questionSet!.number_of_questions,
      negativeMark: questionSet!.negative_marking,
    };
  }, [questionSet, assessment]);
  const exitLink = categoryId
    ? `/${paths.EXAMS}/banking/${categoryId}?activeTab=mcq`
    : `/${paths.EXAMS}/${paths.MY_SUMBISSIONS}`;

  return summaryView ? (
    <Container>
      <div className="flex flex-col  ">
        <div className="p-4 md:p-10 md:pr-2">
          <div className="flex items-center flex-wrap capitalize mb-2 font-medium text-lg">
            <Link className="text-info mr-2 text-sm" to={backLink || exitLink}>
              <Icon icon="arrow_back" /> Back
            </Link>{" "}
            {questionSet?.exam_name?.split("_").join(" ").toLowerCase()}{" "}
          </div>
          <div className="font-semibold text-lg">
            {clearedExam ? "Congratulations!!!" : "Oh Sorry!!!"}
          </div>
          <div className="text-xl font-semibold text-info">
            {profile.profile.first_name} {profile.profile.last_name}
          </div>
          <div className="font-semibold ">
            {clearedExam
              ? "You have met the desired cutoff score."
              : "You did not met the passing cutoff."}
          </div>
          <div className="font-semibold">
            Your Score is{" "}
            <Chip
              icon={"bank"}
              variant={clearedExam ? "success" : "destructive"}
              className="text-base"
            >
              {assessment?.obtained_marks ?? 0}
            </Chip>{" "}
            And cutoff is{" "}
            <Chip icon={"target"} className="text-base">
              {assessment?.cutoff_marks ?? ""}
            </Chip>
          </div>

          <Button
            onClick={() => setSummaryView(false)}
            className="mt-3"
            size={"sm"}
          >
            View Solution
          </Button>
        </div>
        <div className="p-4 md:mx-auto">
          <div className="grid  grid-cols-2 md:grid-cols-3  gap-2">
            <div className="flex flex-col items-center px-4 p-2 bg-white border border-gray-200 rounded-lg shadow hover:scale-105 transition-transform">
              <div className="font-bold text-2xl bg-info/25 text-info flex items-center justify-center w-12 h-12 rounded-full">
                <Icon icon="target" />
              </div>
              <div className="mt-2 text-sm font-semibold  text-gray-900">
                Accuracy{" "}
                <span className="font-bold">
                  {assessment?.raw_assesment_data.summary.accuracy}{" "}
                </span>
              </div>
            </div>

            <div className="flex flex-col items-center px-4 p-2 bg-white border border-gray-200 rounded-lg shadow hover:scale-105 transition-transform">
              <div className="font-bold text-2xl bg-warning/25 text-warning flex items-center justify-center w-12 h-12 rounded-full">
                <Icon icon="tasks" />
              </div>
              <div className="mt-2 text-sm font-semibold  text-gray-900">
                Attempted{" "}
                <span className="font-bold">
                  {assessment?.raw_assesment_data.summary.attempted}{" "}
                </span>
              </div>
            </div>

            <div className="flex flex-col items-center px-4 p-2 bg-white border border-gray-200 rounded-lg shadow hover:scale-105 transition-transform">
              <div className="font-bold text-2xl bg-red-400/25 text-red-500 flex items-center justify-center w-12 h-12 rounded-full">
                <Icon icon="check" />
              </div>
              <div className="mt-2 text-sm font-semibold  text-gray-900">
                Incorrect{" "}
                <span className="font-bold">
                  {assessment?.raw_assesment_data.summary.incorrect}{" "}
                </span>
              </div>
            </div>
            <div className="flex flex-col items-center px-4 p-2 bg-white border border-gray-200 rounded-lg shadow hover:scale-105 transition-transform">
              <div className="font-bold text-2xl bg-green-400/25 text-success flex items-center justify-center w-12 h-12 rounded-full">
                <Icon icon="check" />
              </div>
              <div className="mt-2 text-sm font-semibold  text-gray-900">
                Correct{" "}
                <span className="font-bold">
                  {assessment?.raw_assesment_data.summary.correct}{" "}
                </span>
              </div>
            </div>
            <div className="flex flex-col items-center px-4 p-2 bg-white border border-gray-200 rounded-lg shadow hover:scale-105 transition-transform">
              <div className="font-bold text-2xl bg-blue-400/25 text-blue-500 flex items-center justify-center w-12 h-12 rounded-full">
                <Icon icon="forward" />
              </div>
              <div className="mt-2 text-sm font-semibold  text-gray-900">
                Skipped{" "}
                <span className="font-bold">
                  {assessment?.raw_assesment_data.summary.attempted}{" "}
                </span>
              </div>
            </div>
            <div className="flex flex-col items-center px-4 p-2 bg-white border border-gray-200 rounded-lg shadow hover:scale-105 transition-transform">
              <div className="font-bold text-2xl bg-pink-400/25 text-pink-500 flex items-center justify-center w-12 h-12 rounded-full">
                <Icon icon="clock" />
              </div>
              <div className="mt-2 text-sm font-semibold  text-gray-900">
                Time Taken{" "}
                <span className="font-bold">
                  {assessment?.completed_seconds}
                  {" sec"}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Container>
  ) : (
    <MCQSubmissionView
      {...examMetaData}
      questionSet={questionSet}
      numberOfQuestions={questionSet?.number_of_questions || 0}
      assessment={assessment?.raw_user_submission.attempted_questions || []}
      backLink={backLink || exitLink}
    />
  );
};

export default MCQSubmission;
