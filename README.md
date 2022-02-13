# go-cyclic-number
Toy project to test golang toolset, using as pretext the "cyclic number" problem (famous 142857 number)

First version: https://github.com/amcajal/go-cyclic-number/commit/cabeb0c0102900ca0c0737e0054ba2776d8b98c7

First version just checks that the program produces the expected result. No test files (no TDD obviously), overkill decisions (the recursive function) and no insight about program performance.
GDB was used to check how debugging in Go works. So far, it worked well. Need to take a look into [delve](https://github.com/go-delve/delve) The feedback provided by delve is quite misleading. An error "error layer=debugger can't find build-id note on binary" is reported when launching $> dlv debug; But it seems to work properly.

Second version: statements working as tests have been moved from "main.go" file to proper "_test.go" file. Creation of test cases is quite easy. Followed the tutorial at
https://go.dev/doc/tutorial/add-a-test

