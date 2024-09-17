import { SelectSingleEventHandler } from "react-day-picker";
import { Button } from "../../base/button/button";
import { Popover, PopoverContent, PopoverTrigger } from "../../base/popper";
import { Calendar } from "../date-picker";
import { format } from "date-fns";
import Icon from "../../base/icon";
import { BaseSyntheticEvent } from "react";
type DateFilterSectionProps = {
  date?: Date;
  updateDate: SelectSingleEventHandler;
  label: string;
  toDate?: Date;
  fromDate?: Date;
};
export function DateFilterSection({
  date,
  updateDate,
  label,
  fromDate,
  toDate,
}: DateFilterSectionProps) {
  const clearDate = (e: BaseSyntheticEvent) => {
    e.stopPropagation();
    ///@ts-ignore
    updateDate(undefined);
  };
  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button variant="outline" className="h-8 flex items-center gap-2">
          {label}{" "}
          {date ? (
            <span>
              {format(date, "PPP")}
              <Button onClick={clearDate} variant={"ghost"} className="p-0">
                <Icon icon="xmark" className="ml-4" />
              </Button>
            </span>
          ) : (
            <span>Pick a date</span>
          )}
        </Button>
      </PopoverTrigger>
      <PopoverContent align="start" className="w-auto">
        <div>
          <Calendar
            required
            mode="single"
            selected={date}
            onSelect={updateDate}
            className=""
            toDate={toDate}
            fromDate={fromDate}
          />
        </div>
      </PopoverContent>
    </Popover>
  );
}
