# go-broom
Utility functions for slices, maps and functions.
All broom-classes wrap the default behaviour, so that f.e. len() can still be called, a bslice or bmap can still be iterated in a for loop, etc.  

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
- [ ] bfunc
    - [ ] bfunc.Args - map a bslice to function args
    - [ ] bfunc.Ret - map multi return values to single bslice
- [ ] async version of some functions like bslice.EachAsync
- [ ] support the types like string,int,float where possible 


See the documentation for:
* [bslices](./bslice/bslice.md)
* [bmaps](./bmap/bmap.md)