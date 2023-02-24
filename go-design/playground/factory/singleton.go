package factory

import "fmt"

type Bun interface {
	Create()
}

type MeatBuns struct {}

type VegetableBuns struct {}

func (m MeatBuns) Create() {
	fmt.Println("meat buns")
}

func (v VegetableBuns) Create() {
	fmt.Println("vegetable buns")
}

func SimpleCreateFactory(t string) Bun {
	switch t {
	case "meat":
		return MeatBuns{}
	case "vegetable":
		return VegetableBuns{}
	}
	return nil
}


