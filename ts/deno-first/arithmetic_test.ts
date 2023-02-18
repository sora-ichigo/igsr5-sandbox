import { assertEquals } from "https://deno.land/std@0.170.0/testing/asserts.ts";

Deno.test("add", async () => {
  const { add } = await import("./arithmetic.ts");
  assertEquals(add(1, 2), 3);
});
