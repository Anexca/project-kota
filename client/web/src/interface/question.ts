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
