package prototype

type Prototype interface {
	Print(indentation string)
	Clone() Prototype
}
