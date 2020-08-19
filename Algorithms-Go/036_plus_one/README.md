# Description

Write a recursive function called `PlusOne`. Given a non-empty array of digits representing a non-negative integer, increment one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

# Example Output

|        Input         |   Output    |
| :------------------: | :---------: |
|  PlusOne([1,2,2,1])  |  [1,2,2,2]  |
|                      |             |
| PlusOne([9,4,9,8,4]) | [9,4,9,8,5] |
|                      |             |
|  PlusOne([8,9,9,9])  |  [9,0,0,0]  |
|                      |             |
|   PlusOne([9,9,9])   |  [1,0,0,0]  |
|                      |             |
|  PlusOne([3,8,9,9])  |  [3,9,0,0]  |
|                      |             |
|     PlusOne([9])     |    [1,0]    |

```
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
```

```
Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
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

August 19, 2020

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/036_plus_one)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/036_plus_one)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/036_plus_one)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/036_plus_one)
