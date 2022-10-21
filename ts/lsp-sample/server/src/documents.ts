import { TextDocument } from 'vscode-languageserver-textdocument';
import { TextDocumentChangeEvent, TextDocuments } from 'vscode-languageserver/node';
import { getConnection } from './connections';
import { Diagnostics } from './diagnostics';

// eslint-disable-next-line @typescript-eslint/ban-types
type CompileFiles = Map<string, Object>;

const compile = (uri: string, src: string) => {
	diagnostics.clear();
	const tokens = tokenize(uri, src);

	buffers.set(uri, tokens);
	return buffers;
};

// eslint-disable-next-line @typescript-eslint/ban-types
const tokenize = (uri: string, src: string): Object[] => {
	let i = 0;
	let line = 0;
	let character = 0;
	// eslint-disable-next-line @typescript-eslint/ban-types
	const tokens: Object[] = [];

	const nextChar = () => {
		if (i >= src.length) return;
		if (src[i] === "\n") {
			++i;
			++line;
			character = 0;
		} else {
			++i;
			++character;
		}
	};

	while (true) {
		// skip leading whitespaces
		while (true) {
			if (i >= src.length) return tokens;
			if (" \n\t\r".indexOf(src[i]) === -1) break;
			nextChar();
		}

		const start: { line: number, character: number } = { line, character };

		let text;
		let kind;

		if (src[i] === "(") {
			text = "(";
			kind = "(";
			nextChar();
		} else if (src[i] === ")") {
			text = ")";
			kind = ")";
			nextChar();
		} else if (src[i] === ";") {
			const begin = i;
			while (true) {
				if (i >= src.length) break;
				if (src[i] === "\n") break;
				nextChar();
			}
			text = src.substring(begin, i);
			kind = "comment";
		} else {
			const begin = i;
			while (true) {
				if (i >= src.length) break;
				if (" \n\t\r();".indexOf(src[i]) !== -1) break;
				nextChar();
			}

			text = src.substring(begin, i);

			if (!isNaN(Number(text))) {
				kind = "number";
			} else {
				kind = "variable";
			}
		}

		const end: { line: number, character: number } = { line, character };
		const location = { uri, range: { start, end } };
		tokens.push({ text, kind, location });
	}

	return tokens;
};


const buffers: CompileFiles = new Map();
const diagnostics: Diagnostics = new Map();

export class DocumentManager {
	public documents: TextDocuments<TextDocument> = new TextDocuments(TextDocument);

	constructor() {
		this.documents.onDidChangeContent(this.onDidChangeContent);
		this.documents.onDidOpen(this.onDidOpen);
	}

	private onDidChangeContent(change: TextDocumentChangeEvent<TextDocument>) {
		const uri = change.document.uri;
		const text = change.document.getText();
		// if (text.length !== 0) {
		compile(uri, text);
		// }
		console.log(buffers);
	}

	private onDidOpen(msg: TextDocumentChangeEvent<TextDocument>) {
		const uri = msg.document.uri;
		const text = msg.document.getText();
		compile(uri, text);

		const connection = getConnection();
		const currentDiagnostics = diagnostics.get(uri);
		if (currentDiagnostics) {
			connection.sendDiagnostics({ uri, diagnostics: currentDiagnostics });
		}
	}
}
