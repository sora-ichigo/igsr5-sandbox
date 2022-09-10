import { NextPage } from "next";
import { useState } from "react";
import { Editor, EditorState } from "draft-js";
import "draft-js/dist/Draft.css";

const DraftJS: NextPage = () => {
  const [editorState, setEditorState] = useState(() =>
    EditorState.createEmpty()
  );
  return (
    <div style={{ width: "100%", height: "300px" }}>
      <Editor editorState={editorState} onChange={(v) => setEditorState(v)} />
    </div>
  );
};

export default DraftJS;
