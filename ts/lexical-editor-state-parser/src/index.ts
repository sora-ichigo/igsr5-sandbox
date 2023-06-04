import {
  $createParagraphNode,
  $getRoot,
  createEditor,
  EditorState,
} from "lexical";

export const createEmptyLexicalEditorState = () => {
  // NOTE: This is a hack to create an empty editor state.
  // see. https://github.com/facebook/lexical/blob/6935e7e77bd0195ab394259aa4b35203a45adaf1/packages/lexical-website/docs/concepts/editor-state.md?plain=1#L30-L35
  const editor = createEditor();
  const value =
    '{"root":{"children":[{"children":[],"direction":null,"format":"","indent":0,"type":"paragraph","version":1}],"direction":null,"format":"","indent":0,"type":"root","version":1}}';

  return editor.parseEditorState(value);
};
const editorState = createEmptyLexicalEditorState();
console.log(editorState.isEmpty());
const editor = createEditor();
console.log(editor.getEditorState().isEmpty());
