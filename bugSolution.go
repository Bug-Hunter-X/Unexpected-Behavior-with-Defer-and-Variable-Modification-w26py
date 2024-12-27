```go
func main() {
    var i int
    fmt.Println(i) // This will print 0
    defer func() {
        fmt.Println(i)
    }()
    i = 10
}
```