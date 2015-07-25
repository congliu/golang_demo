package demo2

import (
    "fmt"
    "reflect"
)

type Worker struct {
    name string
    age  int
}

type Staff struct {
    Worker
    level int
}

type Manager struct {
    Staff
    title string
}

func (w Worker) StartWork() {
    fmt.Printf("worker %s starts work\n", w.name)
}

func (s Staff) StartWork() {
    fmt.Printf("Staff %s (level:%d) sit in office\n", s.name, s.level)
}

func (m Manager) StartWork() {
    fmt.Printf("Manger %s (title:%s) holds a meeting\n", m.name, m.title)
}

type IWorker interface {
    StartWork()
}

func Run() {
    mike := Worker{"Mike", 23}
    paul := Staff{Worker{"Paul", 26}, 3}
    dike := Manager{Staff{Worker{"Dick", 30}, 3}, "VP"}
    var li [3]IWorker
    li[0] = mike
    li[1] = paul
    li[2] = dike
    for _, value := range li {
        value.StartWork()
    }

    var x float64 = 3.4
    v := reflect.ValueOf(x)
    t := reflect.TypeOf(x)
    fmt.Println(v.Kind() == reflect.Float64)
    fmt.Println(t)

}
