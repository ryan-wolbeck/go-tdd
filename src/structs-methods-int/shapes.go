package structsmethodsints

import "math"

//Rectangle : struct for the parameters of a rectanlge
type Rectangle struct {
	Width  float64
	Height float64
}

//Area : Method to take in a rectangle object and return the area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Circle : struck for the parameters of a circle
type Circle struct {
	Radius float64
}

//Shape : interface to allow either struct to be passed to the area method properly
type Shape interface {
	Area() float64
}

//Area : Method to take in the radius of a circle and return it's area
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Perimeter : Takes in width and height and returns the perimeter of the rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2*rectangle.Width + 2*rectangle.Height
}

//Area : Takes in width and height and returns the area of a rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
