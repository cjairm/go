# Description

Sort an array

# Example Output

| Input            | Output                                                |
|:----------------:|:-----------------------------------------------------:|
| [3, 1, 2]        | [1, 2, 3] or [3, 2, 1] (depends on the user)          |
|                  |                                                       |
| []               | []                                                    |
|                  |                                                       |
| [9, 10, 1, 2]    | [1, 2, 9, 10] or [10, 9, 2, 1] (depends on the user)  |
|                  |                                                       |
| [1, 0, 0, 1]     | [0, 0, 1, 1] or [1, 1, 0, 0] (depends on the user)    |

# Pseudo code

```
Store the first element as the smallest value you've seen so far.

Compare this item to the next item in the array until you find a smaller number.

If a smaller number is found, designate that smaller number to be the new "minimum" and continue until the end of the array.

If the "minimum" is not the value (index) you initially began with, swap the two values.

Repeat this with the next element until the array is sorted.
```

# How it works?

```
initial state
//   f
// [ 5, 3, 4, 1, 2 ]
//      s

1 state
//   f
// [ 5, 3, 4, 1, 2 ]
//     min s

2 state
//   f
// [ 5, 3, 4, 1, 2 ]
//     min    s

3 state
//   f
// [ 5, 3, 4, 1, 2 ]
//          min  s

4 state
// [ 1, 3, 4, 5, 2 ]

... states
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

Jul 16, 2020

### Same algorithm, different languages

* [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/011_selection_sort)
* [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/011_selection_sort)
* [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/011_selection_sort)
* [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/011_selection_sort)