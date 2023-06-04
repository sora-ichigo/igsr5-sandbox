enum E {
  Foo = 10,
  Bar = 20,
}

function takeValue(e: E) {}
takeValue(E.Bar);
takeValue(20);
//takeValue(30); // error! 4.9以前はエラーにならなかった

enum Color {
  Red = "red",
  Orange = "orange",
  Yellow = "yellow",
  Green = "green",
  Blue = "blue",
  Violet = "violet",
}

type PrimaryColor = Color.Red | Color.Green | Color.Blue;

function isPrimaryColor(c: Color): c is PrimaryColor {
  return c === Color.Red || c === Color.Green || c === Color.Blue;
}

enum E2 {
  Blah = Math.random(),
}

// const c: Color = "red"; // error
const c1: Color = Color.Red; // works
