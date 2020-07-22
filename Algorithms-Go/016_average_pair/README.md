# Description

Write a function called `averagePair.` Given a sorted array of integers and a target average, determine if there is a pair of values in the array where the average of the pair equals the target average. There may be more than one pair that matches the average target.

# Example Output

|               Input               | Output |
| :-------------------------------: | :----: |
|          [1, 2, 3], 2.5           |  true  |
|                                   |        |
| [1, 3, 3, 5, 6, 7, 10, 12, 19], 8 |  true  |
|                                   |        |
|     [-1, 0, 3, 4, 5, 6], 4.1      | false  |
|                                   |        |
|               [], 4               | false  |

# Pseudo code

This was my thinking before resolve the problem.

```
Sort the array

Create two helper pointers

Start looping the array while the end pointer is bigger than start pointer

    Calculate the average

    If average is equal to the num searched return true

    If num is bigger than average, so move start pointer to the rigth

    If average is bigger than num, so move end pointer to the left

return false

```

# How it works?

```
// initial state
// 1 + 19 = 20 / 2 = 10 | 10 bigger the 8
//   f
// [ 1, 3, 3, 5, 6, 7, 10, 12, 19 ], 8
//                             s

// 2 state
// 1 + 12 = 13 / 2 = 6.5 | 8 bigger the 6.5
//      f
// [ 1, 3, 3, 5, 6, 7, 10, 12, 19 ], 8
//                         s

// 3 state
// 3 + 12 = 15 / 2 = 7.5 | 8 bigger the 7.5
//         f
// [ 1, 3, 3, 5, 6, 7, 10, 12, 19 ], 8
//                         s

// 4 state
// 5 + 12 = 17 / 2 = 7.5 | 8 bigger the 7.5
//            f
// [ 1, 3, 3, 5, 6, 7, 10, 12, 19 ], 8
//                         s

...states

// final state
// 6 + 10 = 16 / 2 = 8 | 8 equal 8
//               f
// [ 1, 3, 3, 5, 6, 7, 10, 12, 19 ], 8
//                     s
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

Jul 21, 2020

### Helper

[Merge Sort](https://github.com/cjairm/go/tree/master/Algorithms-Go/014_merge_sort)

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/016_average_pair)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/016_average_pair)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/016_average_pair)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/016_average_pair)
