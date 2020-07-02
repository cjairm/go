# Testing - Description

"
* Tests must:
	* be in a file that ends with “​ _test.go​ ”
	* put the file in ​ the same package​ as the one being tested
	* be in a func with a signature “​ func TestXxx(*testing.T)​ ”
* Run a test `go test`
* Deal with test failure
	* use ​ t.Error​ to signal failure
* Nice idiom
	* expected
	* got
		* Example. t.Error("Got: ", value, "Expected: ", valueExpected)
". Todd McLeod

## Benckmark

"
Part of the testing package allows us to measure the speed of our code. This could also be
called “measuring the performance” of your code, or “benchmarking” your code - finding out how
fast the code runs
". Todd McLeod

## Coverage

"
Coverage in programming is how much of our code is covered by tests. We can use the “-cover”
flag to run coverage analysis on our code. We can use the flag and required file name
“-coverprofile <some file name>” to write our coverage analysis to a file.
". Todd McLeod

* go test -cover
	* go test -coverprofile c.out
		* show in browser:
			* go tool cover -html=c.out
		* learn more
			* go tool cover -h

# Examples

Here is a review of the different commands useful with benchmarks, examples, and tests.
* godoc -http=:8080
* go test
* go test -bench .
	* don’t forget the “.” in the line above
* go test -cover
* go test -coverprofile c.out
* go tool cover -html=c.out

# Resources

* [Testable Examples in Go](https://blog.golang.org/examples)

## Author(s)

* Carlos Mendez

### Created

Jul 2, 2020