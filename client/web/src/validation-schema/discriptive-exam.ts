import * as yup from "yup";

const DiscriptiveExamSchema = yup.object({
  answer: yup
    .string()
    .max(250, "Answer should not be more that 250 words.")
    .required("Answer is required."),
});
type DiscriptiveExamSchemaType = yup.InferType<typeof DiscriptiveExamSchema>;
export { DiscriptiveExamSchema };

export type { DiscriptiveExamSchemaType };
