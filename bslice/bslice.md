# broom-slice
## Create a new broom-slice

A broomslice is a slice wrapper that adds some functionalities

```golang
package main

import (
	"github.com/jpolack/go-broom/bslice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := bslice.NewOf(mySimpleSlice)

	//myBroomSlice is a slice containing 1,2,3

	myBroomSlice2 := bslice.New()
	//myBroomSlice is a empty slice
}
```

## Iterate the bslice

If you change myBroomSlice in the callback it does not affect the iteration since the iteration works on a copy of myBroomSlice

```golang
package main

import (
	"github.com/jpolack/go-broom/bslice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := bslice.NewOf(mySimpleSlice)

	myBroomSlice.Each(func(number interface{}, i int) {
        //do something for each element
	})
}
```


## Map the bslice values
```golang
package main

import (
	"github.com/jpolack/go-broom/bslice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := bslice.NewOf(mySimpleSlice)

	newBroomSlice := myBroomSlice.Map(func(number interface{}, i int) interface{} {
		return number.(int) * 2
    })

    //newBroomSlice contains 2, 4, 6 - myBroomSlice stays unchanged
}
```

## Reduce the bslice values
```golang
package main

import (
	"github.com/jpolack/go-broom/bslice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := bslice.NewOf(mySimpleSlice)

	totalSum := myBroomSlice.Reduce(func(sum interface{}, number interface{}, i int) interface{} {
		return sum.(int) + number.(int)
	}, 0)

    //totalSum equals 6
}
```


## Filter the bslice values
```golang
package main

import (
	"github.com/jpolack/go-broom/bslice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := bslice.NewOf(mySimpleSlice)

	newBroomSlice := myBroomSlice.Filter(func(number interface{}, i int) bool {
		return number.(int)%2!=0
	})

    //newBroomSlice contains 1, 3 - myBroomSlice stays unchanged
}
```


## Combine Functions
```golang
package main

import (
	"fmt"

	"github.com/jpolack/go-broom/bslice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := bslice.NewOf(mySimpleSlice)

	sum := myBroomSlice.Map(func(number interface{}, i int) interface{} {
		return number.(int) + 2
	}).Each(func(number interface{}, i int) {
		fmt.Printf("Index: %v, number: %v\n", i, number)
	}).Filter(func(number interface{}, i int) bool {
		return number.(int)%2 != 0
	}).Reduce(func(sum interface{}, number interface{}, i int) interface{} {
		fmt.Printf("Sum is %v adding %v\n", sum, number)
		return sum.(int) + number.(int)
	}, 0)

	fmt.Printf("Everthing added sums up to: %v\n", sum)
}
```