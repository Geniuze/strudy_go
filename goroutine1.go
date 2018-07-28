package main

import "fmt"

func Add (x, y int, ch chan int) {
    z := x + y
    ch <- z
}

func main() {
    chs := make([]chan int, 10)
    for i:=0; i<10; i++ {
        chs[i] = make(chan int)
        go Add(i, i, chs[i])
    }

    for _,ch := range chs {
        z := <-ch
        fmt.Println(z)
    }
}

