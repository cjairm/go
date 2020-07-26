# Description

Write a function called `minSubArrayLen` which accepts two parameters - an array of positive integers and a positive integer.

This function should return the minimal length of a `contigous` subarray of which the sum is greater than or equal to the integer passed to the function. If there isn't one, return 0 instead.

# Example Output

|                   Input                    | Output |
| :----------------------------------------: | :----: |
|           [2, 3, 1, 2, 4, 3], 7            |   2    |
|                                            |        |
|             [2, 1, 6, 5, 4], 9             |   2    |
|                                            |        |
| [3, 1, 7, 11, 2, 9, 8, 21, 62, 33, 19], 52 |   1    |
|                                            |        |
|     [1, 4, 16, 22, 5, 7, 8, 9, 10], 39     |   3    |
|                                            |        |
|     [1, 4, 16, 22, 5, 7, 8, 9, 10], 55     |   5    |
|                                            |        |
|         [4, 3, 3, 8, 1, 2, 3], 11          |   2    |
|                                            |        |
|     [1, 4, 16, 22, 5, 7, 8, 9, 10], 95     |   0    |

# Pseudo code

This was my thinking before resolve the problem.

```
Init helper variables

Init variable that tends to infinity

Init loop while first pointer is lower than length of the array

    If current total sum is lower than the number expected 

        Check that length of the second pointer and break if necessary

        Increment current sum by adding the current element (array[position])

        Increment the count of the elements

        Increment the second pointer

    If the current sum is equal or higher than the number expected

        If the count of elements is lower than the variable that tends to infinity, save that count in the infinity variable

        Reset count of elements, current total and increment the first pointer

        Then make equal the second pointer to the first one

Just in case the infinity variable never change, return zero

Return infinity variable
```

# How it works?

```
// initial state
// length = 0
// count = 1
// total = 2
//  f
// [2, 3, 1, 2, 4, 3], 7
//  s

// 2 state
// length = 0
// count = 2
// total = 2 + 3 = 5
//  f
// [2, 3, 1, 2, 4, 3], 7
//     s

// 3 state
// length = 0
// count = 3
// total = 2 + 3 + 1 = 6
//  f
// [2, 3, 1, 2, 4, 3], 7
//        s

// 4 state
// length = 4
// count = 4
// total = 2 + 3 + 1 + 2 = 8
//  f
// [2, 3, 1, 2, 4, 3], 7
//           s

// 5 state
// length = 4
// count = 1
// total = 3
//     f
// [2, 3, 1, 2, 4, 3], 7
//     s

// 6 state
// length = 4
// count = 2
// total = 3 + 1 = 4
//     f
// [2, 3, 1, 2, 4, 3], 7
//        s

// 7 state
// length = 4
// count = 3
// total = 3 + 1 + 2 = 6
//     f
// [2, 3, 1, 2, 4, 3], 7
//           s

// 8 state
// length = 4
// count = 4
// total = 3 + 1 + 2 + 4 = 10
//     f
// [2, 3, 1, 2, 4, 3], 7
//              s

// 9 state
// length = 4
// count = 1
// total = 1
//        f
// [2, 3, 1, 2, 4, 3], 7
//        s

// ...states

// final state
// length = 2 -> because [4, 3] is the smalles subarray
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
