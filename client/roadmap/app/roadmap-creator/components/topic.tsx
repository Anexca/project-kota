"use client";
import { Handle, Position } from "@xyflow/react";
import { memo } from "react";

function TopicNode({ data }) {
  return (
    <>
      <div style={{ padding: 10, display: "flex" }}>
        <div>{data.label}</div>
        <Handle
          style={{ position: "relative", left: 0, transform: "none" }}
          id="a"
          type="source"
          position={Position.Bottom}
        />
        <Handle
          style={{ position: "relative", left: 0, transform: "none" }}
          id="b"
          type="source"
          position={Position.Bottom}
        />
      </div>
    </>
  );
}

export default memo(TopicNode);
