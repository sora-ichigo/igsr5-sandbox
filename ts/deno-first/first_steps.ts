// console.log("hello world");
//
// const url = Deno.args[0];
// const res = await fetch(url);
//
// const body = new Uint8Array(await res.arrayBuffer());
// await Deno.stdout.write(body);
//

// import { copy } from "https://deno.land/std@0.170.0/streams/copy.ts";
//
// const filenames = Deno.args;
// filenames.forEach(async (filename) => {
//   const file = await Deno.open(filename);
//   await copy(file, Deno.stdout);
//   file.close();
// });
//

import { serve } from "https://deno.land/std@0.170.0/http/server.ts";

const handler = async (request: Request) => {
  const resp = await fetch("https://api.github.com/users/denoland", {
    headers: {
      accept: "application/json",
    },
  });

  return new Response(resp.body, {
    status: resp.status,
    headers: {
      "content-type": "application/json",
    },
  });
};

console.log("listening on port 8000");
serve(handler);
