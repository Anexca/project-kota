import { useEffect, useState } from "react";
import { getPastAttemptedSubmissions } from "../../services/exam.service";

import { Link } from "react-router-dom";
import { paths } from "../../routes/route.constant";

import Icon from "../../componnets/base/icon";
import AttemptedSubmissionsCard from "../../componnets/shared/attempted-submissions-card";
import { DateFilterSection } from "../../componnets/shared/filter-section";
import { useToast } from "../../hooks/use-toast";
import { IPastExamAttempt } from "../../interface/past-submission";
import { PaginationComponent } from "../../componnets/shared/pagination";
import { DateFilterType, IPaginationType } from "../../interface/utils";

const paginationInit = {
  current_page: 1,
  total_items: 0,
  per_page: 10,
  total_pages: 1,
};
const PreviousSubmissionPage = () => {
  const [dateFilter, setDateFilters] = useState<DateFilterType>({});
  const [pagination, setPagination] = useState<IPaginationType>(paginationInit);
  const [submissions, setQuestions] = useState<IPastExamAttempt[]>([]);
  const [loading, setLoading] = useState(false);
  const { toast } = useToast();
  const getSubmissionList = async (
    dateFilter: DateFilterType,
    pagination: IPaginationType
  ) => {
    setLoading(true);
    try {
      const data = await getPastAttemptedSubmissions({
        from: dateFilter?.from?.toISOString(),
        to: dateFilter?.to?.toISOString(),
        limit: pagination.per_page,
        page: pagination.current_page,
      });
      setQuestions(data.data || []);
      setPagination(data.pagination);
    } catch (error) {
      toast({
        title: "Oh ho Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in processing your request.",
      });
    }
    setLoading(false);
  };
  useEffect(() => {
    getSubmissionList(dateFilter, pagination);
  }, []);

  const updateDateFilter = (key: string) => (date?: Date) => {
    const newDate = { ...dateFilter, [key]: date };
    setDateFilters(newDate);
    setPagination(paginationInit);
    getSubmissionList(newDate, pagination);
  };
  const updatePagination = (pageNumber: number) => {
    const newPagination = { ...pagination, current_page: pageNumber };
    setPagination(newPagination);
    getSubmissionList(dateFilter, newPagination);
  };
  return (
    <div className="pt-2 w-full md:max-w-2xl 2xl:max-w-2xl mx-auto flex flex-col gap-2 p-4">
      <div className="py-2">
        <div className="flex gap-2 items-center">
          <Link to={`/${paths.HOMEPAGE}`} className="p-0">
            <Icon icon="arrow_back" className="text-info text-lg" />
          </Link>
          <span className="text-lg font-semibold">
            All of your previously attempted submissions.
          </span>
        </div>
        <div className="relative flex items-center gap-2 text-sm flex-wrap mt-2">
          <DateFilterSection
            label="From -"
            date={dateFilter.from}
            updateDate={updateDateFilter("from")}
            toDate={dateFilter.to}
          />
          <DateFilterSection
            label="To -"
            date={dateFilter.to}
            updateDate={updateDateFilter("to")}
            fromDate={dateFilter.from}
          />
        </div>
      </div>
      {loading ? (
        <div className="flex flex-col gap-2 justify-center items-center">
          <span className="rounded-full w-8 h-8 animate-spin border-2 border-info border-t-info/30"></span>
          Getting Your Submissions
        </div>
      ) : (
        <div className="animate-fadeIn flex flex-col gap-2">
          {submissions.length ? (
            submissions.map((item) => {
              return item.attempts.length ? (
                <AttemptedSubmissionsCard
                  key={item.attempted_exam_id}
                  topic={item.topic}
                  type={item.type}
                  srNumber={item.attempted_exam_id}
                  category={item.exam_category}
                  submissions={item.attempts}
                />
              ) : null;
            })
          ) : (
            <div className="flex flex-col gap-2 justify-center items-center">
              You don't have any submissions yet. Please attempt some exams.
            </div>
          )}
        </div>
      )}
      <div className="mt-4">
        <PaginationComponent
          currentPage={pagination.current_page}
          onChange={updatePagination}
          totalPage={pagination.total_pages}
        />
      </div>
    </div>
  );
};

export default PreviousSubmissionPage;
