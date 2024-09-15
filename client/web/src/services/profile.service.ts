import axiosInstance from "./base";

export const getUserProfile = async () => {
  const response = await axiosInstance.get("/user");
  return response.data;
};
export const updateUserProfile = async (data: any) => {
  const response = await axiosInstance.put("/user", data);
  return response.data;
};
export const getUserTransactions = async () => {
  const response = await axiosInstance.get("/user/transactions");
  return response.data;
};
