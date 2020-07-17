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
Start by picking the second element in the array

Now compare the second element with the one before it and swap if necessary.

Continue to the next element and if it is in the incorrect order, iterate through the sorted portion (i.e. the left side) to place the element in the correct place.

Repeat until the array is sorted.
```

# How it works?

```
initial state
//   f
// [ 5, 3, 4, 1, 2 ]
//      s
// compare f vs s and swap if necesary

1 state
//      f
// [ 3, 5, 4, 1, 2 ]
//         s
// compare f vs s and swap if necesary

2 state
//   f
// [ 3, 5, 5, 1, 2 ]
//         s
// compare f vs s and swap if necesary

3 state
//         f
// [ 3, 4, 5, 1, 2 ]
//            s
// compare f vs s and swap if necesary

... states

// final state
// [ 1, 2, 3, 4, 5 ]

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

Jul 17, 2020

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/012_insertion_sort)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/012_insertion_sort)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/012_insertion_sort)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/012_insertion_sort)
