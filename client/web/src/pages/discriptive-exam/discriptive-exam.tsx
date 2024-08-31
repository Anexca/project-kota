import React, { useEffect, useState } from "react";
import { diffChars, diffLines, diffWords } from "diff"; // Importing the diff functions
import { modified, original } from "./testdata";

type Props = {};

// src/DiffChecker.js

const DiffChecker = () => {
  const [text1, setText1] = useState(original);
  const [text2, setText2] = useState(modified);
  const [diffResult, setDiffResult] = useState([]);

  const handleTextChange = (e) => {
    const { name, value } = e.target;
    if (name === "text1") {
      setText1(value);
    } else if (name === "text2") {
      setText2(value);
    }
  };

  const handleDiff = () => {
    // You can choose to use diffChars or diffLines based on your requirement
    const diff = diffWords(text1, text2);
    setDiffResult(diff);
  };

  useEffect(() => {
    handleDiff();
  }, []);
  console.log(diffResult);

  return (
    <div>
      <h1>Diff Checker</h1>
      <div>
        <textarea
          name="text1"
          value={text1}
          onChange={handleTextChange}
          placeholder="Enter first string"
          rows="10"
          cols="50"
        />
        <textarea
          name="text2"
          value={text2}
          onChange={handleTextChange}
          placeholder="Enter second string"
          rows="10"
          cols="50"
        />
      </div>
      <button onClick={handleDiff}>Check Diff</button>
      <div>
        <h2>Differences:</h2>
        <div
          style={{ textAlign: "center", width: "60vw", marginInline: "auto" }}
        >
          {diffResult.map((part, index) => (
            <span
              key={index}
              style={{
                color: part.added ? "green" : part.removed ? "red" : "black",
                textDecoration: part.added
                  ? "underline"
                  : part.removed
                  ? "line-through"
                  : "none",
              }}
            >
              {part.value}
            </span>
          ))}
        </div>
      </div>
    </div>
  );
};
DiffChecker;

const DiscriptiveExam = (props: Props) => {
  return (
    <div>
      <DiffChecker />
    </div>
  );
};

export default DiscriptiveExam;
