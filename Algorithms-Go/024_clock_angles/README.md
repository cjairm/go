# Description

Write a function called `clockAngle` that returns the acute angle between two clock hands, with two integers for the number of hours and number of minutes.

# Example Output

|      Input      | Output |
| :-------------: | :----: |
| ClockAngle(3,0) |   90   |
|                 |        |
| ClockAngle(6,0) |  180   |

# Pseudo code

This was my thinking before resolve the problem.

```
Check if hours are valid
Check if minutes are valid

If hours are equal to 12 PM/AM just has to be reset to zero.

If minutes are equal to 60 minutes just has to be reset to zero.
But if we reset we have to increment one hour and fotmat it to AM/PM.

MINUTES: The clock rotates 360 in 60 minutes so each minute represents 6 degrees.
HOURS: The clock rotates 360 in 12 hours so each ohur represent 30 degrees.
PAST HOURS: current minute divided by total minutes in a hours, then multiplied for 30 degrees (hour degrees)
(m / 60) * 30 = (m / 30) * 15 = (m / 6) * 3 = (m / 2) * 1 = m / 2
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

Jul 29, 2020

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/024_clock_angles)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/024_clock_angles)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/024_clock_angles)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/024_clock_angles)
