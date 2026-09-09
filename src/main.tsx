import React from "react";
import { createRoot } from "react-dom/client";

function App() { return <main><h1>业务工作台</h1><p>领域功能将在此接入。</p></main>; }
createRoot(document.getElementById("root")!).render(<App />);
