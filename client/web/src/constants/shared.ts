export enum QUESTION_STATE {
  ATTEMPED = "ATTEMPED",
  NOT_ATTEMPED = "NOT-ATTEMPED",
  FOR_REVIEW = "FOR-REVIEW",
  NOT_ANSWERED = "NOT-ANSWERED",
}

export const questionType: Record<string, string> = {
  formal_letter: "Formal Letter",
  essay: "Essay",
};
