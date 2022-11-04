import * as React from "react";
import { Suspense, useMemo, useState } from "react";
import "./App.css";

const App: React.FC = () => {
  return (
    <div className="text-center">
      <h1 className="text-2xl">React App!</h1>
      <Suspense fallback={<p>loading...</p>}>
        <Component />
      </Suspense>

      {[1, 2, 3].map((i) => {
        return <Comment key={1} />;
      })}
    </div>
  );
};

const Comment: React.FC = () => {
  const [comment, setComment] = useState("a");
  const third = "aa";
  const first = "bb";

  React.useEffect(() => {
    const a = comment;
    first;
  }, [third]);

  return (
    <input
      type="text"
      value={comment}
      onChange={(e) => setComment(e.target.value)}
    />
  );
};

export default App;

const sleep = (ms: number) => {
  return new Promise((resolve) => setTimeout(resolve, ms));
};

const Component: React.FC = () => {
  const [data, setData] = useState<string | null>(null);
  if (data === null) {
    throw fetchData1().then(setData);
  }
  const _ = useMemo(() => console.log("loading"), [data]);

  return <p>Data is {data}</p>;
};

async function fetchData1(): Promise<string> {
  await sleep(1000);
  return `Hello, ${(Math.random() * 1000).toFixed(0)}`;
}
