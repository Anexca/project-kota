import { paths } from "../../routes/route.constant";
import DescriptiveSubmission from "../descriptive-submission";

const ViewPastSubmission = () => {
  return (
    <DescriptiveSubmission
      backLink={`/${paths.EXAMS}/${paths.MY_SUMBISSIONS}`}
    />
  );
};

export default ViewPastSubmission;
