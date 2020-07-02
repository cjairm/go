#  Error handling - Description


# Instructions

Start with this code​. Instead of using the blank identifier, make sure the code is checking and handling the error.
```
type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))

}
```

# How to use it

* Download main.go
* Save it in some folder
* Exexute: `go run ./my/path/main.go`

## Author(s)

* Carlos Mendez

### Created

Jul 1, 2020