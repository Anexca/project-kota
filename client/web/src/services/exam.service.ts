import { IPastExamAttempt } from "../interface/past-submission";
import axiosInstance from "./base";

export const getQuestions = async (isOpenExam?: boolean) => {
  const response = await axiosInstance.get("/exams/banking/descriptive", {
    params: {
      isopen: !!isOpenExam,
    },
  });
  return response.data;
};

export const getQuestionById = async (questionId: string) => {
  const response = await axiosInstance.get(`/exams/${questionId}`);
  return response.data;
};
export const sendAnswerForAssesment = async ({
  questionId,
  answer,
  completedTime,
}: {
  questionId: number;
  answer: string;
  completedTime: number;
}) => {
  const response = await axiosInstance.post(
    `/exams/banking/descriptive/${questionId}/evaluate`,
    { completed_seconds: completedTime, content: answer }
  );
  return response.data;
};
export const getAssesmetsResult = async (assesmetId: number) => {
  const response = await axiosInstance.get(`/exams/assesments/${assesmetId}`);
  return response.data;
};
export const getPastSubmission = async (examId: number) => {
  const response = await axiosInstance.get(`/exams/${examId}/assessments`);
  return response.data;
};
export const getPastAttemptedSubmissions = async () => {
  const response = await axiosInstance.get<{ data: IPastExamAttempt[] }>(
    `/exams/history`
  );
  return response.data;
};
