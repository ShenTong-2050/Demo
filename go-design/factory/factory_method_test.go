package factory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParserFactory(t *testing.T) {

	type args struct {
		typeName	string
	}

	tests := []struct{
		name   string
		args   args
		want   IRuleConfigParserFactory
	}{
		{
			name : "json",
			args : args{typeName: "json"},
			want : jsonRuleConfigParserFactory{},
		},
		{
			name: "yaml",
			args: args{typeName: "yaml"},
			want: yamlRuleConfigParserFactory{},
		},
	}
	
	for _,val := range tests {
		t.Run(val.name, func(t *testing.T) {
			if got := NewIRuleConfigParserFactory(val.args.typeName); !reflect.DeepEqual(got,val.want) {
				t.Errorf("NewIRuleConfigParserFactory() => %v, want => %v",got,val.want)
			}
		})
	}

}
