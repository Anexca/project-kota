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
} from "@xyflow/react";

// import {
//   nodes as initialNodes,
//   edges as initialEdges,
// } from "./initial-elements";
import AnnotationNode from "./anotatenode";
import ToolbarNode from "./toolbarnode";
import ResizerNode from "./resizendode";
import CircleNode from "./circlenode";
import TextNode from "./textnode";
import ButtonEdge from "./buttonedge";

import "@xyflow/react/dist/style.css";
import "./overview.css";

const nodeTypes = {
  annotation: AnnotationNode,
  tools: ToolbarNode,
  resizer: ResizerNode,
  circle: CircleNode,
  textinput: TextNode,
};

const edgeTypes = {
  button: ButtonEdge,
};

import { MarkerType } from "@xyflow/react";

const initialNodes = [
  {
    id: "annotation-1",
    type: "annotation",
    draggable: false,
    selectable: false,
    data: {
      level: 1,
      label:
        "Built-in node and edge types. Draggable, deletable and connectable!",
      arrowStyle: {
        right: 0,
        bottom: 0,
        transform: "translate(-30px,10px) rotate(-80deg)",
      },
    },
    position: { x: -80, y: -30 },
  },
  {
    id: "1-1",
    type: "input",
    data: {
      label: "Input Node",
    },
    position: { x: 150, y: 0 },
  },
  {
    id: "1-2",
    type: "default",
    data: {
      label: "Default Node",
    },
    position: { x: 0, y: 100 },
  },
  {
    id: "1-3",
    type: "output",
    data: {
      label: "Output Node",
    },
    position: { x: 300, y: 100 },
  },
  {
    id: "annotation-2",
    type: "annotation",
    draggable: false,
    selectable: false,
    data: {
      level: 2,
      label: "Sub flows, toolbars and resizable nodes!",
      arrowStyle: {
        left: 0,
        bottom: 0,
        transform: "translate(5px, 25px) scale(1, -1) rotate(100deg)",
      },
    },
    position: { x: 220, y: 200 },
  },
  {
    id: "2-1",
    type: "group",
    position: {
      x: -170,
      y: 250,
    },
    style: {
      width: 380,
      height: 180,
      backgroundColor: "rgba(208, 192, 247, 0.2)",
    },
  },
  {
    id: "2-2",
    data: {
      label: "Node with Toolbar",
    },
    type: "tools",
    position: { x: 50, y: 50 },
    style: {
      width: 80,
      height: 80,
      background: "rgb(208, 192, 247)",
    },
    parentId: "2-1",
    extent: "parent",
  },
  {
    id: "2-3",
    type: "resizer",
    data: {
      label: "resizable node",
    },
    position: { x: 250, y: 50 },
    style: {
      width: 80,
      height: 80,
      background: "rgb(208, 192, 247)",
      color: "white",
    },
    parentId: "2-1",
    extent: "parent",
  },
  {
    id: "annotation-3",
    type: "annotation",
    draggable: false,
    selectable: false,
    data: {
      level: 3,
      label: <>Nodes and edges can be anything and are fully customizable!</>,
      arrowStyle: {
        right: 0,
        bottom: 0,
        transform: "translate(-35px, 20px) rotate(-80deg)",
      },
    },
    position: { x: -40, y: 570 },
  },
  {
    id: "3-2",
    type: "textinput",
    position: { x: 150, y: 650 },
    data: {},
  },
  {
    id: "3-1",
    type: "circle",
    position: { x: 350, y: 500 },
    data: {},
  },
];

const initialEdges = [
  {
    id: "e1-2",
    source: "1-1",
    target: "1-2",
    label: "edge",
    type: "smoothstep",
  },
  {
    id: "e1-3",
    source: "1-1",
    target: "1-3",
    animated: true,
    label: "animated edge",
  },
  {
    id: "e2-2",
    source: "1-2",
    target: "2-2",
    type: "smoothstep",
    markerEnd: {
      type: MarkerType.ArrowClosed,
    },
  },
  {
    id: "e2-3",
    source: "2-2",
    target: "2-3",
    type: "smoothstep",
    markerEnd: {
      type: MarkerType.ArrowClosed,
    },
  },
  {
    id: "e3-3",
    source: "2-3",
    sourceHandle: "a",
    target: "3-2",
    type: "button",
    animated: true,
    style: { stroke: "rgb(158, 118, 255)", strokeWidth: 2 },
  },
  {
    id: "e3-4",
    source: "2-3",
    sourceHandle: "b",
    target: "3-1",
    type: "button",
    style: { strokeWidth: 2 },
  },
];

const nodeClassName = (node) => node.type;

const Editor = () => {
  const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
  const [rfInstance, setRfInstance] = useState(null);
  const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges);
  const onConnect = useCallback((params) => {
    setEdges((eds) => addEdge(params, eds));
  }, []);

  const addNewNode = () => {
    setNodes([
      ...nodes,
      {
        id: `1-${nodes.length + 1}`,
        type: "default",
        data: {
          label: "Default Node",
        },
        position: { x: 50, y: 100 },
      },
    ]);
  };
  console.log(rfInstance?.toObject());

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
