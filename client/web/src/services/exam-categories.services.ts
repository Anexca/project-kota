import axiosInstance from "./base";

export const getQuestionsCategories = async (isOpenExam?: boolean) => {
  const response = await axiosInstance.get("/exams/banking", {
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
  const response = await axiosInstance.get(`/exams/banking/${categoryId}`);
  return response.data;
};
export const getCategoryById = async ({
  examGroupId,
}: {
  examGroupId: number | string;
}) => {
  const response = await axiosInstance.get(`/categories/exams/${examGroupId}`);
  return response.data;
};
