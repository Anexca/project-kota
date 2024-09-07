type EvaluationPending = {
  id: number;
  completed_seconds: number;
  status: "PENDING";
  created_at: string;
  updated_at: string;
};
type EvaluationCompleted = {
  id: number;
  completed_seconds: number;
  raw_assesment_data: {
    corrected_version: string;
    rating: string;
    strengths: string[];
    weakness: string[];
  };
  raw_user_submission: {
    content: string;
  };
  status: "COMPLETED";
  created_at: string;
  updated_at: string;
};

export type { EvaluationCompleted, EvaluationPending };
