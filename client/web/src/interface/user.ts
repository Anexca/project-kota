export interface IUserProfile {
  id: string;
  email: string;
  first_name: string;
  last_name: string;
  phone_number: string;
  payment_provider_customer_id: string;
  active_subscriptions: IUserSubscription[] | null;
}

export interface IUserSubscription {
  subscription_id: number;
  provider_subscription_id: string;
  provider_plan_id: string;
  duration_in_months: 1;
  start_date: Date;
  end_date: Date;
  payment_details: {
    amount: number;
    payment_date: Date;
    payment_status: string;
    payment_method: string;
  };
}

export interface IMyTransactions {
  amount: number;
  payment_date: string;
  payment_status: string;
  payment_method: string;
}
