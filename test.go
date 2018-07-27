package main

import "fmt"

type Integer int

func (a Integer) Less (b Integer) bool {
    return a<b
}
func (a *Integer) Add (b Integer) {
    *a += b
}

type Base struct {
    Name string
}
func (base *Base) Foo(){...}
func (base *Base) Bar(){...}

type Foo struct {
    Base
    ...
}

func (foo *Foo) Bar(){
    foo.Base.Bar()
}

func main(){
    var i Integer = 3
    fmt.Println("a Less 5", i.Less(5))
    i.Add(5)
    fmt.Println("a Add 5 =", i)
}
