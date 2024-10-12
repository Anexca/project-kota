import remarkGfm from "remark-gfm";
import RemarkMathPlugin from "remark-math";
import ReactMarkdown, { Components } from "react-markdown";
import m from "./m.module.scss";
import { MathJax, MathJaxContext } from "better-react-mathjax";

function MarkdownRender({ ...props }) {
  const cp: Components = {
    code: (props) => {
      return (
        <MathJax style={{ display: "inline" }}>\({props.children}\)</MathJax>
      );
    },
    table: (props) => {
      return <span className={m.markdown}>{props.children}</span>;
    },
  };
  const newProps = {
    ...props,
    remarkPlugins: [RemarkMathPlugin, remarkGfm],
    components: cp,
  };
  return (
    <MathJaxContext>
      <ReactMarkdown {...newProps} />
    </MathJaxContext>
  );
}

export default MarkdownRender;
