crate
=====

Go library collected for common usage.

### Q & A
#### why remove import "github.com/***"
it's bad when you are coding go project.
To compile or test,clone to dir/src/crate and make dir as gopath
```sh
git clone https://github.com/tanxunrong/crate crate-dir/src/crate
export GOPATH=$GOPATH:./crate-dir
```