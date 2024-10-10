export enum QUESTION_STATE {
  ATTEMPTED = "ATTEMPTED",
  UN_ATTEMPTED = "UN-ATTEMPTED",
  FOR_REVIEW = "FOR-REVIEW",
  NOT_ANSWERED = "NOT-ANSWERED",
}

export const questionType: Record<string, string> = {
  formal_letter: "Formal Letter",
  essay: "Essay",
  precis: "Precis",
};

export const ScreenSizeQuery = {
  smallScreen: "(min-width: 640px)",
  mediumScreen: "(min-width: 768px)",
  largeScreen: "(min-width: 1024px)",
};
