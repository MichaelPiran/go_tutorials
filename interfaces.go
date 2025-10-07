package main

/*
 *  In Go, an interface is a type that defines a set of method signatures. 
 *  Any type that implements all of the methods defined in the interface is said to satisfy the interface.
 *  This allows you to write code that is generic and can be used with any type that satisfies the interface.
 */

import (
  "fmt"
  "math"
)

type rect struct{
  width, height float64
}
type circle struct{
  radius float64
}

type geometry interface {
    area() float64
    perim() float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    /* 
     * Func measure expect a geometry interface.
     * struct rect has received method area an perim 
     * so it implements a geometry interface
     */
    measure(r)
    measure(c)
}
