import { QUESTION_STATE } from "../constants/shared";

export type MCQAnswersModel = {
  state: `${QUESTION_STATE}`;
  selectedOption: null | number;
  questionNumber: number;
};
export type MCQFormModal = {
  activeQuestionIndex: number;
  activeSection: string;
  answers: Record<string, MCQAnswersModel[]>;
};
