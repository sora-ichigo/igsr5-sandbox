import { getPageOGPMetadata } from "metagros/move/cometPunch.js";

const url = "https://www.wantedly.com/companies/wantedly/post_articles/385515";

const metadata = getPageOGPMetadata(url);
console.log(metadata);
