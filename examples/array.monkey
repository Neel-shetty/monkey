let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            accumulated
        } else {
            iter(rest(arr), push(accumulated, f(first(arr))))
        }
    };
    iter(arr, [])
};

let double = fn(x) {
    x+x
}

let numbers = [1, 2, 3, 4, 5];
puts("First number:", numbers[0]);
puts("Last number:", numbers[4]);

let withExpr = [1 + 1, 2 * 3, 10 - 5];
puts("Expression array:", withExpr);
puts("Doubled expression array:", map(withExpr, double))

let fruits = ["apple", "banana", "orange"];
puts("Number of fruits:", len(fruits));
puts("First fruit:", first(fruits));
puts("Last fruit:", last(fruits));
puts("Rest of fruits:", rest(fruits));

let reduce = fn(arr, initial, f) {
    let iter = fn(arr, result) {
        if (len(arr) == 0) {
        result
        } else {
        iter(rest(arr), f(result, first(arr)));
        }
    };
    iter(arr, initial);
};

let sum = fn(arr) {
    reduce(arr, 0, fn(initial, el) { initial + el });
};

puts("Sum of numbers:", sum([1, 2, 3, 4, 5]));