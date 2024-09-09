import * as yup from "yup";

const DiscriptiveExamSchema = yup.object({
  maxWords: yup.number(),
  answer: yup
    .string()
    .test(
      "word-count",
      "Answer can't be more than 250 words.",
      function (value) {
        console.log(this.parent);
        const wordCount = value?.match(/\b\w+(?:[.,!;?])?\b/g)?.length || 0;
        return wordCount < 250;
      }
    )
    .required("Answer is required."),
});
type DiscriptiveExamSchemaType = yup.InferType<typeof DiscriptiveExamSchema>;
export { DiscriptiveExamSchema };

export type { DiscriptiveExamSchemaType };
