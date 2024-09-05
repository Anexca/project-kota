import { Change, diffWords } from "diff";
import { useEffect, useState } from "react";

const DiffChecker = ({
  originalText,
  modifiedText,
}: {
  originalText: string;
  modifiedText: string;
}) => {
  const [diffResult, setDiffResult] = useState<Change[]>([]);

  const handleDiff = () => {
    // You can choose to use diffChars or diffLines based on your requirement
    const diff = diffWords(originalText, modifiedText);
    setDiffResult(diff);
  };

  useEffect(() => {
    handleDiff();
  }, []);

  return (
    <div className="mt-4">
      <h2 className="mb-2">AI Assesed answer :</h2>
      <p
        style={{ whiteSpace: "pre-wrap" }}
        className="text-sm rounded p-2 border shadow-sm"
      >
        {diffResult.map((part, index) =>
          part.removed || part.added ? (
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
          ) : (
            part.value
          )
        )}
      </p>
    </div>
  );
};
export default DiffChecker;
