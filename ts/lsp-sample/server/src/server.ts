import { buildConnection } from './connections';
import { DocumentManager } from './documents';

const connection = buildConnection();
const documents = new DocumentManager().documents;

documents.listen(connection);
connection.listen();
