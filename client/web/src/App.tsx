import { useEffect } from "react";
import { RouterProvider } from "react-router-dom";
import "./App.css";
import { Toaster } from "./componnets/base/toaster";
import SessionProvider from "./componnets/shared/session-provider";
import { router } from "./routes/router";
import useSessionStore from "./store/auth-store";

function App() {
  const { loadSession } = useSessionStore();
  useEffect(() => {
    loadSession();
  }, []);
  return (
    <>
      <Toaster />
      <SessionProvider>
        <RouterProvider router={router} />
      </SessionProvider>
    </>
  );
}

export default App;
