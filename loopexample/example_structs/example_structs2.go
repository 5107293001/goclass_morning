package examplestructs

type rectangle struct {
	width, height float32
}
type circle struct {
	radius float32
}

func (r rectangle) area() float32 {
	return r.height * r.width
}

func (r rectangle) perimeter() float32 {
	return r.height*2 + r.width*2
}

func (r circle) area() float32 {
	return 2 * 3.14 * r.radius * r.radius

}
func (r circle) perimeter() float32 {
	return 2 * 3.14 * r.radius

}

func InitExampleStructs2() {
	rectangle := rectangle{
		height: 2,
		width:  4,
	}

	measure(rectangle, "rectangle")

	circle := circle{
		radius: 10,
	}

	measure(circle, "circle")

}
