export interface FilterPagination {
  from?: string;
  to?: string;
  page?: number;
  limit?: number;
}
export type DateFilterType = {
  from?: Date;
  to?: Date;
};
export type IPaginationType = {
  current_page: number;
  total_pages: number;
  per_page: number;
  total_items: number;
};
