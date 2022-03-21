console.log(require.main === module);

const notMain = require("./require-not-main");
console.log(notMain.isEntryFile);
