package demo1

import (
    "fmt"
    "mymath"
)

func demo(a int) (b int) {
    fmt.Println(a)
    b = a + 100
    return
}

func add(a int, b int, c *int) {
    *c = a + b
}

type Person struct {
    name string
    age  int
}

func (p Person) print() {
    fmt.Println("name:", p.name, " age:", p.age)
}

func Run() {
    fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))

    s := "hello"
    fmt.Printf("%s\n", s[1:3])
    var arr [10]int
    for i := 0; i < 10; i++ {
        arr[i] = i
    }

    sl := arr[2:5]

    fmt.Printf("%d, cap:%d\n", sl, cap(sl))
    sl = append(sl, 100, 200, 300, 400, 500)
    sl = append(sl, 600)
    fmt.Printf("%d\n", arr)

    var m map[string]int
    m = make(map[string]int)
    m["one"] = 1
    fmt.Printf("m[one] = %d\n", m["one"])

    rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
    csharpRating, ok := rating["C#"]
    if ok {
        fmt.Println("c# rating is", csharpRating)
    } else {
        fmt.Println("c# rating is not found")
        rating["C#"] = 9.9
    }
    delete(rating, "C")
    for k, v := range rating {
        fmt.Printf("map key:%s  value:%f\n", k, v)
    }

    ret := demo(3)
    fmt.Println(ret)

    add1 := 22
    add2 := 33
    add(add1, add2, &ret)
    fmt.Println(ret)

    ptr := new(Person)
    ptr.name = "liucong"
    ptr.age = 12
    fmt.Println(ptr)

    var p1 Person
    p1.name, p1.age = "bad", 12
    fmt.Println(p1)
    p1.print()
}
