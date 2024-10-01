import { QUESTION_STATE } from "../constants/shared";

export type MCQAnswersModel = {
  state: `${QUESTION_STATE}`;
  selectedOption: null | number;
};
export type MCQFormModal = {
  activeQuestionIndex: number;
  answers: MCQAnswersModel[];
};
