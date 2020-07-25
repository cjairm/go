# Description

Given an array of integers and a number, write a function called `maxSubarraySum`, which finds the maximum sum of a subarray with the length of the number passed to the function.

# Example Output

|               Input                | Output |
| :--------------------------------: | :----: |
|      [100, 200, 300, 400], 2       |  700   |
|                                    |        |
| [1, 4, 2, 10, 23, 3, 1, 0, 20], 4  |   39   |
|                                    |        |
|      [-3, 4, 0, -2, 6, -1], 2      |   5    |
|                                    |        |
| [3, -2, 7, -4, 1, -1, 4, -2, 1], 2 |   5    |
|                                    |        |
|             [2, 3], 3              |   0    |

# Pseudo code

This was my thinking before resolve the problem.

```
Save in a variable called currentSum the sum of the first n elements of the array

Create a loop while pointer is lower than the length of the array

    Save in temporal variable the sum of the array from pointer to pointer plus n

    If temporal sum is bigger than currentSum make it equal to temporal sum

    Increment pointer

Return currentSum
```

# How it works?

```
// initial state
// [100, 200, 300, 400], 2

// 2 state
// sum = 100 + 200 = 300

// 3 state
// sum = 200 + 300 = 500

// 4 state
// sum = 300 + 400 = 700

// final state
// sum = 700
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

Jul 25, 2020

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/020_max_sum_sub_array)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/020_max_sum_sub_array)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/020_max_sum_sub_array)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/020_max_sum_sub_array)
