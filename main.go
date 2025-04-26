package main 

import "fmt"

type Person struct {
	Name string
	Age  int
}

type woman struct {
	Person// embedding Person struct which is compulsory for every person 
	hipSize int  // something specific to women
}

func printShapeInfo(s shape){
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
	fmt.Println("Shape type:", s)
	// fmt.Println("Shape type:", s.(circle))
 //fmt.Println("Shape type:", s.(rectangle))
	// fmt.Println("Shape type:", s.(shape))
	
}

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {	
	length  float64
	width   float64	 
}
type circle struct {
	radius float64
}
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}
func (r rectangle) area() float64 {
	return r.length * r.width
}

//closures 
func addMultipleTime() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
	
}


func main()  {
	x := 10
	y := &x
	z := *y
	fmt.Println("print %v and %v",y,z)

	
}