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

### Slices
```go
/*Define a slice*/
cards := []string{"a","b"}

/*Add elements to slice*/
cards := append(cards,"c")

/*Iterating a slice*/
for _, card := range cards {
	fmt.Println(card)
}

/*Splitting a slice. Approach 1*/
slice[startIndexIncluding, endIndexNotIncluding]

/*Splitting a slice. Approach 2*/
slice[:endIndexNotIncluding] --Fetch all the elements from start till the end index

/*Splitting a slice. Approach 3*/
slice[startIndex:] --Fetch all the elements from the start index till the end
```

### Functions





