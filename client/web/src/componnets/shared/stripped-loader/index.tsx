import { cn } from "../../../lib/utils";
import styles from "./loader.module.scss";
type Props = {
  className?: string;
};

const StrippedLoader = ({ className }: Props) => {
  return <div className={cn(styles.animate_stripes, className)}></div>;
};

export default StrippedLoader;
