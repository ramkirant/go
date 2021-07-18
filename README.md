# Go cheatsheet

### Variables in Go

#### Approach 1
Declare the variable using var and the variable datatype as below
```go
var card string = "Ace of spades"
```
#### Approach 2
Here, Go assumes the datatype of the variable using the value assigned to it
```go
card := "Ace of spades"
```

### Custom datatypes with receivers
```go
type deck []string
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}
```





