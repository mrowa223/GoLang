package pkg

type ServerPc struct {
	Cpu     int
	Memory  int
	Wrapper Wrapper
}

func (pc ServerPc) GetPrice() float64 {
	return pc.Wrapper.GetPrice() * float64(pc.Cpu) * float64(pc.Memory)
}
