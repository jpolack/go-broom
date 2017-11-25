# broom-function
## Create a new broom-func

A broom-func is a slice wrapper that adds some functionalities

```golang
package main

import (
	"github.com/jpolack/go-broom/bfunc"
)

func main() {
	broomFunc := bfunc.NewOf(func(x int, y int) int {
		return x + y
    })
    
    broomFunc(1,2)   //[]interface{}{3}
    broomFunc(1,2,3) //panic
}
```

## Call a brrom-func with a slice

```golang
package main

import (
	"github.com/jpolack/go-broom/bslice"
)

func main() {
	broomFunc := bfunc.NewOf(func(x int, y int) int {
		return x + y
    })
    
    broomFunc([]int{1,2})   //[]interface{}{3}
    broomFunc([]int{1,2,3}) //panic
}
```