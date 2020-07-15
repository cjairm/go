# Description

Write a function called "SameFrecuency". Given two positive integers, find out if the two numbers have the same frequency of digits.

The solution MUST have the following complexities: o(n)

# Example Output

| Input            | Output   |
|:----------------:|:--------:|
| 182, 281         | true     |
|                  |          |
| 34. 14           | false    |
|                  |          |
| 3589578, 5879385 | true     |
|                  |          |
| 22, 222          | false    |

# Pseudo code

This was my thinking before resolve the problem.
```
Create SameFrecuency function

    Convert n1 and n2 to string

    Create two helper slices xi1 to store n1 data and xi2 to store n2 data

    Loop over n1 as string 

        Create key in the slice xi1 and / or increment value.

    Loop over n2 as string 

        Create key in the slice xi2 and / or increment value.

    Compare xi1 and xi2
```

# How it works?

```
initial state
// 3589578, 5879385

2 state
// [3,5,8,9,5,7,8], [5,8,7,9,3,8,5]

final state
// {
//     3: 1,
//     5: 2,
//     8: 2,
//     9: 1,
//     7: 1
// }, {
//     5: 2,
//     8: 2,
//     7: 1,
//     9: 1,
//     3: 1
// }
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

Jul 14, 2020

### Same algorithm, different languages

* [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/008_frecuency_counter)
* [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/008_frecuency_counter)
* [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/008_frecuency_counter)
* [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/008_frecuency_counter)