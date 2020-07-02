# Godoc - Description

"
Before writing documentation, we are going to look at reading documentation. There are several
things to know about documentation:
* godoc.org
	* standard library and third party package documentation
* golang.org
	* standard library documentation
* go doc
	* command to read documentation at the command line
* godoc
	* command to read documentation at the command line
	* also can run a local server showing documentation
". Todd McLeod


# Instructions

Create a dog package. The dog package should have an exported func “Years” which takes
human years and turns them into dog years (1 human year = 7 dog years). Document your
code with comments. Use this code in func main.
* Run your program and make sure it works
* Run a local server with godoc and look at your documentation.

# How to use it

* Download main.go
* Save it in some folder
* Exexute: `go run ./my/path/main.go`
	* To check documentation `godoc -http=:6060`
		* Then check `http://localhost:6060/`

# Resources


* [Command godoc](https://godoc.org/golang.org/x/tools/cmd/godoc)
* [Godoc: documenting Go code](https://blog.golang.org/godoc)

## Author(s)

* Carlos Mendez

### Created

Jul 1, 2020