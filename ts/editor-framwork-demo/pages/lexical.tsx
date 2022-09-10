import {
  $getSelection,
  $isRangeSelection,
  EditorState,
  Klass,
  LexicalNode,
} from "lexical";

import { LexicalComposer } from "@lexical/react/LexicalComposer";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import { OnChangePlugin } from "@lexical/react/LexicalOnChangePlugin";
import { AutoFocusPlugin } from "@lexical/react/LexicalAutoFocusPlugin";
import { HeadingNode } from "@lexical/rich-text";
import { HistoryPlugin } from "@lexical/react/LexicalHistoryPlugin";
import { useCallback, useState } from "react";
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import { $wrapLeafNodesInElements } from "@lexical/selection";
import { HeadingTagType, $createHeadingNode } from "@lexical/rich-text";

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
      <ToolbarPlugin />

      <div>
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

function ToolbarPlugin() {
  const SupportedBlockType = {
    paragraph: "Paragraph",
    h1: "Heading 1",
    h2: "Heading 2",
    h3: "Heading 3",
    h4: "Heading 4",
    h5: "Heading 5",
    h6: "Heading 6",
  } as const;
  type BlockType = keyof typeof SupportedBlockType;

  const [blockType, setBlockType] = useState<BlockType>("paragraph");
  const [editor] = useLexicalComposerContext();
  const formatHeading = useCallback(
    (type: HeadingTagType) => {
      if (blockType !== type) {
        editor.update(() => {
          const selection = $getSelection();
          if ($isRangeSelection(selection)) {
            $wrapLeafNodesInElements(selection, () => $createHeadingNode(type));
          }
        });
        setBlockType(type);
      }
    },
    [blockType, editor]
  );

  return (
    <div>
      <button type="button" onClick={() => formatHeading("h1")}>
        h1
      </button>
      <button type="button" onClick={() => formatHeading("h2")}>
        h2
      </button>
      <button type="button" onClick={() => formatHeading("h3")}>
        h3
      </button>
    </div>
  );
}

export default Editor;
