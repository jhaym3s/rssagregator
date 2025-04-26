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
	car := struct {
		Make  string
		Model string
		Year  int
	}{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2020,
	}
	blueShape := rectangle{
		length: 5,
		width:  10,
	}

	mapSample := map[string]map[string]int{
		"John": {
			"age":  30,
			"height": 180,
		},		}
	fmt.Println("Area of rectangle:", blueShape.area())
	const name= "John Doe"
	//const person int  = name
	 someone := Person{Name: name, Age: 30}
	 mother := woman{
		Person:  someone,
		hipSize: 36,
	}
	printShapeInfo(blueShape)
	myCar := car
	fmt.Println("person  details:", mother)
	fmt.Println("Car details:", myCar)
	
}