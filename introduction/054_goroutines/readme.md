#  Go Routines - Description

-

# Instructions

* Using goroutines, create an incrementer program
  * have a variable to hold the incrementer value
  * launch a bunch of goroutines
    * each goroutine should
      * read the incrementer value
        * store it in a new variable
      * yield the processor with runtime.Gosched()
      * increment the new variable
      * write the value in the new variable back to the incrementer variable
* use waitgroups to wait for all of your goroutines to finish
* the above will create a race condition.
* Prove that it is a race condition by using the -race flag
* if you need help, here is a hint: ​ https://play.golang.org/p/FYGoflKQej

# How to use it

* Download main.go
* Save it in some folder
* Execute `go run ./my/path/main.go`

## Author(s)

* Carlos Mendez

### Created

Jun 22, 2020