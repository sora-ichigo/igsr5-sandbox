import fs from "fs";
import { createYoga, createSchema } from "graphql-yoga";
import { Resolvers } from "../../__generated__/gql/resolvers";

const typeDefs = fs.readFileSync("./src/schema.graphql", "utf8");

const resolvers: Resolvers = {
  Query: {
    users() {
      return [{ name: "Nextjs" }];
    },
  },
};

const schema = createSchema({
  typeDefs,
  resolvers,
});

export default createYoga({
  schema,
  // Needed to be defined explicitly because our endpoint lives at a different path other than `/graphql`
  graphqlEndpoint: "/api/graphql",
  graphiql: true,
});
