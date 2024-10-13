import axiosInstance from "./base";

export const evaluateMcqExam = async ({
  id,
  payload,
}: {
  id: string;
  payload: any;
}) => {
  const res = await axiosInstance.post(
    `/exams/banking/mcq/${id}/evaluate`,
    payload
  );
  return res.data;
};
