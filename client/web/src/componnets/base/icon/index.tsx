import { ICONS, IconType } from "../../../constants/icons";
import { cn } from "../../../lib/utils";

type Props = { icon: IconType; className?: string };

const Icon = ({ icon, className }: Props) => {
  return <i className={cn(ICONS[icon], className)}></i>;
};

export default Icon;
