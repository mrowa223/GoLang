package pkg

type HomePc struct {
	Cpu          int
	GraphicsCard int
	Wrapper      Wrapper
}

func (pc HomePc) GetPrice() float64 {
	return pc.Wrapper.GetPrice() * float64(pc.Cpu) * float64(pc.GraphicsCard)
}
