package builder

const(
	TRUCK = "truck"
	SEDAN = "sedan"
	MOTORCYCLE = "motorcycle"
)

type CarBuilder interface {
	SetWindowType()
	SetSeatsType()
	SetRadioType()
	GetCar() Car
}

func GetBuilder(builderType string) CarBuilder{
	switch builderType {
	case TRUCK:
		return &TruckBuilder{}
	case SEDAN:
		return &SedanBuilder{}
	default:
		return nil
	}

	return nil
}