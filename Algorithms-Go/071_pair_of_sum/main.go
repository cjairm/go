package main

import (
  "fmt"
)

func numberOfWays(arr []int, k int) int {
  // Write your code here
  mp := make(map[int]int)

  for _, v := range arr {
    mp[v]++
  }

  var tc int

  for _, cv := range arr {
    tc += mp[k-cv]

    if k-cv == cv {
      tc--
    }
  }

  return tc/2
}

func main() {
  // Call numberOfWays() with test cases here
  got1 := numberOfWays([]int{1, 2, 3, 4, 3}, 6)
  if got1 != 2 {
    fmt.Println("Got: ", got1, ". Expected: ", 2)
  } else {
    fmt.Println("✔ Test 1")
  }

  got2 := numberOfWays([]int{1, 5, 3, 3, 3}, 6)
  if got2 != 4 {
    fmt.Println("Got: ", got2, ". Expected: ", 4)
  } else {
    fmt.Println("✔ Test 2")
  }
}
