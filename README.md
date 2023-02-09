# go_downLoadCheck
> ## documentation
- installed and configured Go
- install go vs code extensions
- to install go vs code tools press ` ctrl + shift + p ` inside vs code, select all/check all boxes to install tools
- create module folder in root of project OR where the developer wants to hold source code for the project,
- open code editor in the root of project/ developer disired source code folder.  
- if in  project root or developer desired root file path
- inside code editor run `go mod init`. This will create go.mod file and add initial information to organize source code.
- if already have github run ` go mod init [github project http path] ` will create go.mod file to organize source code
- if needed edit module declaration with the hosting root source code. mine is github and looks like this
`module github.com/wilmer88/go_downLoadCheck` `go1.20` this will organize project/source code in the disired root repository
- to run program. with the go.mod file module declaration run `go run [source code hosting path]` or ` go run .` with out  executable file
- to run executable file run ` "file name.exe" ` in the terminal/ command prompt (make sure your not in git bash) of your code editor in the correct directory/root of project  


> creating models
- in root of project creat a folder with the name of models
- inside created models folder create a file named `desiered name.go` file
- inside that file start by declaring package name ` package models `
- declare struts and variables

> after building http server and routing logic

- `go build .` creates project executable file 
- to run executable run ` "the exucutable file name" ` in the terminal/ command prompt (not git bash) of your code editor in the correct directory/root of project  
- open browser to ` localhost:3000/nameOfTheController` generating a static route

## in practice folder or project
- functions
- arguments and parameters
- returning results
- Methods 
- branching
- controlling program flow
- struts
- switches
- data types
- Looping
- loop till condition
- loop till condition with post clause
- infinite loops
- loop over collections
- if statement
- else if
- else
- panic method
- slices
- creating server code
- build command 


