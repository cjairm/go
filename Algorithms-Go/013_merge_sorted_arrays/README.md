# Description

Merge two sorted arrays.

# Example Output

|          Input           |        Output         |
| :----------------------: | :-------------------: |
| [1,3,4,7], [6,10,11,12]  | [1,3,4,6,7,10,11,12]  |
|                          |                       |
|       [], [1,2,3]        |        [1,2,3]        |
|                          |                       |
|      [1, 2], [3,4]       |       [1,2,3,4]       |
|                          |                       |
| [10,100,200], [1,11,111] | [1,10,11,100,111,200] |

# Pseudo code

```
Create an empty array, take a look at the smallest values in each input array

While there are still values we haven't looked at...

    If the value in the first array is smaller than the value in the second array, push the value in the first array into our results and move on to the next value in the first array

    If the value in the first array is larger than the value in the second array, push the value in the second array into our results and move on to the next value in the second array

    Once we exhaust one array, push in all remaining values from the other array
```

# How it works?

```
initial state
//  f             s 
// [1, 3, 4, 7], [6, 10, 11, 12]
// [1]

2 state
//     f          s 
// [1, 3, 4, 7], [6, 10, 11, 12]
// [1, 3]

3 state
//        f       s 
// [1, 3, 4, 7], [6, 10, 11, 12]
// [1, 3, 4]

4 state
//           f        s 
// [1, 3, 4, 7], [6, 10, 11, 12]
// [1, 3, 4, 6]

5 state
//           f        s
// [1, 3, 4, 7], [6, 10, 11, 12]
// [1, 3, 4, 6, 7]

... states

// final state
// [ 1, 3, 4, 6, 7, 10, 11, 12 ]

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

Jul 17, 2020

### Same algorithm, different languages

-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS/012_insertion_sort)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py/012_insertion_sort)
-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS/012_insertion_sort)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp/012_insertion_sort)
