# Description

It is a simple factorial calculation. Recursion.

# Example Output

| Input     | Output         |
|:---------:|:--------------:|
| 0         | 0              |
|           |                |
| 1         | 1              |
|           |                |
| 3         | 6              |
|           |                |
| 16        | 20922789888000 |

# Pseudo code

This was my thinking before resolve the problem.
```
Create a function called "FactorialNumber"

If the n is equal to 0 return 0

If the n is equal to 1 return 1

Return n times Factorial function less 1
```

# How it works?

```
// Input - 3

initial state
// 3 * FactorialNumber(3 - 1)
// 2 * FactorialNumber(2 - 1)
// 1 * 1

2 state
// 3 * FactorialNumber(3 - 1)
// 2 * 1

final state
// 3 * 2 = 6

```

# How to use it

* Download main.go and main_test.go
* Save them in some folder
* Run `go run main.go`
	* To run test `go test`

# Author

Carlos Mendez

## Why 0(n)?

This is just to obtain vizually how fast is the program, and is called Big O notation.

### Created at 

Jul 13, 2020

### Same algorithm, different languages

* [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/007_factorial_number)
* [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/007_factorial_number)
* [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/007_factorial_number)
* [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/007_factorial_number)