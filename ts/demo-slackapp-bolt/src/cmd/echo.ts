import { SlackCommandMiddlewareArgs } from "@slack/bolt";

export const echo = async ({
  command,
  ack,
  say,
}: SlackCommandMiddlewareArgs) => {
  ack();
  say(`${command.text}`);
};
