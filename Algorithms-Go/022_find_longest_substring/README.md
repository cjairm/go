# Description (NOT WORKING) 

Write a function called `findLongestSubstring`, which accepts a string and returns the length of the longest substring with all distinct characters.

# Example Output

|       Input        | Output |
| :----------------: | :----: |
|         ''         |   0    |
|                    |        |
|   'rithmschool'    |   7    |
|                    |        |
|  'thisisawesome'   |   6    |
|                    |        |
|  'thecatinthehat'  |   7    |
|                    |        |
|      'bbbbbb'      |   1    |
|                    |        |
| 'longestsubstring' |   8    |
|                    |        |
| 'thisishowwedoit'  |   6    |

# Pseudo code

This was my thinking before resolve the problem.

```

```

# How it works?

```
// Initial state
// letterSeen = {

// } 
// currTotal = 0 
// maxLength = 0
// i
// thecatinthehat
// j

// 2 state
// letterSeen = {
    t: 1
// } 
// currTotal = 1 
// maxLength = 0
// i
// thecatinthehat
//  j

// 3 state
// letterSeen = {
    t: 1
    h: 1
// } 
// currTotal = 2 
// maxLength = 0
// i
// thecatinthehat
//   j

// 4 state
// letterSeen = {
    t: 1
    h: 1
    e: 1
// } 
// currTotal = 3 
// maxLength = 0
// i
// thecatinthehat
//    j

// 5 state
// letterSeen = {
    t: 1
    h: 1
    e: 1
    c: 1
// } 
// currTotal = 4 
// maxLength = 0
// i
// thecatinthehat
//     j

// 6 state
// letterSeen = {
    t: 1
    h: 1
    e: 1
    c: 1
    a: 1
// } 
// currTotal = 5 
// maxLength = 0
// i
// thecatinthehat
//      j

// 7 state
// letterSeen = {

// } 
// currTotal = 0
// maxLength = 4
//  i
// thecatinthehat
//  j

// Final state
// 7
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

Jul 27, 2020

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/022_find_longest_substring)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/022_find_longest_substring)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/022_find_longest_substring)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/022_find_longest_substring)
