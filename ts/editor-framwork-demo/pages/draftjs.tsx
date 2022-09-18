import { NextPage } from "next";
import { useState } from "react";
import { convertFromRaw, Editor, EditorState } from "draft-js";
import "draft-js/dist/Draft.css";

const emptyContentState = convertFromRaw({
  entityMap: {},
  blocks: [
    {
      key: "foo",
      type: "unstyled",
      text: "",
      depth: 0,
      inlineStyleRanges: [],
      entityRanges: [],
    },
  ],
});

const DraftJS: NextPage = () => {
  const [editorState, setEditorState] = useState(() =>
    EditorState.createWithContent(emptyContentState)
  );
  return (
    <div style={{ width: "100%", height: "300px" }}>
      <Editor
        editorState={editorState}
        onChange={(v) => {
          setEditorState(v);
          console.log(editorState.getCurrentContent());
        }}
      />
    </div>
  );
};

export default DraftJS;
