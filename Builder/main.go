package main

import (
	"fmt"
	"github.com/sredxny/golang-design-patterns/Builder/builder"
)

func main() {
	truckBuilder := builder.GetBuilder(builder.TRUCK)
	sedanBuilder := builder.GetBuilder(builder.SEDAN)

	director := builder.NewDirector(truckBuilder)
	truck := director.BuildCar()

	fmt.Printf("truck... Windows: %+v Seats: %+v Radio: %+v", truck.WindowType, truck.SeatsType, truck.RadioType)

	director.SetBuilder(sedanBuilder)
	sedan := director.BuildCar()
	fmt.Printf("sedan... Windows: %+v Seats: %+v Radio: %+v", sedan.WindowType, sedan.SeatsType, sedan.RadioType)

}
