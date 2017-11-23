# go-broom
Utility functions for slices, maps and functions

## Create a new broom-slice
```golang
package main

import (
	"github.com/jpolack/go-broom/slice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := slice.New(mySimpleSlice)
}
```

## Iterate the slice
```golang
package main

import (
	"github.com/jpolack/go-broom/slice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := slice.New(mySimpleSlice)

	myBroomSlice.Each(func(number interface{}, i int) {
        //do something for each element
	})
}
```


## Map the slice values
```golang
package main

import (
	"github.com/jpolack/go-broom/slice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := slice.New(mySimpleSlice)

	newBroomSlice := myBroomSlice.Map(func(number interface{}, i int) interface{} {
		return number.(int) * 2
    })

    //newBroomSlice contains 2, 4, 6 - myBroomSlice stays unchanged
}
```

## Reduce the slice values
```golang
package main

import (
	"github.com/jpolack/go-broom/slice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := slice.New(mySimpleSlice)

	totalSum := myBroomSlice.Reduce(func(sum interface{}, number interface{}, i int) interface{} {
		return sum.(int) + number.(int)
	}, 0)

    //totalSum equals 6
}
```


## Filter the slice values
```golang
package main

import (
	"github.com/jpolack/go-broom/slice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := slice.New(mySimpleSlice)

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

	"github.com/jpolack/go-broom/slice"
)

func main() {
	mySimpleSlice := []int{1, 2, 3}
	myBroomSlice := slice.New(mySimpleSlice)

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