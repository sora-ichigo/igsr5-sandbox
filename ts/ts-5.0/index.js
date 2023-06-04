var E;
(function (E) {
    E[E["Foo"] = 10] = "Foo";
    E[E["Bar"] = 20] = "Bar";
})(E || (E = {}));
function takeValue(e) { }
takeValue(E.Bar);
takeValue(20);
takeValue(30); // error! 4.9以前はエラーにならなかった
var Color;
(function (Color) {
    Color["Red"] = "red";
    Color["Orange"] = "orange";
    Color["Yellow"] = "yellow";
    Color["Green"] = "green";
    Color["Blue"] = "blue";
    Color["Violet"] = "violet";
})(Color || (Color = {}));
function isPrimaryColor(c) {
    return c === Color.Red || c === Color.Green || c === Color.Blue;
}
var E2;
(function (E2) {
    E2[E2["Blah"] = Math.random()] = "Blah";
})(E2 || (E2 = {}));
// const c: Color = "red"; // error
var c1 = Color.Red; // works
