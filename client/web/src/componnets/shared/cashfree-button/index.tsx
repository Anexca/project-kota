import {
  BaseSyntheticEvent,
  PropsWithChildren,
  useMemo,
  useState,
} from "react";
import { useToast } from "../../../hooks/use-toast";
import { delay } from "../../../lib/utils";
import { buySubscription } from "../../../services/payment.service";
import useUserProfileStore from "../../../store/user-info-store";
import { Button } from "../../base/button/button";
import Loader from "../loder";
///@ts-ignore
import { load } from "@cashfreepayments/cashfree-js";
import { paths } from "../../../routes/route.constant";
import { useNavigate } from "react-router-dom";

type Props = {
  subscriptionId: string;
  id: string;
  isDisabled?: boolean;
};

const CashFreeButton = ({
  children,
  id,
  subscriptionId,
  isDisabled,
}: PropsWithChildren<Props>) => {
  const { isComplete, profile, getProfile } = useUserProfileStore();
  const { toast } = useToast();
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const getSubscriptionId = async () => {
    try {
      const res = await buySubscription(
        id,
        `${import.meta.env.VITE_CASHFREE_REDIRECT_URL}/${paths.PRICING_PLAN}`
      );
      return res.data;
    } catch (error) {
      toast({
        title: "Something went wrong",
        description:
          "There is some error on our side. Please raise a query to us at support@pueudotest.pro if your money is deducted.",
        variant: "destructive",
      });
      return;
    }
  };

  const handleRazorPay = async () => {
    try {
      await delay(2000);
      await getProfile();
      toast({
        title: "Congratulations 🎉",
        description:
          "Your payment is successful. Your plan is active now. All the best for preparation.",
        variant: "success",
      });
      navigate(`/${paths.HOMEPAGE}`);
    } catch (error) {
      toast({
        title: "Something went wrong",
        description:
          "There is some error on our side. Please raise a query to us at support@pueudotest.pro if your money is deducted.",
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
        description: "Please complete your user profile in profile section.",
        variant: "warning",
        duration: 10000,
      });
      return;
    }
    setLoading(true);

    const res = await getSubscriptionId();
    if (!res) {
      setLoading(false);
      return;
    }

    ///@ts-ignore
    const cashfree = await load({
      mode: import.meta.env.VITE_CASHFREE_PAYMENT_MODE, //or production
    });

    let checkoutOptions = {
      paymentSessionId: res.payment_session_id,
      redirectTarget: "_modal", //optional ( _self, _blank, or _top)
    };

    const cashfreeRes = await cashfree.checkout(checkoutOptions);

    if (cashfreeRes.error) {
      if (cashfreeRes.error.code == "payment_aborted") {
        // This will be true whenever user clicks on close icon inside the modal or any error happens during the payment
        toast({
          title: "Payment Cancelled",
          description: "Payment is cancelled by you.",
          variant: "destructive",
        });
      } else {
        toast({
          title: "Something went wrong",
          description:
            "There is some error on our side. Please raise a query to us at support@pueudotest.pro if your money is deducted.",
          variant: "destructive",
        });
      }
      setLoading(false);
    }

    if (cashfreeRes.paymentDetails) {
      // This will be called whenever the payment is completed irrespective of transaction status
      await handleRazorPay();
    }
    e.preventDefault();
  };

  const isActive = useMemo(() => {
    return profile?.active_subscriptions?.some(
      (i) => i.provider_plan_id == subscriptionId
    );
  }, [profile]);
  return (
    <Button
      disabled={isActive || isDisabled || loading}
      className="w-full"
      variant={"success"}
      onClick={openRazorPay}
    >
      {loading ? <Loader /> : children}
    </Button>
  );
};

export default CashFreeButton;
