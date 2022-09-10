import { EditorState, Klass, LexicalNode } from "lexical";

import { LexicalComposer } from "@lexical/react/LexicalComposer";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import { OnChangePlugin } from "@lexical/react/LexicalOnChangePlugin";
import { AutoFocusPlugin } from "@lexical/react/LexicalAutoFocusPlugin";
import { ListPlugin } from "@lexical/react/LexicalListPlugin";
import { HeadingNode } from "@lexical/rich-text";
import { HistoryPlugin } from "@lexical/react/LexicalHistoryPlugin";

// When the editor changes, you can get notified via the
// LexicalOnChangePlugin!
function onChange(editorState: EditorState) {
  console.log(editorState);
}

function onError(error: Error) {
  console.error(error);
}

function Editor() {
  const nodes: Klass<LexicalNode>[] = [HeadingNode];

  const initialConfig = {
    namespace: "demo",
    nodes,
    onError,
  };
  return (
    <LexicalComposer initialConfig={initialConfig}>
      <div>
        <ListPlugin />
        <RichTextPlugin
          contentEditable={
            <ContentEditable className="TableNode__contentEditable" />
          }
          placeholder={"something"}
        />
        <AutoFocusPlugin />
        <OnChangePlugin onChange={onChange} />
        <HistoryPlugin />
      </div>
    </LexicalComposer>
  );
}

export default Editor;
