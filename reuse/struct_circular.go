package reuse

type Fruit struct {
	Name string
	Veg  *Veg
}

type Veg struct {
	Name  string
	Fruit *Fruit
}

func NewFruit(name string) *Fruit {
	return &Fruit{Name: name, Veg: nil}
}

func NewVeg(name string) *Veg {
	return &Veg{Name: name, Fruit: nil}
}

func (f *Fruit) Tweak(veg *Veg) {
	f.Veg = veg
}

func (v *Veg) Tweak(fruit *Fruit) {
	v.Fruit = fruit
}
