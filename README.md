# go-broom
Utility functions for slices, maps and functions.

All broom-classes wrap the default behaviour, so that f.e. len() can still be called, a bslice or bmap can still be iterated in a for loop, etc.

## [broom-slice](./bslice/bslice.md)
Tired of writing your mapper/filter/reduce functions? I did that for you! All you have to do is some type assertions.

See [bslice](./bslice/bslice.md) for a documentation

## [broom-map](./bmap/bmap.md) 
Threads like [this (nice way of getting a slice of values from a map?)](https://stackoverflow.com/questions/13422578/in-golang-is-there-a-nice-way-of-getting-a-slice-of-values-from-a-map) asked for a nice way and here it is!

See [bmap](./bmap/bmap.md) for a documentation

## [broom-func](./bfunc/bfunc.md)
`multiple value in single-value context`? Never again!

See [bfunc](./bfunc/bfunc.md) for a documentation


## Todos
- [x] bslice
    - [x] bmap.Each
    - [x] bmap.EachAsync
    - [x] bmap.Map
    - [x] bmap.Filter
    - [x] bmap.Reduce
- [x] bmap
    - [x] bmap.Values - get map values as bslice
    - [x] bmap.Keys - get map keys as bslice
    - [x] bmap.Each - iterate over map
    - [x] bmap.Map - map values in bmap to new bmap 
    - [x] bmap.Contains - get if a key is present in a bmap
- [x] bfunc
    - [x] bfunc.Args - map a bslice to function args
- [ ] async version of some functions like bslice.EachAsync
- [ ] support the types like string,int,float where possible 


## Contributions and more
- Please let me know about featurewishes!
- Contributors are welcome!
- I am not that used to open source development so please let me know where I can improve :shit: 