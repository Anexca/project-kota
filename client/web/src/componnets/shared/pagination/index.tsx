import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "../../base/pagination";
import { usePagination } from "../../../hooks/use-pagination";

type PaginationComponent = {
  totalPage: number;
  currentPage: number;
  onChange: (p: number) => void;
};
export function PaginationComponent({
  currentPage,
  onChange,
  totalPage,
}: PaginationComponent) {
  const pagination = usePagination({
    total: totalPage,
    page: currentPage,
    onChange,
  });

  return (
    <Pagination>
      <PaginationContent>
        <PaginationItem>
          <PaginationPrevious onClick={() => pagination.previous()} />
        </PaginationItem>

        {pagination.range.map((item) => {
          if (item == "dots") {
            return (
              <PaginationItem>
                <PaginationEllipsis />
              </PaginationItem>
            );
          }
          return (
            <PaginationItem>
              <PaginationLink
                isActive={item == currentPage}
                onClick={() =>
                  !(item == currentPage) && pagination.setPage(item)
                }
              >
                {item}
              </PaginationLink>
            </PaginationItem>
          );
        })}
        <PaginationItem>
          <PaginationNext onClick={() => pagination.next()} />
        </PaginationItem>
      </PaginationContent>
    </Pagination>
  );
}
