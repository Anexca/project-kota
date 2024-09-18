interface RawExamData {
  hints: string[];
  topic: string;
  type: string;
  max_number_of_words: string;
  total_marks: string;
}

export interface IQuestion {
  id: number;
  raw_exam_data: RawExamData;
  user_attempts: number;
  max_attempts: number;
  duration_seconds: number;
  number_of_questions: number;
  created_at: string;
  updated_at: string;
}

export const typeOfExam = ["DESCRIPTIVE"] as const;
export const categoryOfExam = ["BANKING"] as const;
export interface ICategory {
  exam_type_id: number;
  exam_name: string;
  description: string;
  type_of_exam: (typeof typeOfExam)[number];
  is_active: boolean;
  category_name: (typeof categoryOfExam)[number];
  category_id: number;
}
