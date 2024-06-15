package observer

import "fmt"

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.NotifyAll()
}

func (i *Item) Register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) Deregister(o Observer) {
	for j, observer := range i.observerList {
		if observer == o {
			i.observerList = append(i.observerList[:j], i.observerList[j+1:]...)
			break
		}
	}
}

func (i *Item) NotifyAll() {
	for _, observer := range i.observerList {
		observer.Update(i.name)
	}
}
