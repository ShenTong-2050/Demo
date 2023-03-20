package composite

// Origination 组织人数接口
type Origination interface {
	Count()		int
}

// Person 成员结构体
type Person struct {
	Name 	string
}

// Department 部门名称
type Department struct {
	Name 		string
	depart 		[]Origination
}

func (p *Person) Count() int {
	return 1
}

func (d *Department) add(o Origination) {
	d.depart = append(d.depart,o)
}

// Count 部门总人数
func (d *Department) Count() int {
	num := 0
	for _,d := range d.depart {
		num += d.Count()
	}
	return num
}

func NewDepartNum() Origination {
	var d = &Department{Name: "root"}
	for i :=0; i<5; i++ {
		// d.add 是给 d.depart 切片中添加元素
		// 由于 Person 与 Department 结构体 均 实现了 Origination 接口
		// 所以 Person 与 Department 均可 往 d.depart 中添加
		d.add(&Person{})
		d.add(&Department{Name: "sub department",depart: []Origination{&Person{}}})
		// d.add(&Department{Name: "sub department",depart: []Origination{&Department{Name: "sub sub department",depart: []Origination{&Department{}}}}})
	}
	return d
}