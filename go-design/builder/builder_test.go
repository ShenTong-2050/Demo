package builder

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResourcePoolConfigBuilder(t *testing.T) {
	/*type arg struct {
		name 	string
		maxTotal,maxIdle,minIdle int
	}

	type test struct {
		name 		string
		args		arg
		builder 	*resourcePoolConfigBuilder
		want    	*resourcePoolConfig
		wantError	bool
	}

	var tests = []test{
		{
			name: "FirstRequest",
			args: arg{name:"FirstRequest",maxTotal: 11,maxIdle: 9,minIdle: 2},
			builder: &resourcePoolConfigBuilder{},
			want: &resourcePoolConfig{
				name: "FirstRequest",
				maxTotal: defaultMaxTotal,
				maxIdle: defaultMaxIdle,
				minIdle: defaultMinIdle,
			},
			wantError: false,
		},
		{
			name: "SecondRequest",
			args: arg{name:"SecondRequest",maxTotal: 12,maxIdle: 10,minIdle: 2},
			builder: &resourcePoolConfigBuilder{},
			want: &resourcePoolConfig{
				name: "SecondRequest",
				maxTotal: defaultMaxTotal,
				maxIdle: 13,
				minIdle: 2,
			},
			wantError: true,
		},
		{
			name: "ThirdRequest",
			args: arg{name:"ThirdRequest",maxTotal: 13,maxIdle: 9,minIdle: 4},
			builder: &resourcePoolConfigBuilder{},
			want: &resourcePoolConfig{
				name: "ThirdRequest",
				maxTotal: 13,
				maxIdle: 9,
				minIdle: -1,
			},
			wantError: true,
		},
	}

	for _,val := range tests {
		t.Run(val.name, func(t *testing.T) {
			val.builder.setName(val.args.name)
			val.builder.setMaxTotal(val.args.maxTotal)
			val.builder.setMaxIdle(val.args.maxIdle)
			val.builder.setMinIdle(val.args.minIdle)

			got,err := val.builder.Build()

			require.Equalf(t, val.wantError,err != nil,"Build() error = %v, wantError = %v",err,val.wantError)

			assert.Equal(t, got,val.want)

		})
	}*/

	type test struct {
		name string
		builder *resourcePoolConfigBuilder
		want  *resourcePoolConfig
		wantError bool
	}

	var tests = []test {
		{
			name:"FirstRequest",
			builder:&resourcePoolConfigBuilder{
				name: "FirstRequest",
				maxTotal: 11,
				maxIdle: 12,
				minIdle: 1,
			},
			want:&resourcePoolConfig{
				name:"FirstRequest",
				maxTotal:11,
				maxIdle:10,
				minIdle:1,
			},
			wantError:false,
		},
		{
			name:"SecondRequest",
			builder:&resourcePoolConfigBuilder{
				name: "SecondRequest",
				maxTotal: 12,
				maxIdle: 11,
				minIdle: 1,
			},
			want:&resourcePoolConfig{
				name:"SecondRequest",
				maxTotal:12,
				maxIdle:11,
				minIdle:1,
			},
			wantError:false,
		},
		{
			name:"ThirdRequest",
			builder:&resourcePoolConfigBuilder{
				name: "ThirdRequest",
				maxTotal: 13,
				maxIdle: 12,
				minIdle: 1,
			},
			want:&resourcePoolConfig{
				name:"ThirdRequest",
				maxTotal:13,
				maxIdle:12,
				minIdle:1,
			},
			wantError:false,
		},
	}

	for _,val := range tests {
		t.Run(val.name,func(t *testing.T) {
			// 设置 请求资源名称
			val.builder.setName(val.name)

			got,err := val.builder.Build()

			// 判断 当前测试 是否出错
			require.Equalf(t,err != nil,val.wantError,"got error is %v",err)

			// 判断 获取的结果是否正确
			assert.Equal(t,got,val.want)
		})
	}

}
