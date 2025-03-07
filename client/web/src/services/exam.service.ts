import { IPastExamAttempt } from "../interface/past-submission";
import { IMCQQuestionSet } from "../interface/question";
import { FilterPagination, IPaginationType } from "../interface/utils";
import axiosInstance from "./base";

export const getQuestionsCategories = async (isOpenExam?: boolean) => {
  const response = await axiosInstance.get("/exams/banking/descriptive", {
    params: {
      isopen: !!isOpenExam,
    },
  });
  return response.data;
};

export const getQuestions = async ({
  categoryId,
}: {
  categoryId: number | string;
}) => {
  const response = await axiosInstance.get(
    `/exams/banking/descriptive/${categoryId}`
  );
  return response.data;
};

export const getQuestionById = async (questionId: string, isOpen?: boolean) => {
  const response = await axiosInstance.get(`/exams/${questionId}`, {
    params: {
      isopen: isOpen,
    },
  });
  return response.data;
};
export const sendAnswerForAssesment = async ({
  questionId,
  answer,
  completedTime,
  isOpen,
}: {
  questionId: number;
  answer: string;
  completedTime: number;
  isOpen?: boolean;
}) => {
  const response = await axiosInstance.post(
    `/exams/banking/descriptive/${questionId}/evaluate`,
    { completed_seconds: completedTime, content: answer },
    {
      params: {
        isopen: isOpen,
      },
    }
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
export const getPastAttemptedSubmissions = async (
  filters?: FilterPagination
) => {
  const response = await axiosInstance.get<{
    data: IPastExamAttempt[];
    pagination: IPaginationType;
  }>(`/exams/history`, { params: filters });
  return response.data;
};

export const getMCQQuestionsCategories = async (isOpenExam?: boolean) => {
  const response = await axiosInstance.get("/exams/banking/mcq", {
    params: {
      isopen: !!isOpenExam,
    },
  });
  return response.data;
};
export const getMCQQuestions = async ({
  categoryId,
}: {
  categoryId: number | string;
}) => {
  const response = await axiosInstance.get(`/exams/banking/mcq/${categoryId}`);
  return response.data;
};
export const getMCQQuestionById = async (questionId: string) => {
  const response = await axiosInstance.get<{ data: IMCQQuestionSet }>(
    `/exams/${questionId}`,
    {}
  );
  return response.data;
};
