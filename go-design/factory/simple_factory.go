package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

// yamlRuleConfigParser 实现了 Parse 接口
type yamlRuleConfigParser struct {}

func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("yaml implements")
}

// jsonRuleConfigParser 实现了 Parse 接口
type jsonRuleConfigParser struct {}

func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("json implements")
}

// NewIRuleConfig 返回的 结构体 都 实现了 IRuleConfigParser 接口
func NewIRuleConfig(s string) IRuleConfigParser {
	switch s {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}




