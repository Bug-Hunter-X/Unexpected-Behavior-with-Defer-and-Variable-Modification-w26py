# Go Defer Statement Bug

This example demonstrates a common pitfall in Go when using the `defer` statement with variables that are modified after the defer call.

The `bug.go` file contains the code exhibiting the issue. The output is unexpected because the deferred `fmt.Println(i)` prints the value of `i` as it was *at the time of the defer call*, not its value at the time of execution. 

The `bugSolution.go` file shows the corrected approach.