import { RouterProvider } from "react-router-dom";
import "./App.css";
import { router } from "./routes/router";
import { Toaster } from "./componnets/base/toaster";

function App() {
  return (
    <>
      <Toaster />
      <RouterProvider router={router} />
    </>
  );
}

export default App;
