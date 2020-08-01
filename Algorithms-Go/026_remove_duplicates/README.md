# Description

Write a function called `RemoveDuplicates`.Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by `modifying the input array` in-place with O(1) extra memory.

# Example Output

|                  Input                  |   Output    |
| :-------------------------------------: | :---------: |
|        RemoveDuplicates([1,1,2])        |    [1,2]    |
|                                         |             |
| RemoveDuplicates([0,0,1,1,1,2,2,3,3,4]) | [0,1,2,3,4] |
|                                         |             |

```
Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
```

```
Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
```

# Pseudo code

This was my thinking before resolve the problem.

```
Sort array

Init helper variables (two pointers)

Start loop until first pointer reach the end of the array

    If first pointer is equal to the second pointer move pointer two until find diferent character and swap them.

    Otherwise move two pointers.

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

Aug 1, 2020

### Helper

[Merge Arrays](https://github.com/cjairm/go/tree/master/Algorithms-Go/014_merge_sort)

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/026_remove_duplicates)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/026_remove_duplicates)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/026_remove_duplicates)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/026_remove_duplicates)
