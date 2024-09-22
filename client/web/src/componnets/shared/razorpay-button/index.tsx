// import {
//   BaseSyntheticEvent,
//   PropsWithChildren,
//   useMemo,
//   useState,
// } from "react";
// import { useToast } from "../../../hooks/use-toast";
// import {
//   alertBackendForSubscription,
//   buySubscription,
// } from "../../../services/payment.service";
// import useUserProfileStore from "../../../store/user-info-store";
// import { Button } from "../../base/button/button";
// import Loader from "../loder";
// import { delay } from "../../../lib/utils";

// type Props = {
//   subscriptionId: string;
//   id: string;
//   isDisabled?: boolean;
// };

// const RazorpayButton = ({
//   children,
//   id,
//   subscriptionId,
//   isDisabled,
// }: PropsWithChildren<Props>) => {
//   const { isComplete, profile, getProfile } = useUserProfileStore();
//   const { toast } = useToast();
//   const [loading, setLoading] = useState(false);
//   const getSubscriptionId = async () => {
//     try {
//       const res = await buySubscription(id);
//       return res.data;
//     } catch (error) {
//       toast({
//         title: "Something went wrong",
//         description:
//           "There is some error on our side. Please raise a query to us at support@pueudotest.pro if your money is deducted.",
//         variant: "destructive",
//       });
//       return;
//     }
//   };

//   const handleRazorPay = async (e: {
//     razorpay_payment_id: string;
//     razorpay_signature: string;
//     razorpay_subscription_id: string;
//     id: string;
//   }) => {
//     try {
//       await alertBackendForSubscription({
//         paymentId: e.razorpay_payment_id,
//         signature: e.razorpay_signature,
//         userSubsId: e.id,
//       });
//       await delay(2000);
//       await getProfile();
//     } catch (error) {
//       toast({
//         title: "Something went wrong",
//         description:
//           "There is some error on our side. Please raise a query to us at support@pueudotest.pro if your money is deducted.",
//         variant: "destructive",
//       });
//     } finally {
//       setLoading(false);
//     }
//   };
//   const openRazorPay = async (e: BaseSyntheticEvent) => {
//     if (!isComplete) {
//       toast({
//         title: "Suggestion",
//         description: "Please complete your user profile in profile section.",
//         variant: "warning",
//         duration: 10000,
//       });
//       return;
//     }
//     setLoading(true);

//     const res = await getSubscriptionId();
//     if (!res) {
//       setLoading(false);
//       return;
//     }
//     const options = {
//       key: import.meta.env.VITE_RAZORPAY_KEY,
//       subscription_id: res.provider_subscription_id,
//       name: "PsuedoTest",
//       description: "Golden Pass",
//       handler: async (e: any) => await handleRazorPay({ ...e, id: res.id }),
//       redirect: false,
//       theme: {
//         color: "#7e22ce",
//       },
//       modal: {
//         ondismiss: function () {
//           setLoading(false);
//         },
//       },
//     };

//     ///@ts-ignore
//     const rpay = new Razorpay(options);
//     rpay.open();
//     e.preventDefault();
//   };

//   const isActive = useMemo(() => {
//     return profile?.active_subscriptions?.some(
//       (i) => i.provider_plan_id == subscriptionId
//     );
//   }, [profile]);
//   return (
//     <Button
//       disabled={isActive || isDisabled || loading}
//       className="w-full"
//       variant={"success"}
//       onClick={openRazorPay}
//     >
//       {loading ? <Loader /> : children}
//     </Button>
//   );
// };

// export default RazorpayButton;
