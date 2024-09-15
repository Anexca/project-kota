import { useState } from "react";
import { questionType } from "../../../constants/shared";
import { Button } from "../../base/button/button";
import Chip from "../../base/chip";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "../../base/collapsible";
import Icon from "../../base/icon";
import { cn } from "../../../lib/utils";
import { Attempt } from "../../../interface/past-submission";
import { StyledLink } from "../../base/styled-link";
import { paths } from "../../../routes/route.constant";

type Props = {
  srNumber?: number | null;
  topic: string;
  type: string;
  category: string;
  submissions: Attempt[];
};

const iconSet = {
  COMPLETED: (
    <Icon
      icon="check_solid"
      className="text-green-700 text-xl mr-4 self-start md:self-auto"
    />
  ),
  PENDING: (
    <Icon
      icon="exclaimation_circle"
      className="text-yellow-700 text-xl mr-4 self-start md:self-auto"
    />
  ),
  REJECTED: (
    <Icon
      icon="xmark_circle"
      className="text-destructive text-xl mr-4 self-start md:self-auto"
    />
  ),
};
const AttemptedSubmissionsCard = ({
  srNumber = null,
  topic,
  type,
  category,
  submissions,
}: Props) => {
  const [listOpen, setListOpen] = useState(false);
  return (
    <article className="rounded-md shadow-sm bg-white flex flex-col gap-1 p-3 px-4 md:p3 text-sm">
      <div className="flex flex-col md:flex-row gap-4">
        <div className="flex-1">
          <p className="font-medium text-balance md:text-pretty text-black mb-2">
            #{srNumber} - {topic}
          </p>
        </div>
        <div className="flex flex-col items-stretch gap-2 md:justify-center md:w-32">
          <div className="flex gap-2 flex-wrap ">
            <Chip icon="tags" variant={"info"} className="capitalize">
              {category}
            </Chip>
            <Chip icon="file" variant={"success"}>
              {questionType[type] || "--"}
            </Chip>
          </div>
        </div>
      </div>
      <Collapsible
        open={listOpen}
        onOpenChange={setListOpen}
        className="w-full space-y-2"
      >
        <CollapsibleTrigger asChild>
          <div className="flex items-center justify-between space-x-4 cursor-pointer">
            <>
              <h4 className="text-sm font-semibold">Submissions</h4>
              <Button variant="ghost" size="sm" className="w-9 p-0">
                <Icon
                  icon="chevron_down"
                  className={cn(
                    "h-4 w-4 transition-transform",
                    listOpen && "rotate-180"
                  )}
                />
                <span className="sr-only">Toggle</span>
              </Button>
            </>
          </div>
        </CollapsibleTrigger>

        <CollapsibleContent className="space-y-2">
          {submissions.map((item) => (
            <article className="rounded-md shadow bg-white flex justify-between gap-4 p-3 px-4 md:p3 text-sm">
              <div className="flex flex-1 items-center text-sm text-black">
                <div>
                  <div className="mb-2 font-medium flex items-center">
                    {iconSet.COMPLETED}
                    Attempt: {item.attempt_number}
                  </div>
                  <div className="flex gap-2 flex-wrap">
                    <Chip icon="calender_solid" variant={"info"}>
                      {new Date(item.attempt_date).toLocaleString()}
                    </Chip>
                  </div>
                </div>
              </div>

              <StyledLink
                size="sm"
                className="text-sm py-1"
                variant={"secondary"}
                to={`/${paths.EXAMS}/${paths.MY_SUMBISSIONS}/${srNumber}/${paths.SUBMISSION}/${item.assessment_id}`}
              >
                <Icon icon="send" className="mr-2" /> View
              </StyledLink>
            </article>
          ))}
        </CollapsibleContent>
      </Collapsible>
    </article>
  );
};

export default AttemptedSubmissionsCard;
