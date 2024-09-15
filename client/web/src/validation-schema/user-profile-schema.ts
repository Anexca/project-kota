import * as yup from "yup";

const UserProfileSchema = yup.object({
  firstName: yup.string().required("First Name is required."),
  lastName: yup.string().required("Last Name is required."),
  phoneNumber: yup
    .string()
    .matches(
      /^(\+91[\-\s]?)?[6-9]\d{9}$/,
      "Please provide a valid phone number."
    )
    .required("Phone Number is required."),
});

type UserProfileType = yup.InferType<typeof UserProfileSchema>;

export { UserProfileSchema };

export type { UserProfileType };
