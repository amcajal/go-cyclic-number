# go-cyclic-number
Toy project to test golang toolset, using as pretext the "cyclic number" problem (famous 142857 number)

First version: https://github.com/amcajal/go-cyclic-number/commit/cabeb0c0102900ca0c0737e0054ba2776d8b98c7

First version just checks that the program produces the expected result. No test files (no TDD obviously), overkill decisions (the recursive function) and no insight about program performance.
GDB was used to check how debugging in Go works. So far, it worked well. Need to take a look into [delve](https://github.com/go-delve/delve) The feedback provided by delve is quite misleading. An error "error layer=debugger can't find build-id note on binary" is reported when launching $> dlv debug; But it seems to work properly.

Second version: statements working as tests have been moved from "main.go" file to proper "_test.go" file. Creation of test cases is quite easy. Followed the tutorial at
https://go.dev/doc/tutorial/add-a-test

Third version: a "resources usage" test is added to the cyclic_test.go, in order to retrieve profiling information. Profiling in go is quite easy with the "go test <profiling-flags>" and "go tool pprof". Profiling performed follown tutorial at https://golangdocs.com/profiling-in-golang

Fourth version: https://github.com/amcajal/go-cyclic-number/commit/9f9c322769e856a29e0f73eb1636b63ff9fe4e6b
  
In the fourth version, "IsCyclic" is modified, so instead of being a recursive function, it performs a plain for loop, generating the same results (but increasing the overall readability of the function). This change is good to test the TDD principles:
- Test goes first: it is changed; Test does not compile (good)
- Source code is modified; Test then compiles and still passing (good)
  
However, no performance improvements take place with this change.
  
Test coverage is already at 100%. Test coverage is really easy to obtain using the "go test -cover" command and "go tool" commands (see https://golangdocs.com/code-coverage-in-golang)
