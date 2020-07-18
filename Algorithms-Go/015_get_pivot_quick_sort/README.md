# Description

Merge two sorted arrays.

# Example Output

|           Input            | Output |
| :------------------------: | :----: |
| [ 5, 2, 1, 8, 4, 7, 6, 3 ] |   4    |

# Pseudo code

```
It will help to accept three arguments: an array, a start index, and an end index (these can default to 0 and the array length minus 1, respectively)

Grab the pivot from the start of the array

Store the current pivot index in a variable (this will keep track of where the pivot should end up)

Loop through the array from the start until the end

    If the pivot is greater than the current element, increment the pivot index variable and then swap the current element with the element at the pivot index

Swap the starting element (i.e. the pivot) with the pivot index

Return the pivot index
```

# How it works?

```
initial state
//   p = 5
//   s
// [ 5, 2, 1, 8, 4, 7, 6, 3 ]
//      i

2 state
//   p > 2
//      s
// [ 5, 1, 2, 8, 4, 7, 6, 3 ]
//         i

3 state
//   p < 8
//         s
// [ 5, 1, 2, 8, 4, 7, 6, 3 ]
//            i

4 state
//   p > 4
//            s
// [ 5, 1, 2, 4, 8, 7, 6, 3 ]
//               i

5 state
//   p < 7
//            s
// [ 5, 1, 2, 4, 8, 7, 6, 3 ]
//                  i

6 state
//   p < 6
//            s
// [ 5, 1, 2, 4, 8, 7, 6, 3 ]
//                     i

7 state
//   p > 3
//               s
// [ 5, 1, 2, 4, 3, 7, 6, 8 ]
//                        i

// final state
// [ 3, 1, 2, 4, 5, 7, 6, 8 ]

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

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/015_get_pivot_quick_sort)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/015_get_pivot_quick_sort)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/015_get_pivot_quick_sort)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/015_get_pivot_quick_sort)
