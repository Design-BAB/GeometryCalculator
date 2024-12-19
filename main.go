package main

import (
  "fmt"
  "math"
)

//this is the circle object
type Circle struct {
  radius float32
}
//this is the circleconstructor
func newCircle(radius float32) *Circle {
  c := Circle{radius: radius}
  return &c
}
func (c Circle) circleArea() float64 {
  radius := float64(c.radius)
  area := radius * radius
  area = area * math.Pi
  area = area * 100
  area = math.Round(area)
  return area / 100
}

//this is the Rectangle object
type rectangle struct {
  length float32
  width float32
}
//this is the rectangle constructor
func newRectangle(length float32, width float32) *rectangle {
  r := rectangle{length: length, width: width}
  return &r
}
func (r rectangle) rectangleArea() float32 {
  return r.length * r.width
}


//Now for the triangle!
type triangle struct {
  base float32
  height float32
}
//this is the rectangle constructor
func newTriangle(base float32, height float32) *triangle {
  t := triangle{base: base, height: height}
  return &t
}
func (t triangle) triangleArea() float32 {
  return t.base * t.height * 0.5
}

func main()  {
  var input int
  var number1 float32
  var number2 float32
  for {
    fmt.Println("\n\nGeometry Calculator")
    fmt.Println("1. Calculate the Area of a circle 🟢")
    fmt.Println("2. Calculate the area of a rectangle 🟦")
    fmt.Println("3. Calcualte the area of a triangle 🛆")
    fmt.Println("4. To quit")
    fmt.Print("Please ender you choice: ")
    _, err := fmt.Scanln(&input)

		if err != nil || input < 1 || input > 4 {
			fmt.Println("Invalid input. Please enter a number between 1 and 4.")
			continue
		}
    
    switch input {
    case 1:
        fmt.Println("Ok, time to calculate the area of the circle!")
        fmt.Println("Please enter the radius: ")
        fmt.Scanln(&number1)
        circle := newCircle(number1)
        fmt.Println("The area of the circle is:", circle.circleArea())
    case 2:
        fmt.Println("Ok, time to calculate the area of the rectangle!")
       fmt.Println("Please enter the length: ")
       fmt.Scanln(&number1)
       fmt.Println("Please enter the width: ")
       fmt.Scanln(&number2)
       Rectangle := newRectangle(number1, number2)
       fmt.Println("The area of the rectangle is:", Rectangle.rectangleArea())
    case 3:
        fmt.Println("Now we will calculate the area of the triangle!")
        fmt.Println("Please enter the base")
        fmt.Scanln(&number1)
        fmt.Println("Now please enter the height")
        fmt.Scanln(&number2)
        triangle := newTriangle(number1, number2)
        fmt.Println("The area of the triable is: ", triangle.triangleArea())
    case 4:
        fmt.Println("Have a nice day!")
    }
    if input == 4 {
      break
    }
  }
}
