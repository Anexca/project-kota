export interface IMCQSubmission {
  id: number;
  completed_seconds: number;
  obtained_marks: number;
  total_marks: number;
  cutoff_marks: number;
  status: string;
  raw_assesment_data: RawAssesmentData;
  raw_user_submission: RawUserSubmission;
  created_at: string;
  updated_at: string;
}

export interface RawAssesmentData {
  summary: Summary;
}

export interface Summary {
  accuracy: number;
  attempted: number;
  correct: number;
  incorrect: number;
}

export interface RawUserSubmission {
  attempted_questions: AttemptedQuestion[];
}

export interface AttemptedQuestion {
  question_number: number;
  time_taken_in_seconds: number;
  user_selected_option_index: number[];
}
