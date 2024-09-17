import { useState } from "react";

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

export function PaginationComponent({}) {
  const [page, onChange] = useState(1);
  const pagination = usePagination({ total: 10, page, onChange });

  return (
    <Pagination>
      <PaginationContent>
        <PaginationItem>
          <PaginationPrevious />
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
              <PaginationLink onClick={() => pagination.setPage(item)}>
                {item}
              </PaginationLink>
            </PaginationItem>
          );
        })}
        <PaginationItem>
          <PaginationNext />
        </PaginationItem>
      </PaginationContent>
    </Pagination>
  );
}
