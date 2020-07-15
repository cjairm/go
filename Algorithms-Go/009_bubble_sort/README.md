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
Start looping from with a variable called i the end of the array towards the beggining

    Start an inner loop with a variable called j from the beggining until i - 1

    If arr[j] is greater|lower than arr[j+1], swap those values 

Return sorted arr
```

# How it works?

```
initial state
// [100, 3, 90, 1, 0]

2 state
// [3, 100, 90, 1, 0]

3 state
// [3, 90, 100, 1, 0]

4 state
// [3, 90, 1, 100, 0]

5 state
// [3, 90, 100, 0, 1]

...states

final state
// [0, 1, 3, 90, 100]
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

Jul 15, 2020

### Same algorithm, different languages

* [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/009_bubble_sort)
* [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/009_bubble_sort)
* [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/009_bubble_sort)
* [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/009_bubble_sort)