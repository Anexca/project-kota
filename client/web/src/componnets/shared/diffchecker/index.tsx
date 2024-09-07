import { Change, diffWords } from "diff";
import { useEffect, useState } from "react";
import Icon from "../../base/icon";
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
    <div className="mt-4 shadow-sm rounded p-2 md:p-4 bg-white text-sm">
      <h2 className="mb-4">
        AI Assessment rating{" "}
        <span className="ml-2 inline-flex items-center justify-center rounded-full bg-green-100 px-2.5 py-0.5 text-green-700">
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
      <div className=" rounded p-2 shadow-sm">
        <label className="font-medium underline">
          <Icon icon="dumbbell" className="text-green-700 text-base mr-2" />
          Strengths
        </label>
        <ul className="pl-1 mt-2 space-y-2">
          {strength.map((item) => (
            <li className="flex items-start">
              <Icon icon="caret_up" className="text-green-700 text-base mr-2" />
              <span>{item}</span>
            </li>
          ))}
        </ul>
      </div>
      <div className="rounded p-2 shadow-sm">
        <label className="font-medium underline">
          <Icon
            icon="exclaimation"
            className="text-orange-700 text-base mr-2"
          />
          Weakness
        </label>
        <ul className="pl-1 mt-2 space-y-2">
          {weaknesses.map((item) => (
            <li className="flex items-start">
              <Icon
                icon="caret_down"
                className="text-orange-700 text-base mr-2"
              />
              <span>{item}</span>
            </li>
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
      className="text-sm rounded p-2 px-3  bg-neutral-100/25 shadow-sm"
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
