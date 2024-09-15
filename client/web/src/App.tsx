import { RouterProvider } from "react-router-dom";
import "./App.css";
import { router } from "./routes/router";
import { Toaster } from "./componnets/base/toaster";
import useSessionStore from "./store/auth-store";
import { useEffect } from "react";
import SessionProvider from "./componnets/shared/session-provider";

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
