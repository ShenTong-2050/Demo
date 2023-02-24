package builder

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResourcePoolConfig(t *testing.T) {
	/*type arg struct {
		name string
		rf    []ResourcePoolConfigOptionFunc
	}
	type test struct {
		name 		string
		args 		arg
		want 		*ResourcePoolConfigOption
		wantError	bool
	}

	var tests = []test{
		{
			name: "empty name",
			args: arg{name: "empty name"},
			want: &ResourcePoolConfigOption{
				MaxTotal: 10,
				MaxIdle: 9,
				MinIdle: 1,
			},
			wantError: false,
		},
		{
			name: "max total < max idle",
			args: arg{name: "max total < max idle"},
			want: &ResourcePoolConfigOption{
				MaxTotal: 10,
				MaxIdle: 9,
				MinIdle: 1,
			},
			wantError: false,
		},
		{
			name: "success request",
			args: arg{name: "success request"},
			want: &ResourcePoolConfigOption{
				MaxTotal: 10,
				MaxIdle: 9,
				MinIdle: 1,
			},
			wantError: false,
		},
	}

	for _,val := range tests {
		t.Run(val.name, func(t *testing.T) {
			got,err := NewResourcePoolConfig(val.args.name)
			require.Equalf(t, err != nil,val.wantError,"new resource pool config error: %v, want error %v",err,val.wantError)
			assert.Equal(t, got,val.want)
		})
	}*/

	type arg struct {
		name string
	}
	type test struct {
		name string
		args arg
		want *ResourcePoolConfigOption
		wantError bool
	}
	var tests = []test{
		{
			name:"empty name",
			args:arg{name:"empty name"},
			want:&ResourcePoolConfigOption{
				MaxTotal:10,
				MaxIdle:9,
				MinIdle:1,
			},
			wantError:false,
		},{
			name:"maxIdle is Invalid",
			args:arg{name:"maxIdle is Invalid"},
			want:&ResourcePoolConfigOption{
				MaxTotal:10,
				MaxIdle:9,
				MinIdle:1,
			},
			wantError:false,
		},
		{
			name:"success",
			args:arg{name:"success"},
			want:&ResourcePoolConfigOption{
				MaxTotal:10,
				MaxIdle:9,
				MinIdle:1,
			},
			wantError:false,
		},
	}

	for _,val := range tests {
		t.Run(val.name,func(t *testing.T){
			got,err := NewResourcePoolConfig(val.args.name)
			require.Equalf(t,err!=nil,val.wantError,"NewResourcePoolConfig get error %v, want error %v",err,val.wantError)
			assert.Equal(t,got,val.want)
		})
	}
}
