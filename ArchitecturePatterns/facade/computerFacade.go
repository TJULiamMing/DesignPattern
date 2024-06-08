package facade

type ComputerFacade struct {
	cpu    *CPU
	memory *Memory
	disk   *Disk
}

func NewComputerFacaxde(speed float32, capacity int, size int) *ComputerFacade {
	return &ComputerFacade{
		cpu:    NewCPU(speed),
		memory: NewMemory(capacity),
		disk:   NewDisk(size),
	}
}

func (cf *ComputerFacade) Run() {
	cf.cpu.Run()
	cf.memory.Load()
	cf.disk.Read()
}
