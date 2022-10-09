const fs = require("fs");
const log = fs.openSync("./log.txt", w);

const languageServer = () => {
  process.stdin.on("readable", () => {
    let chunk = process.stdin.read();
    fs.writeSync(log, chunk);
  });
};

if (process.argv.length !== 3) {
  console.log(`usage: ${process.argv[1]} [--language-server|FILE]`);
} else if (process.argv[2] === "--language-server") {
  languageServer();
} else {
  // TODO:
}
