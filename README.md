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
### if else
```go
if  condition-1 { 
    // code to be executed if condition-1 is true
} 
else if condition-2 {
    // code to be executed if condition-2 is true
} 
else {
    // code to be executed if both condition1 and condition2 are false
}
```

### if else with initialization
```go
if x := 100; x == 100 {
	fmt.Println("Germany")
}
```

### Switch case
```go
switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
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

### struts
```go
/*Declare a strut*/
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo		//If we have to declare one strut inside another, we can just give its name.
}

/*Assign values to strut (Least desirable) - Approach 1*/
deadpool := person{"Ryan", "Renolds"}

/*Assign values to strut - Approach 2*/
var captainAmerica person
captainAmerica.firstName = "Steve"
captainAmerica.lastName = "Rodgers"

/*Assign values to strut - Approach 3*/
/*For multiline declaration, it is mandatory to end all the lines with a comma*/
ironMan := person{
	firstName: "Tony",
	lastName:  "Stark",
	contactInfo: contactInfo{
		email:   "tonystark@starkindustries.com",
		zipCode: 10010,
	},
}

/*Print formatted output of strut*/
fmt.Printf("%+v", ironMan)
```
