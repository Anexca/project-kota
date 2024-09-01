import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import ExamCenter from "./pages/exam-center/exam-center";
import DiscriptiveExam from "./pages/discriptive-exam/discriptive-exam";
import { Dashboard } from "./pages/login/login";
import { RegisterForm } from "./pages/register/register";

function App() {
  return (
    <>
      {/* <ExamCenter/> */}
      {/* <DiscriptiveExam /> */}
      <Dashboard />
      {/* <RegisterForm /> */}
    </>
  );
}

export default App;
