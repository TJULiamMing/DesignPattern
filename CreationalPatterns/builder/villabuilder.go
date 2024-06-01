package builder

type Villabuilder struct {
	Wall   string
	Door   string
	Window string
}

func (v *Villabuilder) BuildWall() {
	v.Wall = "wall"
}

func (v *Villabuilder) BuildDoor() {
	v.Door = "door"
}

func (v *Villabuilder) BuildWindow() {
	v.Window = "window"
}

func (v *Villabuilder) GetHouse() string {
	return v.Wall + " " + v.Door + " " + v.Window
}
