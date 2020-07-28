# Description

Write a function called `power` which accepts a base and an exponent. The function should return the power of the base to the exponent.

# Example Output

|   Input    | Output |
| :--------: | :----: |
| Power(2,0) |   1    |
|            |        |
| Power(2,2) |   4    |
|            |        |
| Power(2,4) |   16   |

# Pseudo code

This was my thinking before resolve the problem.

```
Check if the exponent is equal to zero return base of exponent zero = one

return base number times the result of the next value with exponent less one
```

# How it works?

```
// Statement
// 2^4

// stack 2 * (2 ^ 3)

// stack 2 * (2 ^ 2)

// stack 2 * (2 ^ 1)

// stack 2 * (2 ^ 0)

// return 2 * (2 ^ 0) = 2 * 1 = 2

// return 2 * (2 ^ 1) = 2 * 2 = 4

// return 2 * (2 ^ 2) = 2 * 4 = 8

// return 2 * (2 ^ 3) = 2 * 8 = 16

// Final state
// 16
```

# How to use it

-   Download main.go and main_test.go
-   Save them in some folder
-   Run `go run main.go` \* To run test `go test`

# Author

Carlos Mendez

## Why 0(n)?

This is just to obtain vizually how fast is the program, and is called Big O notation.

### Created at

Jul 28, 2020

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/023_power)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/023_power)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/023_power)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/023_power)
