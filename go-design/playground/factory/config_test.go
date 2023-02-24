package factory

import (
	"reflect"
	"testing"
)

func TestCreateFactory(t *testing.T) {
	type args struct {
		rule string
	}

	type test struct {
		name  string
		args  args
		want  IRuleConfigFactory
	}

	var tests = []test{
		{
			name: "GetJsonRuleFactory",
			args: args{rule: "json"},
			want: JsonRuleConfigFactory{},
		},
		{
			name: "GetYamlRuleFactory",
			args: args{rule: "yaml"},
			want: YamlRuleConfigFactory{},
		},
		{
			name: "GetXmlRuleFactory",
			args: args{rule: "xml"},
			want: XmlRuleConfigFactory{},
		},
	}

	for _,val := range tests {
		t.Run(val.name, func(t *testing.T) {
			if got := CreateFactory(val.args.rule); !reflect.DeepEqual(got,val.want) {
				t.Errorf("got %v, want %v",got,val.want)
			}
		})
	}
}
