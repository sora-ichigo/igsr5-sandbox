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

class PersonImpl {
  name: string;
  age: number;
  constructor(name: string, age: number) {
    this.name = name;
    this.age = age;
  }
}

interface Person {
  new (name: string, age: number): void;
  name: string;
  age: number;
}

const person1: PersonImpl = new PersonImpl("", 1);
const person2: InstanceType<typeof PersonImpl> = new PersonImpl("", 1);

type PersonType = {
  new (name: string, age: number): void;
  name: string;
  age: number;
};

type Personc = ConstructorParameters<PersonType>;

type T1 = Parameters<(s: string, n: number) => void>;
type T2 = ConstructorParameters<ErrorConstructor>;
const t1: T1 = ["hello", 1];
(...t1: T1) => {};
type T3 = [
  arg: {
    a: number;
    b: string;
  }
];

const t3: T3[0] = [{ a: 1, b: "2" }];

function toHex(this: Number, hoge: stinterface DB {
  filterUsers(filter: (this: User) => boolean): User[];
}
 
type c = Capitalize<"hello world">
type uc = Uncapitalize<"Hello world">
type u = Uppercase<"hello world">
type l = Lowercase<"HEllo world">

type ObjectDescriptor<D, M> = {
  data?: D;
  methods?: M; // Type of 'this' in methods is D & M
};

function makeObject<D, M>(desc: ObjectDescriptor<D, M>): D & M {
  let data: object = desc.data || {};
  let methods: object = desc.methods || {};
  return { ...data, ...methods } as D & M;
}
 
makeObject.bind({ x: 1, y: 2 });
let obj = makeObject({
  data: { x: 0, y: 0 },
  methods: {
    moveBy(dx: number, dy: number) {
      this.x += dx; // Strongly typed this
      this.y += dy; // Strongly typed this
    },
  },
});

 
obj.x = 10;
obj.y = 20;
obj.moveBy(5, 5);
