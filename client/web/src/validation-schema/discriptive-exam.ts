import * as yup from "yup";

const DiscriptiveExamSchema = yup.object({
  answer: yup.string().length(250).required("Answer is required."),
});
type DiscriptiveExamSchemaType = yup.InferType<typeof DiscriptiveExamSchema>;
export { DiscriptiveExamSchema };

export type { DiscriptiveExamSchemaType };
