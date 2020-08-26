package infra

type Initailizer interface {
	Init()
}

type InitailizerRegister struct {
	Initailizers []Initailizer
}

func (i *InitailizerRegister) Register(ai Initailizer) {
	i.Initailizers = append(i.Initailizers, ai)
}
