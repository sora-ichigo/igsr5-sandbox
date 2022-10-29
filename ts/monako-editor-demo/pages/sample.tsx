import { NextPage } from "next";
import React, { useEffect } from "react";
import Editor, { useMonaco } from "@monaco-editor/react";

const SampleEditor: NextPage = () => {
  return <MonacoEditor />;
};
export default SampleEditor;

function MonacoEditor() {
  const monaco = useMonaco();

  useEffect(() => {
    console.log("monaco", monaco);
  }, [monaco]);

  return (
    <Editor
      height="90vh"
      defaultLanguage="javascript"
      defaultValue="// some comment"
    />
  );
}
