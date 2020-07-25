# Description

Write a function called `isSubsequence` which takes in two strings and checks whether the characters in the first string form a subsequence of the characters in the second string. In other words, the function should check whether the characters in the first string appear somewhere in the second string, `without their order changing.`

# Example Output

|         Input          | Output |
| :--------------------: | :----: |
| "hello", "hello world" |  true  |
|                        |        |
|    "sing", "sting"     |  true  |
|                        |        |
|  "abc", "abracadabra"  |  true  |
|                        |        |
|      "abc", "acb"      | false  |

# Pseudo code

This was my thinking before resolve the problem.

```
Create helpers

Start loop while the second array pointer is bigger than length of array

    If array1 at first pointer if equal to array2 at second pointer, increment the first pointer

    If first pointer meet array1's length return true

    Increment second pointer

Return false

```

# How it works?

```
// init state
//  i
// "sing", "sting"
//          j

// 1 state
//   i
// "sing", "sting"
//           j

// 2 state
//   i
// "sing", "sting"
//            j

// 3 state
//    i
// "sing", "sting"
//             j

// 4 state
//     i
// "sing", "sting"
//              j

// 5 state
// true
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

Jul 24, 2020

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/018_is_subsequence)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/018_is_subsequence)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/018_is_subsequence)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/018_is_subsequence)
