# Description

Implement a function called, `areThereDuplicates` which accepts a `variable number of arguments`, and checks whether there are any duplicates among the arguments passed in. You can solve this using the frequency counter pattern OR the multiple pointers pattern.

# Example Output

| Input     | Output |
|:---------:|:------:|
| [1, 2, 3] | false  |
|           |        |
| [2, 2, 1] | true   |
|           |        |
| [0, 0, 0] | true   |

# Pseudo code

```
Sort xi

Init two helpers (pointers) h1, h2

Loop xi until the pointer is greater than the length of xi

    If the value at h1 it's equal to h2

        Return true because duplicates exists

    Increment in one h1 and h2

Return false because duplicate was not found
    
```

# How it works?

```
initial state
// [2, 2, 1]

2 state
// [1, 2, 2]

3 state
// compate 1 vs 2

...states
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

* [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/010_are_there_duplicates)
* [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/010_are_there_duplicates)
* [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/010_are_there_duplicates)
* [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/010_are_there_duplicates)