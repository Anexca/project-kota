import { create } from "zustand";
import { IUserProfile } from "../interface/user";
import { getUserProfile } from "../services/profile.service";

interface UserProfileState {
  profile: IUserProfile;
  setProfile: (profile: IUserProfile) => void;
  clearProfile: () => void;
  getProfile: () => Promise<IUserProfile>;
  disableNotification: () => void;
  notificationShow: boolean;
  isComplete: boolean;
}

const emptyResponse: IUserProfile = {
  email: "",
  first_name: "",
  id: "",
  last_name: "",
  payment_provider_customer_id: "",
  phone_number: "",
  active_subscriptions: [],
};

const useUserProfileStore = create<UserProfileState>((set) => ({
  profile: emptyResponse,
  notificationShow: true,
  isComplete: false,
  setProfile: (profile) => {
    const isComplete = !!(
      profile.first_name &&
      profile.last_name &&
      profile.phone_number
    );
    set((p) => ({ ...p, isComplete, profile }));
  },
  clearProfile: () => set((p) => ({ ...p, profile: emptyResponse })),
  getProfile: async () => {
    try {
      const res = await getUserProfile();

      const { first_name, last_name, phone_number } = res.data;
      const isComplete = !!(first_name && last_name && phone_number);

      set((p) => ({
        ...p,
        profile: res.data,
        isComplete,
      }));
      return res.data;
    } catch (error) {
      throw error;
    }
  },
  disableNotification: () => {
    set((p) => ({ ...p, notificationShow: false }));
  },
}));

export default useUserProfileStore;
