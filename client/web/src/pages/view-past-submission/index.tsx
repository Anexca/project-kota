import { paths } from "../../routes/route.constant";
import DescriptiveSubmission from "../descriptive-submission";

const ViewPastSubmission = ({ isOpenMode }: { isOpenMode?: boolean }) => {
  return (
    <DescriptiveSubmission
      isOpenMode={isOpenMode}
      backLink={`/${isOpenMode ? paths.COMMUNITY_EXAMS : paths.EXAMS}/${
        paths.MY_SUMBISSIONS
      }`}
    />
  );
};

export default ViewPastSubmission;
