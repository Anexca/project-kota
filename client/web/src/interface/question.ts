interface RawExamData {
  hints: string[];
  topic: string;
  type: string;
  max_number_of_words: string;
  total_marks: string;
}

export interface IQuestion {
  exam_id: number;
  raw_exam_data: RawExamData;
  user_attempts: number;
  max_attempts: number;
  duration_seconds: number;
  number_of_questions: number;
  created_at: string;
  updated_at: string;
  exam_type: (typeof typeOfExam)[0];
}

export const typeOfExam = ["DESCRIPTIVE", "MCQ"] as const;
export const categoryOfExam = ["BANKING"] as const;
export interface ICategory {
  exam_type_id: number;
  exam_name: string;
  description: string;
  type_of_exam: (typeof typeOfExam)[number];
  is_active: boolean;
  category_name: (typeof categoryOfExam)[number];
  category_id: number;
  logo_url: string;
}

export interface IMCQExam {
  exam_id: number;
  exam_name: string;
  user_attempts: number;
  max_attempts: number;
  duration_seconds: number;
  number_of_questions: number;
  negative_marking: number;
  created_at: string;
  updated_at: string;
  exam_type: (typeof typeOfExam)[1];
  exam_stage: string;
}

export interface IMCQQuestionSet {
  id: number;
  exam_type: string;
  exam_name: string;
  raw_exam_data: {
    questions: IMCQQuestion[];
    content_groups: IContentGroup[];
  };
  user_attempts: number;
  max_attempts: number;
  duration_seconds: number;
  number_of_questions: number;
  negative_marking: number;
  created_at: string;
  updated_at: string;
}
export interface IMCQQuestion {
  answer: string;
  content_reference_id: string;
  explanation: string;
  options: string[];
  question: string;
  question_number: number;
}
export interface IContentGroup {
  content: string;
  content_id: string;
  instructions: string;
}
