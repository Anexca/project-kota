import * as yup from "yup";

const LoginSchema = yup.object({
  email: yup.string().email("Invalid email.").required("Email is required."),
  password: yup.string().required("Password is required."),
});
const ForgotPasswordSchema = yup.object({
  email: yup.string().email("Invalid email.").required("Email is required."),
});
type LoginType = yup.InferType<typeof LoginSchema>;
type ForgotPasswordType = yup.InferType<typeof ForgotPasswordSchema>;
export { LoginSchema, ForgotPasswordSchema };

export type { LoginType, ForgotPasswordType };
