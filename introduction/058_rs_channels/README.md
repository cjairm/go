#  Bidirectional Channel - Description


# Instructions

* Get this code work: 
```
func main() {
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
```

# How to use it

* Download main.go
* Save it in some folder
* Exexute: `go run ./my/path/main.go`

## Author(s)

* Carlos Mendez

### Created

Jun 28, 2020