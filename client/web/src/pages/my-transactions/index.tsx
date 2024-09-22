import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import Icon from "../../componnets/base/icon";
import { useToast } from "../../hooks/use-toast";
import { IMyTransactions } from "../../interface/user";
import { paths } from "../../routes/route.constant";
import { getUserTransactions } from "../../services/profile.service";
import Chip from "../../componnets/base/chip";

const paymentStatusComponents: { [key: string]: JSX.Element } = {
  CREATED: (
    <Chip icon={"clock"} variant={"info"}>
      Payment Created
    </Chip>
  ),
  AUTHORIZED: <Chip icon={"thumbs_up"}>Payment Authorized</Chip>,
  SUCCESS: (
    <Chip icon={"check_solid"} variant={"success"}>
      Payment Captured
    </Chip>
  ),
  FAILED: (
    <Chip icon={"xmark"} variant={"danger"}>
      Payment Failed
    </Chip>
  ),
  VOID: (
    <Chip icon={"target"} variant={"success"}>
      Payment Void
    </Chip>
  ),
  NOT_ATTEMPTED: (
    <Chip icon={"undo_alt"} variant={"warning"}>
      Payment Not Attempted
    </Chip>
  ),
  PENDING: (
    <Chip icon={"hourglass_half"} variant={"warning"}>
      Payment Pending
    </Chip>
  ),
  PROCESSING: (
    <Chip icon={"sync"} variant={"outline"}>
      Payment Processing
    </Chip>
  ),
  CANCELLED: (
    <Chip icon={"ban"} variant={"danger"}>
      Payment Cancelled
    </Chip>
  ),
  USER_DROPPED: (
    <Chip icon={"exclaimation_circle"} variant={"warning"}>
      User Dropped
    </Chip>
  ),
};

// const paymentMethod: { [key: string]: string } = {
//   upi: "UPI",
//   card: "CARD",
// };
const MyTransactions = () => {
  const [transactions, setTransactions] = useState<IMyTransactions[]>([]);
  const [loading, setLoading] = useState(false);
  const { toast } = useToast();
  const getQuestionsList = async () => {
    setLoading(true);
    try {
      const data = await getUserTransactions();
      setTransactions(data.data || []);
    } catch (error) {
      toast({
        title: "Oh ho Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in processing your request.",
      });
    }
    setLoading(false);
  };
  useEffect(() => {
    getQuestionsList();
  }, []);

  return (
    <div className="pt-2 w-full md:max-w-2xl 2xl:max-w-2xl mx-auto flex flex-col gap-2 p-4">
      <div className="py-2">
        <div className="flex gap-2 items-center">
          <Link to={`/${paths.HOMEPAGE}`} className="p-0">
            <Icon icon="arrow_back" className="text-info text-lg" />
          </Link>
          <span className="text-lg font-semibold">
            All of your transactions.
          </span>
        </div>
        <div className="flex items-center gap-2  text-sm"></div>
      </div>
      {loading ? (
        <div className="flex flex-col gap-2 justify-center items-center">
          <span className="rounded-full w-8 h-8 animate-spin border-2 border-info border-t-info/30"></span>
          Getting Your Transactions
        </div>
      ) : (
        <div className="animate-fadeIn flex flex-col gap-2">
          {transactions.length ? (
            transactions.map((item) => {
              return (
                <article className="rounded-md shadow-sm bg-white flex flex-col gap-1 p-3 px-4 md:p3 text-sm">
                  <div className="flex flex-col md:flex-row">
                    <div className="flex-1">
                      <p className="font-medium text-balance md:text-pretty text-black text-lg mb-2">
                        <Icon
                          icon="rupee"
                          className="text-destructive text-base"
                        />{" "}
                        {item.amount}
                      </p>
                      <div className="flex gap-2">
                        <p className="font-medium text-balance md:text-pretty text-black mb-2">
                          Date :{" "}
                          <span className="font-semibold">
                            {" "}
                            {new Date(item.payment_date).toLocaleDateString()}
                          </span>
                        </p>
                        <p className="font-medium text-balance md:text-pretty text-black mb-2">
                          Time :{" "}
                          <span className="font-semibold">
                            {new Date(item.payment_date).toLocaleTimeString()}
                          </span>
                        </p>
                      </div>
                    </div>

                    <div className="flex flex-col items-stretch gap-2 md:justify-center md:w-48">
                      <div className="flex gap-2 flex-wrap ">
                        {paymentStatusComponents[item.payment_status]}
                        <Chip icon="bank" variant={"success"}>
                          {item.payment_method || "Unknown"}
                        </Chip>
                      </div>
                    </div>
                  </div>
                </article>
              );
            })
          ) : (
            <div className="flex flex-col gap-2 justify-center items-center">
              You don't have any transactions yet. Please buy a subscription
              plan.
            </div>
          )}
        </div>
      )}
    </div>
  );
};

export default MyTransactions;
