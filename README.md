## Ways to run a Go app
1. `go run <filename>`

2. `go  build`

    Used to build an executable. Going into the package folder and running `go build` will create an executable with the same name as the package.

3. `go install` 
    
    Creates an executable in the `bin` folder of the `go` folder. Location can be changed using `GOPATH` and `GOBIN`, which are set using 
    
    `go env -w GOPATH="<path>"` \n and 
    
    `go env -w GOBIN="<path>"`