import axiosInstance from "./base";

const getPlans = async () => {
  const response = await axiosInstance.get(`/subscriptions`);
  return response.data;
};
const buySubscription = async (id: string) => {
  const response = await axiosInstance.post(`/subscriptions/${id}/buy`);
  return response.data;
};

const alertBackendForSubscription = async ({
  paymentId,
  signature,
  userSubsId,
}: {
  userSubsId: string;
  paymentId: string;
  signature: string;
}) => {
  const response = await axiosInstance.post(
    `/user-subscriptions/${userSubsId}/activate`,
    {
      payment_id: paymentId,
      signature,
    }
  );
  return response.data;
};

export { getPlans, buySubscription, alertBackendForSubscription };
