import * as React from "react";
import { Suspense } from "react";
import "./App.css";

const App: React.FC = () => {
  return (
    <div className="text-center">
      <h1 className="text-2xl">React App!</h1>
      <Suspense fallback={<p>loading...</p>}>
        <p>test</p>
        <Component />
      </Suspense>
    </div>
  );
};

export default App;

const sleep = (ms: number) => {
  return new Promise((resolve) => setTimeout(resolve, ms));
};

const Component: React.FC = () => {
  throw sleep(1000);
};
