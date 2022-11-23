import { CodegenConfig } from "@graphql-codegen/cli";
import { IGraphQLConfig } from "graphql-config";

const codegenConfig: CodegenConfig = {
  overwrite: true,
  schema: "./src/schema.graphql",
  generates: {
    "src/__generated__/gql/": {
      preset: "client",
      plugins: [],
    },
    "./graphql.schema.json": {
      plugins: ["introspection"],
    },
    "./src/__generated__/gql/resolvers.ts": {
      plugins: ["typescript", "typescript-resolvers"],
    },
  },
};

const config: IGraphQLConfig = {
  schema: "./src/schema.graphql",
  extensions: {
    codegen: codegenConfig,
  },
};

export default config;
