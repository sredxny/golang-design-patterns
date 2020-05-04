package builder

type SedanBuilder struct{
	WindowType string
	SeatType string
	RadioType string
}

func newSedanBuilder() *SedanBuilder {
	return &SedanBuilder{}
}

func (b *SedanBuilder) SetWindowType(){
	b.WindowType = "small sedan windows"
}

func (b *SedanBuilder) SetSeatsType(){
	b.SeatType = "small sedan seats"
}

func (b *SedanBuilder) SetRadioType(){
	b.SeatType = "radio to listen music with family"
}

func (b *SedanBuilder) GetCar() Car {
	return Car{
		WindowType: b.WindowType,
		SeatsType:  b.SeatType,
		RadioType:  b.RadioType,
	}
}
