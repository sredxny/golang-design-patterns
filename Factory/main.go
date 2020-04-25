package main

func main() {
	factory := &ShapeFactory{}

	circle, err := factory.GetShape(CIRCLE)
	if err == nil {
		circle.Draw()
	}

	square, err := factory.GetShape(SQUARE)
	if err == nil{
		square.Draw()
	}
}
