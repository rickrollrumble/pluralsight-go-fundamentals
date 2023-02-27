## Ways to run a Go app
1. `go run <filename>`

2. `go  build`

    Used to build an executable. Going into the package folder and running `go build` will create an executable with the same name as the package.

3. `go install` 
    
    Creates an executable in the `bin` folder of the `go` folder. Location can be changed using `GOPATH` and `GOBIN`, which are set using 
    
    `go env -w GOPATH="<path>"` and 
    
    `go env -w GOBIN="<path>"`

## Working with constants and variables

A variable can be declared at the package level and remain unused. But within a function, unused variables will result in errors.

Go passes arguments by values, not reference. It can be overriden by using pointers. When a value without the `&` is passed to a function, a copy of the input value is made, operated upon, and returned. The original value remains unchanged.

However, when the address to a variable is passed with the `&variableName` notation, the value at the location is changed directly.

### Environment Variables
These are settings on your working system. Can be obtained using `os` package.