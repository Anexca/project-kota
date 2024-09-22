"use client";
import React, { useCallback, useState } from "react";

import {
  ReactFlow,
  addEdge,
  MiniMap,
  Controls,
  Background,
  useNodesState,
  useEdgesState,
  Node,
  Edge,
  Connection,
} from "@xyflow/react";

import "@xyflow/react/dist/style.css";
import "./overview.css";
import Topic from "./components/topic";
import DashedEdge from "./components/dashed-edge";

const nodeClassName = (node) => node.type;

const nodeTypes = {
  topic: Topic,
};
const edgeTypes = {
  dashededge: DashedEdge,
};
const Editor = () => {
  const [nodes, setNodes, onNodesChange] = useNodesState<Node>([]);
  const [rfInstance, setRfInstance] = useState(null);
  const [edges, setEdges, onEdgesChange] = useEdgesState<Edge>([]);
  const onConnect = useCallback((params: Connection) => {
    setEdges((eds) => addEdge(params, eds));
  }, []);

  const addNewNode = () => {
    setNodes([
      ...nodes,
      {
        id: `1-${nodes.length + 1}`,
        type: "topic",
        data: {
          label: "Default Node",
        },
        position: { x: 50, y: 100 },
      },
    ]);
  };

  return (
    <div style={{ width: "100%", height: "100vh", position: "relative" }}>
      <button onClick={addNewNode}>Add Node</button>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        onConnect={onConnect}
        fitView
        attributionPosition="top-right"
        nodeTypes={nodeTypes}
        edgeTypes={edgeTypes}
        className="overview"
        snapToGrid={true}
        onInit={setRfInstance}
      >
        <MiniMap zoomable pannable nodeClassName={nodeClassName} />
        <Controls />
        <Background />
      </ReactFlow>
    </div>
  );
};

export default Editor;
