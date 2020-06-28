#  Bidirectional Channel - Description


# Instructions

* Get this code work: 
```
func main() {
	cr := make(<-chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
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