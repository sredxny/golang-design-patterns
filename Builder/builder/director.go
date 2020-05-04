package builder

type director struct {
	builder CarBuilder
}

func NewDirector(b CarBuilder) *director{
	return &director{
		b,
	}
}

func (d *director) SetBuilder(b CarBuilder){
	d.builder = b
}

func (d *director) BuildCar() Car {
	d.builder.SetRadioType()
	d.builder.SetSeatsType()
	d.builder.SetWindowType()
	return d.builder.GetCar()
}

