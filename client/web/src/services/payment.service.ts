import { paths } from "../routes/route.constant";
import axiosInstance from "./base";

const getPlans = async () => {
  const response = await axiosInstance.get(`/subscriptions`);
  return response.data;
};
const buySubscription = async (id: string) => {
  const response = await axiosInstance.post<{
    data: {
      id: number;
      status: string;
      subscription_id: string;
      payment_session_id: string;
    };
  }>(
    `/subscriptions/${id}/buy`,
    {},
    { params: { returnUrl: `http://localhost:5173/${paths.PRICING_PLAN}` } }
  );
  return response.data;
};

const alertBackendForSubscription = async ({
  paymentId,
}: {
  paymentId: number;
}) => {
  const response = await axiosInstance.post(
    `/user-subscriptions/${paymentId}/activate`
  );
  return response.data;
};
const cancelSubscription = async (id: string) => {
  const response = await axiosInstance.post(`/user-subscriptions/${id}/cancel`);
  return response.data;
};
export {
  getPlans,
  buySubscription,
  alertBackendForSubscription,
  cancelSubscription,
};
