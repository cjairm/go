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
Define a function that accepts list of numbers

Figure out how many digits the largest number has

Loop from k = 0 up to this largest number of digits

For each iteration of the loop:

    Create buckets for each digit (0 to 9)

    Place each number in the corresponding bucket based on its kth digit

Replace our existing array with values in our buckets, starting with 0 and going up to 9

return list at the end!
```

# How it works?

```
// initial state
// [ 8, 3, 5, 4, 7, 6, 1, 2, 302, 99, 498 ]
// bunch for 0s
// []
// bunch for 1s
// [1]
// bunch for 2s
// [2, 302]
// bunch for 3s
// [3]
// bunch for 4s
// [4]
// bunch for 5s
// [5]
// bunch for 6s
// [6]
// bunch for 7s
// [7]
// bunch for 8s
// [8,498]
// bunch for 9s
// [99]

// 2 state
// [ 1, 2, 302, 3, 4, 5, 6, 7, 8, 498, 99 ]
// bunch for 0s
// [1, 2, 302, 3, 4, 5, 6, 7, 8]
// bunch for 1s
// []
// bunch for 2s
// []
// bunch for 3s
// []
// bunch for 4s
// []
// bunch for 5s
// []
// bunch for 6s
// []
// bunch for 7s
// []
// bunch for 8s
// []
// bunch for 9s
// [498, 99]

// 3 state
// [ 1, 2, 302, 3, 4, 5, 6, 7, 8, 498, 99 ]
// bunch for 0s
// [1, 2, 3, 4, 5, 6, 7, 8, 99]
// bunch for 1s
// []
// bunch for 2s
// []
// bunch for 3s
// [302]
// bunch for 4s
// [498]
// bunch for 5s
// []
// bunch for 6s
// []
// bunch for 7s
// []
// bunch for 8s
// []
// bunch for 9s
// []

// final state
// [ 1, 2, 3, 4, 5, 6, 7, 8, 99, 302, 498 ]
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

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/019_radix_sort)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/019_radix_sort)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/019_radix_sort)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/019_radix_sort)
