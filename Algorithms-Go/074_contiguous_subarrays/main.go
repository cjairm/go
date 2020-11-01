package main

import "fmt"

func countSubarrays(arr []int) []int {
  // Write your code here
  var s []int = []int{}
  var c int

  for i := range arr {
    c = 0

    for j := i - 1; j >= 0; j-- {
      if arr[i] < arr[j] {
        break
      } else {
        c++
      }
    }

    for k := i + 1; len(arr) > k; k++ {
      if arr[i] < arr[k] {
        break
      } else {
        c++
      }
    }

    s = append(s, c + 1)
  }

  return s;
}

// Equal is a function just to compare to slices(arrays)
func Equal(a, b []int) bool {
  if len(a) != len(b) {
    return false
  }

  for i, v := range a {
    if v != b[i] {
      return false
    }
  }

  return true
}

func main() {
  // Call countSubarrays() with test cases here
  got1 := countSubarrays([]int{3, 4, 1, 6, 2})
  if !Equal(got1, []int{1, 3, 1, 5, 1}) {
    fmt.Println("Got: ", got1, ". Expected: ", []int{1, 3, 1, 5, 1})
  } else {
    fmt.Println("✔ Test 1")
  }

  got2 := countSubarrays([]int{2, 4, 7, 1, 5, 3})
  if !Equal(got2, []int{1, 2, 6, 1, 3, 1}) {
    fmt.Println("Got: ", got2, ". Expected: ", []int{1, 2, 6, 1, 3, 1})
  } else {
    fmt.Println("✔ Test 2")
  }

  got3 := countSubarrays([]int{})
  if !Equal(got3, []int{}) {
    fmt.Println("Got: ", got3, ". Expected: ", []int{})
  } else {
    fmt.Println("✔ Test 3")
  }

  got4 := countSubarrays([]int{1000})
  if !Equal(got4, []int{1}) {
    fmt.Println("Got: ", got4, ". Expected: ", []int{1})
  } else {
    fmt.Println("✔ Test 4")
  }

  got5 := countSubarrays([]int{10000, 0, 0, 0, 0})
  if !Equal(got5, []int{5, 4, 4, 4, 4}) {
    fmt.Println("Got: ", got5, ". Expected: ", []int{5, 4, 4, 4, 4})
  } else {
    fmt.Println("✔ Test 5")
  }

  got6 := countSubarrays([]int{-1, 0, 1, 2, 10000})
  if !Equal(got6, []int{1, 2, 3, 4, 5}) {
    fmt.Println("Got: ", got6, ". Expected: ", []int{1, 2, 3, 4, 5})
  } else {
    fmt.Println("✔ Test 6")
  }
}
