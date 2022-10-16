import { CompletionItem, CompletionItemKind, Connection, ConnectionOptions, ConnectionStrategy, createConnection, InitializeParams, InitializeResult, ProposedFeatures, TextDocumentPositionParams, TextDocumentSyncKind } from 'vscode-languageserver/node';

let connection: Connection;

export const onInitialize = ((_params: InitializeParams) => {
	const result: InitializeResult = {
		capabilities: {
			textDocumentSync: TextDocumentSyncKind.Incremental,
			completionProvider: {
				resolveProvider: true
			}
		}
	};
	return result;
});

export const onCompletion = (_textDocumentPosition: TextDocumentPositionParams): CompletionItem[] => {
	return [
		{
			label: 'TypeScript',
			kind: CompletionItemKind.Text,
			data: 1
		},
		{
			label: 'JavaScript',
			kind: CompletionItemKind.Text,
			data: 2
		}
	];
};

export const onCompletionResolve = (
	(item: CompletionItem): CompletionItem => {
		if (item.data === 1) {
			item.detail = 'TypeScript details';
			item.documentation = 'TypeScript documentation';
		} else if (item.data === 2) {
			item.detail = 'JavaScript details';
			item.documentation = 'JavaScript documentation';
		}
		return item;
	}
);

export const buildConnection = () => {
	connection = createConnection(ProposedFeatures.all);

	connection.onInitialize(onInitialize);
	connection.onCompletion(onCompletion);
	connection.onCompletionResolve(onCompletionResolve);

	return connection;
};

export const getConnection = () => {
	return connection;
};
