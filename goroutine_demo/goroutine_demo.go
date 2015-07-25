package goroutine_demo

import (
    "fmt"
    //"runtime"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        //runtime.Gosched()
        fmt.Println(s)
    }
}

func sum(a []int, c chan int) {
    total := 0
    for _, v := range a {
        total += v
    }
    c <- total
}

func block(c chan int, output chan int) {
    a := <-c
    b := <-c
    output <- a + b
}

func fibonacci(c, quit chan int) {
    x, y := 1, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func Run() {
    go say("hello")
    say("world")

    a := []int{7, 3, 2, 1, 4, 5, 2, 1}
    c := make(chan int)
    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)
    x, y := <-c, <-c
    fmt.Println(x, y)

    c2 := make(chan int, 2)
    c2 <- 100
    c2 <- 200
    go block(c2, c)
    z := <-c
    fmt.Println(z)

    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
    //fmt.Println(<-c)
}
