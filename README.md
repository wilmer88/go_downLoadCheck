# go_downLoadCheck
> ### documentation
- installed and configured Go
- install go vs extensions
- install go vs code tools press ` ctrl + shift + p ` inside vs code and select all/check all boxes to install tools
- create module folder in root of project OR where the developer wants to hold source code for the project,
- open code editor in the root of project/ developer disired source code folder.  
- if in  project root or developer desired root file path
- inside code editor run `go mod init`. This will create go.mod file and add initial information to organize source code.
- if already have github run ` go mod init [github project http path] ` will create go.mod file to organize source code
- if needed edit module declaration with the hosting root source code. mine is github and looks like this
`module github.com/wilmer88/go_downLoadCheck` `go1.20`
- this will organize project/source code in the disired root repository
- to run program. with the go.mod file module declaration run `go run [source code hosting path]` or ` go run .`
> creating models
- in root of project creat a folder with the name of models
- inside created models folder create a file named `desiered name.go` file
- inside that file start by declaring package name ` package models `
- declare struts and variables
