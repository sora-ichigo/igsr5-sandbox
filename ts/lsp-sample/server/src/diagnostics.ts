import { Diagnostic } from 'vscode-languageserver/node';

export type Diagnostics = Map<string, Diagnostic[]>;
