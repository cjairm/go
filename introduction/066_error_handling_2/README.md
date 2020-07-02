#  Error handling - Description


# Instructions

Start with ​ this code​ . Create a custom error message using “fmt.Errorf”.
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

	bs, err := toJSON(p1)

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) []byte {
	bs, _ := json.Marshal(a)
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