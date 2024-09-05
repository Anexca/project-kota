import { Change, diffWords } from "diff";
import { useEffect, useState } from "react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../../base/tabs";

const DiffChecker = ({
  originalText,
  modifiedText,
  rating,
  strength,
  weaknesses,
}: {
  originalText: string;
  modifiedText: string;
  rating: string;
  weaknesses: string[];
  strength: string[];
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
    <div className="mt-4 border rounded p-2 md:p-4">
      <h2 className="mb-4">
        AI Assesed rating{" "}
        <span className="inline-flex items-center justify-center rounded-full bg-green-100 px-2.5 py-0.5 text-green-700">
          <i className="fa-regular fa-star text-sm mr-2"></i>
          <p className="whitespace-nowrap text-sm"> {rating}</p>
        </span>
      </h2>

      <Tabs defaultValue="assesment" className="w-full">
        <TabsList>
          <TabsTrigger value="assesment">Assesment</TabsTrigger>
          <TabsTrigger value="insights">Insights</TabsTrigger>
        </TabsList>
        <TabsContent value="assesment">
          <DiffResult diffResult={diffResult} />
        </TabsContent>
        <TabsContent value="insights">
          <Insights strength={strength} weaknesses={weaknesses} />
        </TabsContent>
      </Tabs>
    </div>
  );
};
export default DiffChecker;

const Insights = ({
  strength,
  weaknesses,
}: {
  strength: string[];
  weaknesses: string[];
}) => {
  return (
    <div className="flex flex-col gap-2 text-sm">
      <div className="bg-green-100 text-green-700 rounded p-2">
        <label className="font-medium">Strengths :</label>
        <ul className="list-disc ml-4">
          {strength.map((item) => (
            <li>{item}</li>
          ))}
        </ul>
      </div>
      <div className="bg-orange-100 text-orange-700 rounded p-2">
        <label className="font-medium">Weakness :</label>
        <ul className="list-disc ml-4">
          {weaknesses.map((item) => (
            <li>{item}</li>
          ))}
        </ul>
      </div>
    </div>
  );
};

const DiffResult = ({ diffResult }: { diffResult: Change[] }) => {
  return (
    <p
      style={{ whiteSpace: "pre-wrap" }}
      className="text-sm rounded p-2  bg-neutral-100/50"
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
  );
};
