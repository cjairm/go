# Description

Sort an array (IS NOT WORKING YET)

# Example Output

|     Input     |                        Output                        |
| :-----------: | :--------------------------------------------------: |
|   [3, 1, 2]   |     [1, 2, 3] or [3, 2, 1] (depends on the user)     |
|               |                                                      |
|      []       |                          []                          |
|               |                                                      |
| [9, 10, 1, 2] | [1, 2, 9, 10] or [10, 9, 2, 1] (depends on the user) |
|               |                                                      |
| [1, 0, 0, 1]  |  [0, 0, 1, 1] or [1, 1, 0, 0] (depends on the user)  |

# Pseudo code

```
Call the pivot helper on the array

When the helper returns to you the updated pivot index, recursively call the pivot helper on the subarray to the left of that index, and the subarray to the right of that index

Your base case occurs when you consider a subarray with less than 2 elements
```

# How it works?

```
// initial state
// [ 5, 2, 1, 8, 4, 7, 6, 3 ]

// 2 state
// [ 5, 2, 1, 4, 3, 7, 6, 8 ]

// 2.5 state
// [ 3, 2, 1, 4, 5, 7, 6, 8 ]

// 3 state
// [ 3, 2, 1, 4, 5, 7, 6, 8 ]

// 3.5 state
// [ 1, 2, 3, 4, 5, 7, 6, 8 ]

// ... states

// final state
// [ 1, 2, 3, 4, 5, 6, 7, 8 ]
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

Jul 23, 2020

### Helper

[Get Pivot](https://github.com/cjairm/go/tree/master/Algorithms-Go/015_get_pivot_quick_sort)

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/014_merge_sort)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/014_merge_sort)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/014_merge_sort)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/014_merge_sort)
