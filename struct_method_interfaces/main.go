package main

import (
    "fmt"
    "math"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////// STRUCTS ///////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////
// type person struct {
//     name string
//     age int
//     income float64
//     company string
// }
//
// // default if we initialize stuff here
// // company is default here (if everyone is employee) THEORETICALLY
// func newPerson(name string, age int, income float64) *person {
//     return &person {
//         name: name,
//         age: age,
//         income: income,
//         // everyone works for Johnson Brothers here
//         company: "Johnson Brothers",
//     }
// }

///////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////// METHODS ///////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////
// type rect struct {
//     width, height int
// }
//
// func (r *rect) area() int {
//     return r.width * r.height
// }
//
// func (r *rect) perim() int {
//     return 2*r.width + 2*r.height
// }

///////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////// INTERFACE /////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////
type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}

type circ struct {
    radius float64
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (r rect) area() float64 {
    return r.width * r.height
}

func (c circ) perim() float64 {
    return 2 * math.Pi * c.radius
}

func (c circ) area() float64 {
    return math.Pi * c.radius * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////// MAIN //////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
    // STRUCTS
    // employee1 := newPerson("Grace", 24, 65000)
    // fmt.Println(*employee1)
    //
    // employee2 := newPerson("Mckenna", 22, 50000)
    // fmt.Printf("\nName: %v\n", employee2.name)
    // fmt.Printf("Age: %v\n", employee2.age)
    // fmt.Printf("Income: %v\n", employee2.income)
    // fmt.Printf("Company: %v\n", employee2.company)

    // METHODS
    // r := rect{width: 10, height: 4}
    //
    // // defining new struct here
    // rp := rect{width: 20, height: 10}
    //
    // fmt.Println("area: ", r.area())
    // fmt.Println("perimeter: ", r.perim())
    //
    // fmt.Println("area: ", rp.area())
    // fmt.Println("perimeter: ", rp.perim())

    // INTERFACING
    cirlce := circ{radius: 5}
    rectangle := rect{width: 4, height: 3}

    measure(cirlce)
    measure(rectangle)
}
