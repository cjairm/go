# Description

Write a function called `IsValidPalindrome`. Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

`Note`: For the purpose of this problem, we define empty string as valid palindrome.

# Example Output

|                        Input                        | Output |
| :-------------------------------------------------: | :----: |
| IsValidPalindrome("A man, a plan, a canal: Panama") |  true  |
|                                                     |        |
|           IsValidPalindrome("race a car")           | false  |

# Pseudo code

This was my thinking before resolve the problem.

```
Check if string is empty return true

Init helpers variables (pointers) to start from the beggining and the end of the array

Start a loop while right pointer is bigger than left

    if both pointers are alphanumeric

        if pointers values are equal, increment left and decrement right pointers

        otherwise return false

    otherwise
    
        if left pointer is not alphanumeric increment this pointer

        if right pointer is not alphanumeric decrement this pointer
    
return true
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

Jul 31, 2020

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/025_is_valid_palindrome)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/025_is_valid_palindrome)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/025_is_valid_palindrome)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/025_is_valid_palindrome)
