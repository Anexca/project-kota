// import "./index.css";
import { useEffect } from "react";
import { Auth } from "@supabase/auth-ui-react";
import { ThemeSupa } from "@supabase/auth-ui-shared";
import { supabase } from "../../supabase/client";
import useSessionStore from "../../store/auth-store";

export default function SupabaseAuth() {
  const { loadSession, subscribeToAuthChanges, session, logout } =
    useSessionStore();

  useEffect(() => {
    loadSession();
    return subscribeToAuthChanges();
  }, []);

  if (!session) {
    return <Auth supabaseClient={supabase} appearance={{ theme: ThemeSupa }} />;
  } else {
    return (
      <div>
        Logged in!
        <button onClick={logout}>logout</button>
      </div>
    );
  }
}
