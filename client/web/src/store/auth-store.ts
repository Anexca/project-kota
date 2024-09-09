import { create } from "zustand";
import { Session } from "@supabase/supabase-js"; // Import Supabase types
import { supabase } from "../supabase/client";

// Define the store's state and actions types
interface SessionStore {
  session: Session | null;
  setSession: (session: Session | null) => void;
  clearSession: () => void;
  loadSession: () => Promise<Session | null>;
  subscribeToAuthChanges: () => void;
  logout: () => Promise<void>;
}

// Create the Zustand store
export const sessionStore = create<SessionStore>((set) => ({
  session: null,

  setSession: (session) => set({ session }),

  clearSession: () => set({ session: null }),

  loadSession: async () => {
    const {
      data: { session },
    } = await supabase.auth.getSession();
    set({ session });
    return session;
  },

  subscribeToAuthChanges: () => {
    const {
      data: { subscription },
    } = supabase.auth.onAuthStateChange((_event, session) => {
      set({ session });
    });

    // Clean up subscription when not needed
    return () => {
      subscription.unsubscribe();
    };
  },
  logout: async () => {
    try {
      localStorage.clear();
      await supabase.auth.signOut();
      set({ session: null }); // Clear session on successful logout
    } catch (error) {
      set({ session: null }); // Clear session on successful logout
    }
  },
}));
const useSessionStore = sessionStore;
export default useSessionStore;
