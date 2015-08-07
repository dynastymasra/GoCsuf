### GoCsuf 01_Hello

#### Golang workspaces
bin/
pkg/
src/

#### GOPATH environment varibale
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin

#### Import library
mkdir $GOPATH/src/github.com/dynastymasra/library_name
go build github.com/dynastymasra/library_name < This won't produce an output file
go install github.com/dynastymasra/package_name
