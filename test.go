package main

import "fmt"
//import "time"
// time.Sleep(1e9) 1秒钟

func Add (x, y int, ch chan int) {
    z := x + y
    ch <- z
}

func main() {
    // make(chan int, 1)创建一个缓存的channel，无缓存的channel在写入数据时会阻塞等待读取
    ch := make(chan int, 1)

    for {
        select {
            case ch <- 1:
            case ch <- 0:
        }
        i := <- ch
        fmt.Println("Receive value:", i)
    }
}

