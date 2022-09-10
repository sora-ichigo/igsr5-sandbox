import {
  $createParagraphNode,
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
import { ListPlugin } from "@lexical/react/LexicalListPlugin";
import { CheckListPlugin } from "@lexical/react/LexicalCheckListPlugin";
import { HistoryPlugin } from "@lexical/react/LexicalHistoryPlugin";
import { useCallback, useState } from "react";
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import { $wrapLeafNodesInElements } from "@lexical/selection";
import {
  HeadingNode,
  QuoteNode,
  HeadingTagType,
  $createHeadingNode,
  $createQuoteNode,
} from "@lexical/rich-text";
import {
  ListItemNode,
  ListNode,
  INSERT_ORDERED_LIST_COMMAND,
  INSERT_UNORDERED_LIST_COMMAND,
  INSERT_CHECK_LIST_COMMAND,
} from "@lexical/list";

// When the editor changes, you can get notified via the
// LexicalOnChangePlugin!
function onChange(editorState: EditorState) {
  console.log(editorState);
}

function onError(error: Error) {
  console.error(error);
}

function Editor() {
  const nodes: Klass<LexicalNode>[] = [
    HeadingNode,
    QuoteNode,
    ListItemNode,
    ListNode,
  ];

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
        <ListPlugin />
        <CheckListPlugin />
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
    quote: "Quate",
    number: "Numbered List",
    bullet: "Bullet List",
    check: "Check List",
  } as const;
  type BlockType = keyof typeof SupportedBlockType;

  const [blockType, setBlockType] = useState<BlockType>("paragraph");
  const [editor] = useLexicalComposerContext();

  const formatParagraph = useCallback(() => {
    if (blockType !== "paragraph") {
      editor.update(() => {
        const selection = $getSelection();
        if ($isRangeSelection(selection)) {
          $wrapLeafNodesInElements(selection, () => $createParagraphNode());
        }
      });

      setBlockType("paragraph");
    }
  }, [blockType, editor]);

  const formatQuote = useCallback(() => {
    if (blockType !== "quote") {
      editor.update(() => {
        const selection = $getSelection();
        if ($isRangeSelection(selection)) {
          $wrapLeafNodesInElements(selection, () => $createQuoteNode());
        }
      });

      setBlockType("quote");
    }
  }, [blockType, editor]);

  const formatBulletList = useCallback(() => {
    if (blockType !== "bullet") {
      editor.update(() => {
        const selection = $getSelection();
        if ($isRangeSelection(selection)) {
          editor.dispatchCommand(INSERT_UNORDERED_LIST_COMMAND, undefined);
        }
      });
      setBlockType("bullet");
    }
  }, [blockType, editor]);

  const formatNumberList = useCallback(() => {
    if (blockType !== "number") {
      editor.update(() => {
        const selection = $getSelection();
        if ($isRangeSelection(selection)) {
          editor.dispatchCommand(INSERT_ORDERED_LIST_COMMAND, undefined);
        }
      });
      setBlockType("number");
    }
  }, [blockType, editor]);

  const formatCheckList = useCallback(() => {
    if (blockType !== "check") {
      editor.update(() => {
        const selection = $getSelection();
        if ($isRangeSelection(selection)) {
          editor.dispatchCommand(INSERT_CHECK_LIST_COMMAND, undefined);
        }
      });
      setBlockType("check");
    }
  }, [blockType, editor]);

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
      <button type="button" onClick={() => formatParagraph()}>
        p
      </button>
      <button type="button" onClick={() => formatQuote()}>
        quote
      </button>
      <button type="button" onClick={() => formatBulletList()}>
        list
      </button>
      <button type="button" onClick={() => formatNumberList()}>
        number
      </button>
      <button type="button" onClick={() => formatCheckList()}>
        check
      </button>
    </div>
  );
}

export default Editor;
