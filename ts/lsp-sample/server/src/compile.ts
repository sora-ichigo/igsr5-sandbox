// eslint-disable-next-line @typescript-eslint/ban-types
export type CompileFiles = Map<string, Object>;

export const compile = (buffers: CompileFiles, uri: string, src: string): CompileFiles => {
	buffers.set(uri, src);
	return buffers;
};

