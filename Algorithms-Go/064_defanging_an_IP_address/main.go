package main

import "strings"

func main() {

}

func defangIPaddr(address string) string {
	addressSplit := strings.Split(address, "")

	for i, currAdd := range addressSplit {
		if currAdd == "." {
			addressSplit[i] = "[.]"
		}
	}

	return strings.Join(addressSplit, "")
}
