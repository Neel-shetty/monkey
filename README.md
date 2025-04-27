# Monkey Programming Language

A simple, dynamically typed scripting language with C-like syntax, inspired by the book ["Writing An Interpreter In Go"](https://interpreterbook.com/) by Thorsten Ball.

## Features

- C-like syntax with minimal punctuation
- Dynamic typing
- First-class functions and closures
- Higher-order functions
- Built-in data types: integers, booleans, strings, arrays, hashes
- Simple REPL environment

## Example Code

```monkey
let x = 5;
let y = 10;
let name = "Monkey";
let isCool = true;
let array = [1, 2, 3, 4, 5];
let hash = {"name": "Monkey", "age": 5};

let add = fn(a, b) {
    return a + b;
};

let fibonacci = fn(n) {
    if (n == 0) { return 0; }
    if (n == 1) { return 1; }
    return fibonacci(n - 1) + fibonacci(n - 2);
};

puts("Fibonacci of 10 is:", fibonacci(10));
```

## Getting Started

### Prerequisites

- Go 1.13+ (for the reference implementation)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Neel-shetty/monkey.git
cd monkey
```

2. Build the interpreter:
```bash
go build -o monkey .
```

### Usage

**REPL Mode:**
```bash
./monkey
```

**Run a script:**
```bash
./monkey examples/hello.monkey
```

## Language Reference

### Data Types

- **Integers**: `let x = 42;`
- **Booleans**: `let flag = true;`
- **Strings**: `let name = "Monkey";`
- **Arrays**: `let arr = [1, 2, 3];`
- **Hashes**: `let person = {"name": "Alice", "age": 30};`

### Control Structures

**Conditionals:**
```monkey
if (x > y) {
    puts("x is greater");
} else {
    puts("y is greater or equal");
}
```

**Loops:**
```monkey
let i = 0;
while (i < 10) {
    puts(i);
    i = i + 1;
}
```

### Functions

```monkey
// Simple function
let greet = fn(name) {
    return "Hello " + name + "!";
};

// Higher-order function
let applyTwice = fn(f, x) {
    return f(f(x));
};
```

### Built-in Functions

- `puts(...)`: Print values to stdout
- `len(x)`: Get length of string or array
- `first(arr)`: Get first element of array
- `rest(arr)`: Get all elements except first
- `push(arr, val)`: Append value to array

## Examples

See the [examples/](examples/) directory for more code samples:

- `array.monkey`: Array operations and higher-order functions
- `closure.monkey`: Demonstrates closures
- `hashmap.monkey`: Hash map data structure usage
- `hello.monkey`: Simple hello world example
- `math.monkey`: Basic mathematical operations

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

[MIT License](LICENSE)

---

*Monkey was created as part of the book "Writing An Interpreter In Go" by Thorsten Ball. This implementation is for educational purposes.*