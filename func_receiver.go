package main

import "fmt"

type person struct {
  name    string
  surname string
  age     int
}

func (p person) print(){
  fmt.Printf("%s %s is %d old\n", p.name, p.surname, p.age)
}

func (p person) yearsToAge(age int){
  if p.age > age {
    fmt.Printf("The age of %d has been exceeded by %d years\n", age, p.age-age)
  } else if p.age < age{
    fmt.Printf("The age of %d has been reached %d years ago\n", age, age-p.age)
  } else {
    fmt.Printf("The age of %d is now\n", age)
  }
}

func main() {
  alex := person{name: "Alex", surname:"Bello", age: 18}

  fmt.Printf("%s %s is %d old\n", alex.name, alex.surname, alex.age)
  alex.print()

  alex.yearsToAge(25)
}
