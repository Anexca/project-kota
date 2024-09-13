import { BaseSyntheticEvent, PropsWithChildren } from "react";
import {
  alertBackendForSubscription,
  buySubscription,
} from "../../../services/payment.service";

type Props = {
  subscriptionId: string;
  id: string;
  prefill: UserInfo;
};
type UserInfo = {
  name: string;
  email: string;
  contact: string;
};

const RazorpayButton = ({
  children,
  prefill,
  id,
}: PropsWithChildren<Props>) => {
  const getSubscriptionId = async () => {
    const res = await buySubscription(id);
    return res.data;
  };

  const handleRazorPay = async (e: {
    razorpay_payment_id: string;
    razorpay_signature: string;
    razorpay_subscription_id: string;
  }) => {
    console.log(e);
    await alertBackendForSubscription({
      paymentId: e.razorpay_payment_id,
      signature: e.razorpay_signature,
      userSubsId: id,
    });
  };
  const openRazorPay = async (e: BaseSyntheticEvent) => {
    const res = await getSubscriptionId();
    const options = {
      key: import.meta.env.VITE_RAZORPAY_KEY,
      subscription_id: res.provider_subscription_id,
      name: "PsuedoTest",
      description: "Golden Pass",
      handler: handleRazorPay,
      prefill: prefill,
      redirect: false,
      theme: {
        color: "#7e22ce",
      },
    };

    ///@ts-ignore
    const rpay = new Razorpay(options);

    rpay.open();
    e.preventDefault();
  };
  return <button onClick={openRazorPay}>{children}</button>;
};

export default RazorpayButton;
