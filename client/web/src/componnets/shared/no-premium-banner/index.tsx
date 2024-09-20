import { paths } from "../../../routes/route.constant";
import Icon from "../../base/icon";
import { StyledLink } from "../../base/styled-link";

const NoPremiumBanner = () => {
  return (
    <div className="flex p-4 rounded bg-white shadow-sm mt-2 gap-4">
      <div className=" w-2 bg-yellow-400 rounded-full"></div>
      <div className="flex flex-col flex-1">
        <div className="text-sm font-semibold">
          Please buy one of our paid plan to get new question daily.
        </div>
        <div>
          <StyledLink
            to={`/${paths.PRICING_PLAN}`}
            size={"sm"}
            className="px-3 py-1 h-7 mt-2"
            variant={"warning"}
          >
            <Icon icon="send" className="mr-2" /> Buy Plan
          </StyledLink>
        </div>
      </div>
    </div>
  );
};

export default NoPremiumBanner;
