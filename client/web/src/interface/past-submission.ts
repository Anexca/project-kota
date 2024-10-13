export interface Attempt {
  attempt_id: number;
  attempt_number: number;
  assessment_id: number;
  attempt_date: string;
}

export interface IPastExamAttempt {
  attempted_exam_id: number;
  is_active: boolean;
  exam_type: string;
  exam_type_id: number;
  exam_category: string;
  exam_category_id: number;
  topic: string;
  type: string;
  attempts: Attempt[];

  exam_group: string;
  exam_group_id: number;
  exam_name: string;

  exam_stage: string;
  is_sectional: boolean;
}
