import { TextDocument } from 'vscode-languageserver-textdocument';
import { Connection, DiagnosticSeverity, TextDocumentChangeEvent, TextDocuments } from 'vscode-languageserver/node';
import { compile, CompileFiles } from './compile';
import { getConnection } from './connections';
import { Diagnostics } from './diagnostics';

let buffers: CompileFiles = new Map();
const diagnostics: Diagnostics = new Map();

export class DocumentManager {
	public documents: TextDocuments<TextDocument> = new TextDocuments(TextDocument);

	constructor() {
		this.documents.onDidChangeContent(this.onDidChangeContent);
	}

	private onDidChangeContent(change: TextDocumentChangeEvent<TextDocument>) {
		buffers = compile(buffers, change.document.uri, change.document.getText());

		diagnostics.set(change.document.uri, [{
			severity: DiagnosticSeverity.Information,
			range: {
				start: { line: 0, character: 0 },
				end: { line: 0, character: 10 }
			},
			message: "diagnostic message 1"
		}]);

		const connection = getConnection();
		diagnostics.forEach((diagnostics, uri) => {
			connection.sendDiagnostics({ uri: uri, diagnostics: diagnostics });
		});
	}
}
