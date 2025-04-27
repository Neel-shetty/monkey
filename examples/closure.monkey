let newAdder = fn(x) { fn(y) { x + y } };
let addTwo = newAdder(2);
puts(addTwo(3));
let addThree = newAdder(3);
addThree(10);