import axiosInstance from "./base";

export const getUserProfile = async () => {
  const response = await axiosInstance.get("/user");
  return response.data;
};
