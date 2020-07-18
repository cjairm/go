# Description

Sort an array

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
Break up the array into halves until you have arrays that are empty or have one element

Once you have smaller sorted arrays, merge those arrays with other sorted arrays until you are back at the full length of the array

Once the array has been merged back together, return the merged (and sorted!) array
```

# How it works?

```
// initial state
// [ 8, 3, 5, 4, 7, 6, 1, 2 ]

// 2 state
// [ 8, 3, 5, 4 ], [ 7, 6, 1, 2 ]

// 3 state
// [ 8, 3 ], [ 5, 4 ], [ 1, 2 ], [ 7, 6 ]

// 4 state
// [ 8 ], [ 3 ], [ 5 ], [ 4 ], [ 7 ], [ 6 ], [ 1 ], [ 2 ]

// 5 state
// [ 3, 8 ], [ 4, 5 ], [ 6, 7 ], [ 1, 2 ]

// 6 state
// [ 3, 4, 5, 8 ], [ 1, 2, 6, 7 ]

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

Jul 18, 2020

### Helper

[Merge Sorted Arrays](https://github.com/cjairm/go/tree/master/Algorithms-Go/013_merge_sorted_arrays)

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/014_merge_sort)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/014_merge_sort)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/014_merge_sort)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/014_merge_sort)
