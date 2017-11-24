# broom-map
## Create a new broom-map

A broommap is a map wrapper that adds some functionalities

```golang
package main

import (
	"github.com/jpolack/go-broom/bmap"
)

func main() {
	mySimpleMap := map[int]int{1: 2, 2: 4, 3: 9}
	myBroomMap := bmap.NewOf(mySimpleMap)
	//myBroomMap is a map containing 1: 2, 2: 4, 3: 9

	myBroomMap2 := bmap.New()
	//myBroomMap2 is an empty map
}
```

## Get the map Values as a slice

```golang
package main

import (
	"github.com/jpolack/go-broom/bmap"
)

func main() {
	mySimpleMap := map[int]int{1: 2, 2: 4, 3: 9}
	myBroomMap := bmap.NewOf(mySimpleMap)

	myBroomMap.Values() //2,4,9 (not garantueed in that order)
}
```

## Get the map Keys as a slice

```golang
package main

import (
	"github.com/jpolack/go-broom/bmap"
)

func main() {
	mySimpleMap := map[int]int{1: 2, 2: 4, 3: 9}
	myBroomMap := bmap.NewOf(mySimpleMap)

	myBroomMap.Keys() //1,2,3 (not garantueed in that order)
}
```

## Check if a map contains a certain key

```golang
package main

import (
	"github.com/jpolack/go-broom/bmap"
)

func main() {
	mySimpleMap := map[int]int{1: 2, 2: 4, 3: 9}
	myBroomMap := bmap.NewOf(mySimpleMap)

	myBroomMap.Contains(1) //true
	myBroomMap.Contains(4) //false
}
```

## Iterate the map

If you change myBroomMap in the callback it does not affect the iteration since the iteration works on a copy of myBroomMap

```golang
package main

import (
	"github.com/jpolack/go-broom/bmap"
)

func main() {
	mySimpleMap := map[int]int{1: 2, 2: 4, 3: 9}
	myBroomMap := bmap.NewOf(mySimpleMap)

	myBroomMap.Each(func(key interface{}, value interface{}) {
        //do something for each element
	})
}
```

## Map the map

```golang
package main

import (
	"github.com/jpolack/go-broom/bmap"
)

func main() {
	mySimpleMap := map[int]int{1: 2, 2: 4, 3: 9}
	myBroomMap := bmap.NewOf(mySimpleMap)

	myBroomMap.Map(func(key interface{}, value interface{}) interface{}{
        //map each element
	})
}
```