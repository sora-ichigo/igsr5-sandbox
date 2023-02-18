// import { add, multiply } from "./arithmetic.ts";
//
// const totalCost = (outbound: number, inbound: number, tax: number): number => {
//   return multiply(add(outbound, inbound), tax);
// };

// console.log(totalCost(19, 31, 1.1));

// type Fish = {
//   swim: () => void;
// };
//
// type Bird = {
//   fly: () => void;
// };
//
// function isFish(pet: Fish | Bird | {}): pet is Fish {
//   return (pet as Fish).swim !== undefined;
// }
//
// let pet: Fish | Bird | {} = {
//   swim: () => {},
//   fly: () => {},
// };
//
// pet.swim;
// if ("swim" in pet) {
//   pet.swim;
// }
//
// if (isFish(pet)) {
//   pet.swim();
// }
//
//

interface ApiCall {
  data: Response;
}

interface ApiCall {
  // data: string;
  error?: Error;
}

class CustomClient implements ApiCall {
  data = new Response();
}

type ApiCallType = {
  data: Response;
};

type ApiCallType2 = {
  data: string;
  error?: Error;
};

type T = ApiCallType & ApiCallType2;

const t: T = {
  data: "aa" as never,
};
