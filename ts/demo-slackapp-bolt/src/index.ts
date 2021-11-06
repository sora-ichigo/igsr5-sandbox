import { App } from "@slack/bolt";
import dotenv from "dotenv";
import { echo } from "./cmd/echo";

dotenv.config();
const config = {
  token: process.env.BOT_USER_OAUTH_TOKEN,
  signingSecret: process.env.SIGNING_SECRET,
};

const app = new App(config);
app.command("/echo",echo);

app.start(3000);
console.log(`⚡️ Bolt app is running!`);
