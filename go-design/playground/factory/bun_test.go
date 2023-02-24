package factory

import (
	"testing"
)

func TestGenerateBuns(t *testing.T) {

	type test struct {
		name,want  string
		factory interface{}
	}

	var tests = []test{
		{
			name: "BJMeatBuns",
			want: "Beijing meat buns",
			factory: new(BJFactory).Generate("meat").create(),
		},
		{
			name: "BJVegetableBuns",
			want: "Beijing vegetable buns",
			factory: new(BJFactory).Generate("vegetable").create(),
		},
		{
			name: "SHMeatBuns",
			want: "Shanghai meat buns",
			factory: new(SHFactory).Generate("meat").create(),
		},
		{
			name: "SHVegetableBuns",
			want: "Shanghai vegetable buns",
			factory: new(SHFactory).Generate("vegetable").create(),
		},
	}

	for _,val := range tests {
		t.Run(val.name, func(t *testing.T) {
			if val.factory != val.want {
				t.Errorf("Got: %v , Want: %v",val.factory,val.want)
			}
		})
	}
}
