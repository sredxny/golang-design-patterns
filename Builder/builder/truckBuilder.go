package builder

type TruckBuilder struct{
	WindowType string
	SeatType string
	RadioType string
}

func newTruckBuilder() *TruckBuilder {
	return &TruckBuilder{}
}

func (b *TruckBuilder) SetWindowType(){
	b.WindowType = "big truck windows"
}

func (b *TruckBuilder) SetSeatsType(){
	b.SeatType = "big truck seats"
}

func (b *TruckBuilder) SetRadioType(){
	b.SeatType = "radio to listen much much music"
}

func (b *TruckBuilder) GetCar() Car {
	return Car{
		WindowType: b.WindowType,
		SeatsType:  b.SeatType,
		RadioType:  b.RadioType,
	}
}