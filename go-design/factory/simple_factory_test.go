package factory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfig(t *testing.T) {

	type args struct {
		typeName	string
	}

	tests := []struct{
		name	string
		arg		args
		want 	IRuleConfigParser
	}{
		{
			name: "json",
			arg: args{typeName : "json"},
			want: jsonRuleConfigParser{},
		},
		{
			name: "yaml",
			arg: args{typeName: "yaml"},
			want: yamlRuleConfigParser{},
		},
	}

	for _,val := range tests {
		t.Run(val.name, func(t *testing.T) {
			if got := NewIRuleConfig(val.arg.typeName); !reflect.DeepEqual(got,val.want) {
				t.Errorf("NewIRuleConfigParser() => %v, want => %v",got,val.want)
			}
		})
	}
}
