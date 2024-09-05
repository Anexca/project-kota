import { RouterProvider } from "react-router-dom";
import "./App.css";
import { router } from "./routes/router";
import { Toaster } from "./componnets/base/toaster";
import useSessionStore from "./store/auth-store";
import { useEffect } from "react";

function App() {
  const { loadSession } = useSessionStore();
  useEffect(() => {
    loadSession();
  }, []);
  return (
    <>
      <Toaster />
      <RouterProvider router={router} />
    </>
  );
}

export default App;
