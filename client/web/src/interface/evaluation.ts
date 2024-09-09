type EvaluationPending = {
  id: number;
  completed_seconds: number;
  status: "PENDING";
  created_at: string;
  updated_at: string;
};
type EvaluationRejected = {
  id: number;
  completed_seconds: number;
  status: "REJECTED";
  created_at: string;
  updated_at: string;
  raw_assesment_data?: {
    profane_content: string;
    profanity_check: "detected";
  };
};
type EvaluationCompleted = {
  id: number;
  completed_seconds: number;
  raw_assesment_data: {
    corrected_version: string;
    rating: string;
    strengths: string[];
    weaknesses: string[];
  };
  raw_user_submission: {
    content: string;
  };
  status: "COMPLETED";
  created_at: string;
  updated_at: string;
};
type Evalution = EvaluationPending | EvaluationCompleted | EvaluationRejected;

export type {
  EvaluationCompleted,
  EvaluationPending,
  EvaluationRejected,
  Evalution,
};
