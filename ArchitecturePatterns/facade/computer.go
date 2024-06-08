package facade

import "fmt"

type CPU struct {
	speed float32
}

func NewCPU(speed float32) *CPU {
	return &CPU{speed: speed}
}

func (c *CPU) Run() {
	fmt.Println("CPU Run at speed of ", c.speed, "GHZ")
}

type Memory struct {
	capacity int
}

func NewMemory(capacity int) *Memory {
	return &Memory{capacity: capacity}
}

func (m *Memory) Load() {
	fmt.Println("Memory Load with  ", m.capacity, "GB")
}

type Disk struct {
	size int
}

func NewDisk(size int) *Disk {
	return &Disk{size: size}
}

func (d *Disk) Read() {
	fmt.Println("Disk reading data with ", d.size, "KB/s")
}
