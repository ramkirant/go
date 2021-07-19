# Go cheatsheet

### Variables in Go
```go
/*Approach 1*/
/*Declare the variable using var and the variable datatype as below*/
var card string = "Ace of spades"

/*Approach 2*/
/*Here, Go assumes the datatype of the variable using the value assigned to it*/
card := "Ace of spades"

/*Type Conversion*/
byteSlice := []byte("Hi There")
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

/*Length of a slice*/
len(slice)

/*Convert a slice of string to a single string*/
strings.Json(stringSlice, ",")
```

### Functions
```go
/*Function without a return value*/
func printDeck() {
    cards := newDeck()
    cards.print()
}

/*Function with a return value*/
func newDeck() deck {
    cards := newDeck()
    return cards		
}

/*Function with arguments and multiple return values*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*Assigning multiple return values to variables*/
hand, remainingCards := deal(cards, 5)
```

### Null check in go
Use nil for all null checks in go
```go
if error != nil {
	fmt.Println("Error: ",error)
	os.Exit(1)
}
```
