## Frecuency Counters

-   "This pattern uses objects or sets to collect values/frecuencies of values." Colt Steele.
-   "This can often avoid the need for nested loops or O(n^2) operations with arrays/strings." Colt Steele.
-   Examples.
    -   [Anagram](https://github.com/cjairm/go/tree/master/Algorithms-Go/001_anagram)
    -   [Same But Squared](https://github.com/cjairm/go/tree/master/Algorithms-Go/003_same_but_squared)

## Multiple Pointers

-   "Creating pointers of values that correspond to an ondex or position and move towards the beginning, end or middle based on a certain condition." Colt Steele.
-   "Very efficient for solving problems with minimal space complexity as well." Colt Steele.
-   Examples.
    -   [Count Unique Values](https://github.com/cjairm/go/tree/master/Algorithms-Go/002_count_unique_values)
    -   [Pair That Sum Zero](https://github.com/cjairm/go/tree/master/Algorithms-Go/004_pair_that_sum_zero)

## Sliding Window

-   "This pattern involves creating a `window` which can either be an array or number from one position to another." Colt Steele.
-   "Depending on a certain condition, the window either increases or closes (and new window is created)." Colt Steele.
-   "Very useful for keeping track of a subset of data in an array/string etc." Colt Steele.
-   Example.
    -   [Max sub array sum](https://github.com/cjairm/go/tree/master/Algorithms-Go/005_max_sub_array_sum)

## Recursion

-   "It is a process (in these cases functions) that calls itself." Colt Steele.
-   [Martin and The Dragon (History)](https://webdocs.cs.ualberta.ca/~ree/c101-b2/dragonstory0.pdf)
-   Example.
    -   [Recursion Example](https://github.com/cjairm/go/tree/master/Algorithms-Go/007_factorial_number)

## Binary Search (Divide and Conquer)

-   "This pattern involves dividing a data set into smaller chunks and then repeating a process with a subset of data." Colt Steele.
-   "This pattern can tremendously `decrease time complexity`." Colt Steele.
-   Example.
    -   [Search a number](https://github.com/cjairm/go/tree/master/Algorithms-Go/006_search_a_number)

## Sorting

-   "Sorting is a process of rearranging the items in a collection (e.g. an array) so that the items are in some kind of order." Colt Steele.

### Bubble Sort

-   "A sorting algorithm where the largest values bubble up to the top." Colt Steele.
-   Example.
    -   [Bubble Sort](https://github.com/cjairm/go/tree/master/Algorithms-Go/009_bubble_sort)

### Selection Sort

-   "Similar to bubble sort, but instead of first placing large values into sorted position, it places small values into sorted position." Colt Steele.
-   Example.
    -   [Selection Sort](https://github.com/cjairm/go/tree/master/Algorithms-Go/011_selection_sort)

### Insertion Sort

-   "Builds up the sort by gradually creating a larger left half which is always sorted." Colt Steele.
-   Example.
    -   [Insertion Sort](https://github.com/cjairm/go/tree/master/Algorithms-Go/012_insertion_sort)

### Merge Sort

-   "There is a family of sorting algorithms that can improve time complexity from O(n ) to O(n log n)." Colt Steele.
-   "It's a combination of two things - merging and sorting!" Colt Steele.
-   "Exploits the fact that arrays of 0 or 1 element are always sorted." Colt Steele.
-   "Works by decomposing an array into smaller arrays of 0 or 1 elements, then building up a newly sorted array." Colt Steele.
-   Example.
    -   [Merge Sort](https://github.com/cjairm/go/tree/master/Algorithms-Go/014_merge_sort)

### Quick Sort

-   "Like merge sort, exploits the fact that arrays of 0 or 1 element are always sorted." Colt Steele.
-   "Works by selecting one element (called the "pivot") and finding the index where the pivot should end up in the sorted array." Colt Steele.
-   "Once the pivot is positioned appropriately, quick sort can be applied on either side of the pivot." Colt Steele.
-   Example.
    -   [Quick Sort](https://github.com/cjairm/go/tree/master/Algorithms-Go/017_quick_sort)

### Radix Sort

-   "Radix sort is a special sorting algorithm that works on lists of numbers." Colt Steele.
-   "It exploits the fact that information about the size of a number is encoded in the number of digits." Colt Steele.
-   "More digits means a bigger number!." Colt Steele.
-   "It never makes comparisons between elements!." Colt Steele.
-   Example.
    -   [Radix Sort](https://github.com/cjairm/go/tree/master/Algorithms-Go/019_radix_sort)

## Excercises

-   [Frecuency Counter](https://github.com/cjairm/go/tree/master/Algorithms-Go/008_frecuency_counter)
-   [Are There Duplicates?](https://github.com/cjairm/go/tree/master/Algorithms-Go/010_are_there_duplicates)
-   [Merge Sorted Arrays](https://github.com/cjairm/go/tree/master/Algorithms-Go/013_merge_sorted_arrays)
-   [Get Pivot - Quick Sort](https://github.com/cjairm/go/tree/master/Algorithms-Go/015_get_pivot_quick_sort)
-   [Average Pair](https://github.com/cjairm/go/tree/master/Algorithms-Go/016_average_pair)
-   [Is Subsequence?](https://github.com/cjairm/go/tree/master/Algorithms-Go/018_is_subsequence)
-   [Max Sum in SubArray](https://github.com/cjairm/go/tree/master/Algorithms-Go/020_max_sum_sub_array)
-   [Min Sub Array Length](https://github.com/cjairm/go/tree/master/Algorithms-Go/021_min_sub_array_length)
-   [Find Longest Sub String](https://github.com/cjairm/go/tree/master/Algorithms-Go/022_find_longest_substring)
-   [Power(Recursive)](https://github.com/cjairm/go/tree/master/Algorithms-Go/023_power)
-   [Check Clock Angles](https://github.com/cjairm/go/tree/master/Algorithms-Go/024_clock_angles)
-   [Is Valid Palindrome](https://github.com/cjairm/go/tree/master/Algorithms-Go/025_is_valid_palindrome)
-   [Remove Duplicates](https://github.com/cjairm/go/tree/master/Algorithms-Go/026_remove_duplicates)
-   [Remove Consecutive Characthers](https://github.com/cjairm/go/tree/master/Algorithms-Go/027_remove_consecutive_char)
-   [Product of Array](https://github.com/cjairm/go/tree/master/Algorithms-Go/028_array_product)

# Golang Algorithms

My daily practice. These examples are written in different languages

-   [TypeScript](https://github.com/cjairm/typescript/tree/master/Algorithms-TS)
-   [Python](https://github.com/cjairm/python/tree/master/Algoritms-Py)
-   [JavaScript](https://github.com/cjairm/javascript/tree/master/Algorithms-JS)
-   [C++(Arduino)](https://github.com/cjairm/arduino/tree/master/Algorithms-Cpp)
