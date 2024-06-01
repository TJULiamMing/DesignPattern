package builder

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) BuildHouse() {
	d.builder.BuildWall()
	d.builder.BuildDoor()
	d.builder.BuildWindow()
}

func (d *Director) B() {}
