package main

import (
  "math"
  "fmt"
)

func balancedSplitExists(arr []int) bool {
  // Write your code here
  // Sort my array
  // As we should build two arrays, sum all elements and then divide by 2
  // Start to sum from last to first and at the moment we get the splitedSum thats the position to split,
  // In the last step check for the minNum and compared with the other array 

  if len(arr) <= 1 {
    return false;
  }
  a := MergeSort(arr);

  rs := 0
  for _, val := range a {
    rs += val
  }
  splitNum := int(math.Floor(float64(rs) / float64(2)))

  var count int
  var minNum int = math.MaxInt32
  var i int = len(a) - 1

  for i >= 0 {
    count += a[i]
    if count == splitNum {
      minNum = a[i]
      break;
    } else if count >= splitNum {
      return false
    }

    i--
  }

  f := a[0:i]

  if minNum <= f[len(f) - 1] {
    return false;
  } else {
    return true;
  }
}

// MergeSort sorts an array of numbers
// Create MergeSort function
func MergeSort(xi []int) []int {
	if len(xi) <= 1 {
		return xi
	}
	mid := int(math.Floor(float64(len(xi) / 2)))
	left := MergeSort(xi[0:mid])
	right := MergeSort(xi[mid:])
	return MergeSortedArrays(left, right)
}

// MergeSortedArrays as the name says merge two arrays
func MergeSortedArrays(xi1, xi2 []int) []int {
	// Create an empty array, take a look at the smallest values in each input array
	xxi := []int{}
	var i, j int

	// While there are still values we haven't looked at...
	for i < len(xi1) && j < len(xi2) {
		if xi2[j] > xi1[i] {
			// If the value in the first array is smaller than the value in the second array,
			// push the value in the first array into our results and move on to the next value in the first array
			xxi = append(xxi, xi1[i])
			i++
		} else {
			// If the value in the first array is larger than the value in the second array,
			// push the value in the second array into our results and move on to the next value in the second array
			xxi = append(xxi, xi2[j])
			j++
		}
	}

	// Once we exhaust one array, push in all remaining values from the other array
	for i < len(xi1) {
		xxi = append(xxi, xi1[i])
		i++
	}

	for j < len(xi2) {
		xxi = append(xxi, xi2[j])
		j++
	}

	return xxi
}

func main() {
  // Call balancedSplitExists() with test cases here
  got1 := balancedSplitExists([]int{1,5,7,1})
  if got1 != true {
    fmt.Println("Got: ", got1, ". Expected: ", true)
  } else {
    fmt.Println("✔ Test 1")
  }

  got2 := balancedSplitExists([]int{12,7,6,7,6})
  if got2 != false {
    fmt.Println("Got: ", got2, ". Expected: ", false)
  } else {
    fmt.Println("✔ Test 2")
  }

  got3 := balancedSplitExists([]int{2,1,2,5})
  if got3 != true {
    fmt.Println("Got: ", got3, ". Expected: ", true)
  } else {
    fmt.Println("✔ Test 3")
  }

  got4 := balancedSplitExists([]int{3,6,3,4,4})
  if got4 != false {
    fmt.Println("Got: ", got4, ". Expected: ", false)
  } else {
    fmt.Println("✔ Test 4")
  }

  got5 := balancedSplitExists([]int{0})
  if got5 != false {
    fmt.Println("Got: ", got5, ". Expected: ", false)
  } else {
    fmt.Println("✔ Test 5")
  }

  got6 := balancedSplitExists([]int{1,0})
  if got6 != false {
    fmt.Println("Got: ", got6, ". Expected: ", false)
  } else {
    fmt.Println("✔ Test 6")
  }

  got7 := balancedSplitExists([]int{30,30})
  if got7 != false {
    fmt.Println("Got: ", got7, ". Expected: ", false)
  } else {
    fmt.Println("✔ Test 7")
  }

  got8 := balancedSplitExists([]int{15,30,15})
  if got8 != true {
    fmt.Println("Got: ", got8, ". Expected: ", true)
  } else {
    fmt.Println("✔ Test 8")
  }
}
