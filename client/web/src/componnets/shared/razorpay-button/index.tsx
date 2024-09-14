import {
  BaseSyntheticEvent,
  PropsWithChildren,
  useMemo,
  useState,
} from "react";
import { useToast } from "../../../hooks/use-toast";
import {
  alertBackendForSubscription,
  buySubscription,
} from "../../../services/payment.service";
import useUserProfileStore from "../../../store/user-info-store";
import { Button } from "../../base/button/button";
import Loader from "../loder";

type Props = {
  subscriptionId: string;
  id: string;
  isDisabled?: boolean;
};

const RazorpayButton = ({
  children,
  id,
  subscriptionId,
  isDisabled,
}: PropsWithChildren<Props>) => {
  const { isComplete, profile, getProfile } = useUserProfileStore();
  const { toast } = useToast();
  const [loading, setLoading] = useState(false);
  const getSubscriptionId = async () => {
    const res = await buySubscription(id);
    return res.data;
  };

  const handleRazorPay = async (e: {
    razorpay_payment_id: string;
    razorpay_signature: string;
    razorpay_subscription_id: string;
    id: string;
  }) => {
    setLoading(true);
    try {
      await alertBackendForSubscription({
        paymentId: e.razorpay_payment_id,
        signature: e.razorpay_signature,
        userSubsId: e.id,
      });
      await getProfile();
    } catch (error) {
      toast({
        title: "Something went wrong",
        description:
          "There is some error on our side. Please raise a query to us if your money is deducted.",
        variant: "destructive",
      });
    } finally {
      setLoading(false);
    }
  };
  const openRazorPay = async (e: BaseSyntheticEvent) => {
    if (!isComplete) {
      toast({
        title: "Suggestion",
        description: "Plese complete your user profile in profile section.",
        variant: "warning",
        duration: 10000,
      });
      return;
    }
    const res = await getSubscriptionId();
    const options = {
      key: import.meta.env.VITE_RAZORPAY_KEY,
      subscription_id: res.provider_subscription_id,
      name: "PsuedoTest",
      description: "Golden Pass",
      handler: async (e: any) => await handleRazorPay({ ...e, id: res.id }),
      prefill: {
        name: `${profile.first_name} ${profile.last_name}`,
        email: profile.email,
        contact: profile.phone_number,
      },
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

  const isActive = useMemo(() => {
    return profile?.active_subscriptions?.some(
      (i) => i.provider_plan_id == subscriptionId
    );
  }, [profile]);
  return (
    <Button
      disabled={isActive || isDisabled}
      className="w-full"
      variant={"success"}
      onClick={openRazorPay}
    >
      {loading ? <Loader /> : children}
    </Button>
  );
};

export default RazorpayButton;
