# Description

Write a recursive function called `RotateArray`.

Given an array, rotate the array to the right by k steps, where k is non-negative.

# Example Output

|              Input              |     Output      |
| :-----------------------------: | :-------------: |
| RotateArray([1,2,3,4,5,6,7], 3) | [5,6,7,1,2,3,4] |
|                                 |                 |
| RotateArray([-1,-100,3,99], 2)  | [3,99,-1,-100]  |
|                                 |                 |

```
Example 1.
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
```

```
Example 2.
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation: 
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
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

August 7, 2020

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/032_rotate_array)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/032_rotate_array)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/032_rotate_array)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/032_rotate_array)
